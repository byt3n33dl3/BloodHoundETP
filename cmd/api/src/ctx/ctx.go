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

package ctx

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/specterops/bloodhound/errors"
	"github.com/specterops/bloodhound/src/auth"
)

// Use our own type rather than a primitive to avoid collisions
// (https://staticcheck.io/docs/checks#SA1029)
type CtxKey string

const ValueKey = CtxKey("ctx.bhe")

type RequestedWaitDuration struct {
	Value   time.Duration
	UserSet bool
}

// Context holds contextual data that is passed around to functions. This is an extension to Golang's built in context.
type Context struct {
	StartTime time.Time
	Timeout   RequestedWaitDuration
	RequestID string
	AuthCtx   auth.Context
	Host      *url.URL
	RequestIP string
}

func (s *Context) ConstructGoContext() context.Context {
	return context.WithValue(context.Background(), ValueKey, s)
}

// WithUserSession adds the supplied AuthCtx value to the BloodHound Context structure
func (s *Context) WithUserSession(userSession auth.Context) *Context {
	s.AuthCtx = userSession
	return s
}

// WithRequestID adds the supplied RequestID value to the BloodHound Context structure
func (s *Context) WithRequestID(requestID string) *Context {
	s.RequestID = requestID
	return s
}

func (s *Context) WithHost(host *url.URL) *Context {
	s.Host = host
	return s
}

// FromRequest extracts the Golang-builtin-Context from a request and converts it to a BloodHound Context struct
func FromRequest(request *http.Request) *Context {
	return Get(request.Context())
}

// Get converts a Golang-builtin-Context into a BloodHound-defined Context struct
func Get(ctx context.Context) *Context {
	if ctx == nil {
		return &Context{}
	} else if rawValue := ctx.Value(ValueKey); rawValue == nil {
		return &Context{}
	} else if bhCtx, ok := rawValue.(*Context); !ok {
		panic(fmt.Sprintf("Context value for %q was not the the expected type. Wanted Context but got %T.", ValueKey, rawValue))
	} else {
		return bhCtx
	}
}

// Set takes the given golang context and stores the given bh context struct inside it using a well known key
func Set(ctx context.Context, bhCtx *Context) context.Context {
	return context.WithValue(ctx, ValueKey, bhCtx)
}

// RequestID returns the request ID of the HTTP request
func RequestID(request *http.Request) string {
	return FromRequest(request).RequestID
}

// SetRequestContext sets the given BloodHound Context pointer into the request's context. The resulting, new request pointer
// is then returned.
func SetRequestContext(request *http.Request, bhCtx *Context) *http.Request {
	newRequestContext := context.WithValue(request.Context(), ValueKey, bhCtx)
	return request.WithContext(newRequestContext)
}

const (
	ErrAuthContextInvalid = errors.Error("auth context is invalid")
)

// Audit Log Reference

// func NewAuditLogFromContext(ctx Context, idResolver auth.IdentityResolver) (model.AuditLog, error) {
// 	if ctx.AuditCtx.Model == nil {
// 		return model.AuditLog{}, fmt.Errorf("model cannot be nil when creating a new audit log")
// 	} else if ctx.AuditCtx.Action != model.AuditStatusFailure && ctx.AuditCtx.Action != model.AuditStatusSuccess {
// 		return model.AuditLog{}, fmt.Errorf("invalid action specified in audit log: %s", ctx.AuditCtx.Action)
// 	}
// 	authContext := ctx.AuthCtx

// 	if !authContext.Authenticated() {
// 		return model.AuditLog{}, ErrAuthContextInvalid
// 	} else if identity, err := idResolver.GetIdentity(ctx.AuthCtx); err != nil {
// 		return model.AuditLog{}, ErrAuthContextInvalid
// 	} else {
// 		auditLog := model.AuditLog{
// 			ActorID:    identity.ID.String(),
// 			ActorName:  identity.Name,
// 			ActorEmail: identity.Email,
// 			Action:     ctx.AuditCtx.Action,
// 			Fields:     types.JSONUntypedObject(ctx.AuditCtx.Model.AuditData()),
// 			RequestID:  ctx.RequestID,
// 			Source:     ctx.RequestIP,
// 			Status:     ctx.AuditCtx.Status,
// 		}

// 		if auditLog.Status == model.AuditStatusFailure {
// 			auditLog.Fields["error"] = ctx.AuditCtx.ErrorMsg
// 		}

// 		return auditLog, nil
// 	}
// }
