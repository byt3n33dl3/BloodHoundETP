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
// Source: github.com/specterops/bloodhound/src/database (interfaces: AuthContextInitializer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	auth "github.com/specterops/bloodhound/src/auth"
	model "github.com/specterops/bloodhound/src/model"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthContextInitializer is a mock of AuthContextInitializer interface.
type MockAuthContextInitializer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthContextInitializerMockRecorder
}

// MockAuthContextInitializerMockRecorder is the mock recorder for MockAuthContextInitializer.
type MockAuthContextInitializerMockRecorder struct {
	mock *MockAuthContextInitializer
}

// NewMockAuthContextInitializer creates a new mock instance.
func NewMockAuthContextInitializer(ctrl *gomock.Controller) *MockAuthContextInitializer {
	mock := &MockAuthContextInitializer{ctrl: ctrl}
	mock.recorder = &MockAuthContextInitializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthContextInitializer) EXPECT() *MockAuthContextInitializerMockRecorder {
	return m.recorder
}

// InitContextFromToken mocks base method.
func (m *MockAuthContextInitializer) InitContextFromToken(arg0 context.Context, arg1 model.AuthToken) (auth.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitContextFromToken", arg0, arg1)
	ret0, _ := ret[0].(auth.Context)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitContextFromToken indicates an expected call of InitContextFromToken.
func (mr *MockAuthContextInitializerMockRecorder) InitContextFromToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitContextFromToken", reflect.TypeOf((*MockAuthContextInitializer)(nil).InitContextFromToken), arg0, arg1)
}
