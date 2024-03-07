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

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/specterops/bloodhound/src/api (interfaces: Authenticator)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"
	time "time"

	uuid "github.com/gofrs/uuid"
	api "github.com/specterops/bloodhound/src/api"
	auth "github.com/specterops/bloodhound/src/auth"
	model "github.com/specterops/bloodhound/src/model"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthenticator is a mock of Authenticator interface.
type MockAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatorMockRecorder
}

// MockAuthenticatorMockRecorder is the mock recorder for MockAuthenticator.
type MockAuthenticatorMockRecorder struct {
	mock *MockAuthenticator
}

// NewMockAuthenticator creates a new mock instance.
func NewMockAuthenticator(ctrl *gomock.Controller) *MockAuthenticator {
	mock := &MockAuthenticator{ctrl: ctrl}
	mock.recorder = &MockAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticator) EXPECT() *MockAuthenticatorMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockAuthenticator) CreateSession(arg0 context.Context, arg1 model.User, arg2 interface{}) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockAuthenticatorMockRecorder) CreateSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockAuthenticator)(nil).CreateSession), arg0, arg1, arg2)
}

// LoginWithSecret mocks base method.
func (m *MockAuthenticator) LoginWithSecret(arg0 context.Context, arg1 api.LoginRequest) (api.LoginDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginWithSecret", arg0, arg1)
	ret0, _ := ret[0].(api.LoginDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginWithSecret indicates an expected call of LoginWithSecret.
func (mr *MockAuthenticatorMockRecorder) LoginWithSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginWithSecret", reflect.TypeOf((*MockAuthenticator)(nil).LoginWithSecret), arg0, arg1)
}

// Logout mocks base method.
func (m *MockAuthenticator) Logout(arg0 model.UserSession) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Logout", arg0)
}

// Logout indicates an expected call of Logout.
func (mr *MockAuthenticatorMockRecorder) Logout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthenticator)(nil).Logout), arg0)
}

// ValidateRequestSignature mocks base method.
func (m *MockAuthenticator) ValidateRequestSignature(arg0 uuid.UUID, arg1 *http.Request, arg2 time.Time) (auth.Context, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRequestSignature", arg0, arg1, arg2)
	ret0, _ := ret[0].(auth.Context)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ValidateRequestSignature indicates an expected call of ValidateRequestSignature.
func (mr *MockAuthenticatorMockRecorder) ValidateRequestSignature(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRequestSignature", reflect.TypeOf((*MockAuthenticator)(nil).ValidateRequestSignature), arg0, arg1, arg2)
}

// ValidateSecret mocks base method.
func (m *MockAuthenticator) ValidateSecret(arg0 context.Context, arg1 string, arg2 model.AuthSecret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateSecret indicates an expected call of ValidateSecret.
func (mr *MockAuthenticatorMockRecorder) ValidateSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateSecret", reflect.TypeOf((*MockAuthenticator)(nil).ValidateSecret), arg0, arg1, arg2)
}

// ValidateSession mocks base method.
func (m *MockAuthenticator) ValidateSession(arg0 string) (auth.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateSession", arg0)
	ret0, _ := ret[0].(auth.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateSession indicates an expected call of ValidateSession.
func (mr *MockAuthenticatorMockRecorder) ValidateSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateSession", reflect.TypeOf((*MockAuthenticator)(nil).ValidateSession), arg0)
}
