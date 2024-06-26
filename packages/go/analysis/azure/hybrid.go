// Copyright 2024 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package azure

import (
	"context"
	"errors"
	"fmt"

	"github.com/specterops/bloodhound/analysis"
	"github.com/specterops/bloodhound/dawgs/graph"
	"github.com/specterops/bloodhound/dawgs/query"
	"github.com/specterops/bloodhound/dawgs/util/channels"
	"github.com/specterops/bloodhound/graphschema/ad"
	"github.com/specterops/bloodhound/graphschema/azure"
	"github.com/specterops/bloodhound/graphschema/common"
	"github.com/specterops/bloodhound/log"
)

func PostHybrid(ctx context.Context, db graph.Database) (*analysis.AtomicPostProcessingStats, error) {
	tenants, err := FetchTenants(ctx, db)
	if err != nil {
		return &analysis.AtomicPostProcessingStats{}, fmt.Errorf("hybrid post processing: %w", err)
	}

	operation := analysis.NewPostRelationshipOperation(ctx, db, "Hybrid Attack Paths Post Processing")

	err = db.ReadTransaction(ctx, func(tx graph.Transaction) error {
		for _, tenant := range tenants {
			if tenantUsers, err := EndNodes(tx, tenant, azure.Contains, azure.User); err != nil {
				return err
			} else if len(tenantUsers) == 0 {
				return nil
			} else {
				for _, tenantUser := range tenantUsers {
					innerTenantUser := tenantUser
					if err := operation.Operation.SubmitReader(func(ctx context.Context, tx graph.Transaction, outC chan<- analysis.CreatePostRelationshipJob) error {
						var adUser *graph.Node

						if onPremUserID, hasOnPremUser, err := hasOnPremUser(innerTenantUser); err != nil {
							return err
						} else if hasOnPremUser {
							if adUser, err = tx.Nodes().Filter(query.Equals(query.NodeProperty(common.ObjectID.String()), onPremUserID)).First(); err != nil {
								if errors.Is(err, graph.ErrNoResultsFound) {
									if adUser, err = createMissingAdUser(ctx, db, onPremUserID); err != nil {
										return fmt.Errorf("error attempting to create missing AD User node: %w", err)
									}
								} else {
									return err
								}
							}

							SyncedToEntraUserRelationship := analysis.CreatePostRelationshipJob{
								FromID: adUser.ID,
								ToID:   innerTenantUser.ID,
								Kind:   ad.SyncedToEntraUser,
							}

							if !channels.Submit(ctx, outC, SyncedToEntraUserRelationship) {
								return nil
							}

							SyncedToADUserRelationship := analysis.CreatePostRelationshipJob{
								FromID: innerTenantUser.ID,
								ToID:   adUser.ID,
								Kind:   azure.SyncedToADUser,
							}

							if !channels.Submit(ctx, outC, SyncedToADUserRelationship) {
								return nil
							}
						}

						return nil
					}); err != nil {
						return err
					}
				}
			}
		}

		if err != nil {
			return err
		}

		return tx.Commit()
	})

	if opErr := operation.Done(); opErr != nil {
		return &operation.Stats, fmt.Errorf("marking operation as done: %w; transaction error (if any): %v", opErr, err)
	}

	return &operation.Stats, nil
}

func hasOnPremUser(node *graph.Node) (string, bool, error) {
	if onPremSyncEnabled, err := node.Properties.Get(azure.OnPremSyncEnabled.String()).Bool(); errors.Is(err, graph.ErrPropertyNotFound) {
		return "", false, nil
	} else if err != nil {
		return "", false, err
	} else if onPremID, err := node.Properties.Get(azure.OnPremID.String()).String(); errors.Is(err, graph.ErrPropertyNotFound) {
		return onPremID, false, nil
	} else if err != nil {
		return onPremID, false, err
	} else {
		return onPremID, (onPremSyncEnabled && len(onPremID) != 0), nil
	}
}

func createMissingAdUser(ctx context.Context, db graph.Database, objectID string) (*graph.Node, error) {
	var (
		err     error
		newNode *graph.Node
	)

	log.Debugf("Matching AD User node with objectID %s not found, creating a new one", objectID)
	properties := graph.AsProperties(map[string]any{
		common.ObjectID.String(): objectID,
	})

	err = db.WriteTransaction(ctx, func(tx graph.Transaction) error {
		newNode, err = tx.CreateNode(properties, ad.Entity, ad.User)
		return err
	})

	return newNode, err
}
