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

package v2

import (
	"github.com/specterops/bloodhound/log"
	"github.com/specterops/bloodhound/src/model"
	"net/http"

	"github.com/specterops/bloodhound/src/auth"
	"github.com/specterops/bloodhound/src/ctx"

	"github.com/specterops/bloodhound/dawgs/util"
	"github.com/specterops/bloodhound/src/api"
	"github.com/specterops/bloodhound/src/queries"
)

type CypherSearch struct {
	Query             string `json:"query"`
	IncludeProperties bool   `json:"include_properties,omitempty"`
}

func (s Resources) CypherSearch(response http.ResponseWriter, request *http.Request) {
	var (
		payload       CypherSearch
		authCtx       = ctx.FromRequest(request).AuthCtx
		auditLogEntry model.AuditEntry
		preparedQuery queries.PreparedQuery
		err           error
	)

	if err := api.ReadJSONRequestPayloadLimited(&payload, request); err != nil {
		api.WriteErrorResponse(
			request.Context(),
			api.BuildErrorResponse(http.StatusBadRequest, "JSON malformed.", request), response,
		)
		return
	}

	if preparedQuery, err = s.GraphQuery.PrepareCypherQuery(payload.Query); err != nil {
		api.WriteErrorResponse(
			request.Context(),
			api.BuildErrorResponse(http.StatusBadRequest, err.Error(), request), response,
		)
		return
	}

	if preparedQuery.HasMutation {
		if !s.Authorizer.AllowsPermission(authCtx, auth.Permissions().GraphDBMutate) {
			s.Authorizer.AuditLogUnauthorizedAccess(request)
			api.WriteErrorResponse(
				request.Context(),
				api.BuildErrorResponse(http.StatusForbidden, "Permission denied: User may not modify the graph.", request), response,
			)
			return
		}

		// All mutation attempts must be audit logged even when failed
		if auditLogEntry, err = model.NewAuditEntry(
			model.AuditLogActionMutateGraph,
			model.AuditLogStatusIntent,
			model.AuditData{"query": preparedQuery.StrippedQuery}); err != nil {
			api.WriteErrorResponse(
				request.Context(),
				api.BuildErrorResponse(http.StatusInternalServerError, api.ErrorResponseDetailsInternalServerError, request),
				response,
			)
			return
		}

		// create an intent audit log
		if err := s.DB.AppendAuditLog(request.Context(), auditLogEntry); err != nil {
			api.WriteErrorResponse(
				request.Context(),
				api.BuildErrorResponse(http.StatusInternalServerError, "failure creating an intent audit log", request),
				response,
			)
			return
		}
	}

	if graphResponse, err := s.GraphQuery.RawCypherSearch(request.Context(), preparedQuery, payload.IncludeProperties); err != nil {
		// audit log needs to be marked as failed if a mutation is present
		if preparedQuery.HasMutation {
			auditLogEntry.Status = model.AuditLogStatusFailure
		}
		if queries.IsQueryError(err) {
			api.WriteErrorResponse(
				request.Context(),
				api.BuildErrorResponse(http.StatusBadRequest, err.Error(), request), response,
			)
		} else if util.IsNeoTimeoutError(err) {
			api.WriteErrorResponse(request.Context(), api.BuildErrorResponse(http.StatusInternalServerError, "transaction timed out, reduce query complexity or try again later", request), response)
		} else {
			api.WriteErrorResponse(
				request.Context(),
				api.BuildErrorResponse(http.StatusInternalServerError, err.Error(), request), response,
			)
		}
	} else {
		// audit log needs to be marked as success if a mutation is present
		if preparedQuery.HasMutation {
			auditLogEntry.Status = model.AuditLogStatusSuccess
		}
		api.WriteBasicResponse(request.Context(), graphResponse, http.StatusOK, response)
	}

	if preparedQuery.HasMutation {
		if err := s.DB.AppendAuditLog(request.Context(), auditLogEntry); err != nil {
			// We don't want to update response as failed because having info on the mutation response trumps this error
			log.Errorf("failure to create mutation audit log %s", err.Error())
		}
	}
}
