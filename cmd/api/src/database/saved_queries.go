// Copyright 2023 Specter Ops, Inc.
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

package database

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/specterops/bloodhound/src/model"
	"gorm.io/gorm"
)

type SavedQueriesData interface {
	ListSavedQueries(ctx context.Context, userID uuid.UUID, order string, filter model.SQLFilter, skip, limit int) (model.SavedQueries, int, error)
	CreateSavedQuery(ctx context.Context, userID uuid.UUID, name string, query string, description string) (model.SavedQuery, error)
	DeleteSavedQuery(ctx context.Context, id int) error
	SavedQueryBelongsToUser(ctx context.Context, userID uuid.UUID, savedQueryID int) (bool, error)
	GetSharedSavedQueries(ctx context.Context, userID uuid.UUID) (model.SavedQueries, error)
	GetPublicSavedQueries(ctx context.Context) (model.SavedQueries, error)
}

func (s *BloodhoundDB) ListSavedQueries(ctx context.Context, userID uuid.UUID, order string, filter model.SQLFilter, skip, limit int) (model.SavedQueries, int, error) {
	var (
		queries model.SavedQueries
		result  *gorm.DB
		count   int64
		cursor  = s.Scope(Paginate(skip, limit)).WithContext(ctx).Where("user_id = ?", userID)
	)
	// change result to query permissions. check if one from hackathon fufills the requirements. probably need to add scope?
	if filter.SQLString != "" {
		cursor = cursor.Where(filter.SQLString, filter.Params)
		result = s.db.Model(&queries).WithContext(ctx).Where("user_id = ?", userID).Where(filter.SQLString, filter.Params).Count(&count)
	} else {
		result = s.db.Model(&queries).WithContext(ctx).Where("user_id = ?", userID).Count(&count)
	}

	if result.Error != nil {
		return queries, 0, result.Error
	}

	if order != "" {
		cursor = cursor.Order(order)
	}

	// if name != "" {
	// 	cursor = cursor.Order(order)
	// }

	// if description != "" {
	// 	cursor = cursor.Where("description = ?", description)
	// }

	// if query != "" {
	// 	cursor = cursor.Order(order)
	// }

	// if scope != "" {
	// this needs to account for multiple
	// 	cursor = cursor.Order(order)
	// }

	// result = cursor.Find(&queries)

	result = cursor.Joins("JOIN saved_queries_permissions sqp ON sqp.query_id = saved_queries.id AND (sqp.global OR saved_queries.user_id = ?)", userID).Find(&queries)

	return queries, int(count), CheckError(result)
}

func (s *BloodhoundDB) CreateSavedQuery(ctx context.Context, userID uuid.UUID, name string, query string, description string) (model.SavedQuery, error) {
	savedQuery := model.SavedQuery{
		UserID:      userID.String(),
		Name:        name,
		Query:       query,
		Description: description,
	}

	return savedQuery, CheckError(s.db.WithContext(ctx).Create(&savedQuery))
}

func (s *BloodhoundDB) DeleteSavedQuery(ctx context.Context, id int) error {
	return CheckError(s.db.WithContext(ctx).Delete(&model.SavedQuery{}, id))
}

func (s *BloodhoundDB) SavedQueryBelongsToUser(ctx context.Context, userID uuid.UUID, savedQueryID int) (bool, error) {
	var savedQuery model.SavedQuery
	if result := s.db.WithContext(ctx).First(&savedQuery, savedQueryID); result.Error != nil {
		return false, CheckError(result)
	} else if savedQuery.UserID == userID.String() {
		return true, nil
	} else {
		return false, nil
	}
}

// GetSharedSavedQueries returns all the saved queries that the given userID has access to, including global queries
func (s *BloodhoundDB) GetSharedSavedQueries(ctx context.Context, userID uuid.UUID) (model.SavedQueries, error) {
	savedQueries := model.SavedQueries{}

	result := s.db.WithContext(ctx).Where("shared_to_user_id = ?", userID).Find(&savedQueries)

	return savedQueries, CheckError(result)
}

// GetPublicSavedQueries returns all the queries that were shared publicly
func (s *BloodhoundDB) GetPublicSavedQueries(ctx context.Context) (model.SavedQueries, error) {
	savedQueries := model.SavedQueries{}

	result := s.db.WithContext(ctx).Where("public = true").Find(&savedQueries)
	return savedQueries, CheckError(result)
}
