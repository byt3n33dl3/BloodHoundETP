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

package v2

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/specterops/bloodhound/dawgs/graph"
	"github.com/specterops/bloodhound/dawgs/ops"
	"github.com/specterops/bloodhound/log"
	"github.com/specterops/bloodhound/src/api"
	"github.com/specterops/bloodhound/src/model"
)

type DatabaseManagement struct {
	DeleteCollectedGraphData bool `json:"deleteCollectedGraphData"`
	DeleteHighValueSelectors bool `json:"deleteHighValueSelectors"`
	DeleteFileIngestHistory  bool `json:"deleteFileIngestHistory"`
	DeleteDataQualityHistory bool `json:"deleteDataQualityHistory"`
	AssetGroupId             int  `json:"assetGroupId"`
}

func (s Resources) HandleDatabaseWipe(response http.ResponseWriter, request *http.Request) {

	var (
		payload DatabaseManagement
		nodeIDs []graph.ID
		// use this struct to flag any fields that failed to delete
		errors DatabaseManagement
	)

	if err := api.ReadJSONRequestPayloadLimited(&payload, request); err != nil {
		api.WriteErrorResponse(
			request.Context(),
			api.BuildErrorResponse(http.StatusBadRequest, "JSON malformed.", request),
			response,
		)
		return
	}

	if payload.DeleteHighValueSelectors && payload.AssetGroupId == 0 {
		api.WriteErrorResponse(
			request.Context(),
			api.BuildErrorResponse(http.StatusBadRequest, "please provide an assetGroupId to delete", request),
			response,
		)
		return
	}

	commitID, err := uuid.NewV4()
	if err != nil {
		api.WriteErrorResponse(
			request.Context(),
			api.BuildErrorResponse(http.StatusInternalServerError, fmt.Sprintf("failure generating uuid: %v", err.Error()), request),
			response,
		)
		return
	}

	auditEntry := &model.AuditEntry{
		Action: "DeleteBloodhoundData",
		Model: &model.AuditData{
			"options": payload,
		},
		Status:   model.AuditStatusIntent,
		CommitID: commitID,
	}

	// create an intent audit log
	if err := s.DB.AppendAuditLog(request.Context(), *auditEntry); err != nil {
		api.WriteErrorResponse(
			request.Context(),
			api.BuildErrorResponse(http.StatusInternalServerError, "failure creating an intent audit log", request),
			response,
		)
		return
	}

	// delete graph
	if payload.DeleteCollectedGraphData {

		if err := s.Graph.ReadTransaction(request.Context(), func(tx graph.Transaction) error {
			fetchedNodeIDs, err := ops.FetchNodeIDs(tx.Nodes())

			nodeIDs = append(nodeIDs, fetchedNodeIDs...)
			return err
		}); err != nil {
			log.Errorf("%s: %s", "error fetching all nodes", err.Error())
			errors.DeleteCollectedGraphData = true

		} else if err := s.Graph.BatchOperation(request.Context(), func(batch graph.Batch) error {
			for _, nodeId := range nodeIDs {
				// deleting a node also deletes all of its edges due to a sql trigger
				if err := batch.DeleteNode(nodeId); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Errorf("%s: %s", "error deleting all nodes", err.Error())
			errors.DeleteCollectedGraphData = true
		} else {
			// if succesful, kick off analysis
			s.TaskNotifier.RequestAnalysis()
		}

	}

	// delete custom high value selectors
	if payload.DeleteHighValueSelectors {
		if err := s.DB.DeleteAssetGroupSelectorsForAssetGroup(request.Context(), payload.AssetGroupId); err != nil {
			log.Errorf("%s %d: %s", "there was an error deleting asset group with id = ", payload.AssetGroupId, err.Error())
			errors.DeleteHighValueSelectors = true
		} else {
			// if succesful, kick off analysis
			s.TaskNotifier.RequestAnalysis()
		}

	}

	// delete file ingest history
	if payload.DeleteFileIngestHistory {
		if err := s.DB.DeleteAllFileUploads(); err != nil {
			log.Errorf("%s: %s", "there was an error deleting file ingest history", err.Error())
			errors.DeleteFileIngestHistory = true
		}
	}

	// delete data quality history
	if payload.DeleteDataQualityHistory {
		if err := s.DB.DeleteAllDataQuality(); err != nil {
			log.Errorf("%s: %s", "there was an error deleting data quality history", err.Error())
			errors.DeleteDataQualityHistory = true
		}
	}

	// append a failure audit log if anything failed
	if errors.DeleteCollectedGraphData || errors.DeleteHighValueSelectors || errors.DeleteDataQualityHistory || errors.DeleteFileIngestHistory {
		auditEntry.Status = model.AuditStatusFailure
		auditEntry.Model = model.AuditData{
			"deletionFailures": buildMessageForFailureAudit(errors),
		}

		if err := s.DB.AppendAuditLog(request.Context(), *auditEntry); err != nil {
			api.WriteErrorResponse(
				request.Context(),
				api.BuildErrorResponse(http.StatusInternalServerError, "there was an error creating audit log for deleting Bloodhound data", request),
				response,
			)
			return
		}

		// return a user friendly error message indicating what operations failed
		api.WriteErrorResponse(
			request.Context(),
			api.BuildErrorResponse(http.StatusInternalServerError, fmt.Sprintf("we encountered an error while deleting %s.  please submit your request again.", buildMessageForFailureAudit(errors)), request),
			response,
		)
		return
	}

	// otherwise append a success audit log
	auditEntry.Status = model.AuditStatusSuccess
	if err := s.DB.AppendAuditLog(request.Context(), *auditEntry); err != nil {
		api.WriteErrorResponse(
			request.Context(),
			api.BuildErrorResponse(http.StatusInternalServerError, "there was an error creating audit log for deleting Bloodhound data", request),
			response,
		)
		return
	}

	response.WriteHeader(http.StatusNoContent)
}

func buildMessageForFailureAudit(failures DatabaseManagement) string {
	var message []string
	if failures.DeleteCollectedGraphData {
		message = append(message, "collected graph data")
	} else if failures.DeleteDataQualityHistory {
		message = append(message, "data quality history")
	} else if failures.DeleteFileIngestHistory {
		message = append(message, "file ingest history")
	} else if failures.DeleteHighValueSelectors {
		message = append(message, "high value selectors")
	}

	return strings.Join(message, ", ")
}