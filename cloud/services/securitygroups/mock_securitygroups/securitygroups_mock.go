/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network/networkapi (interfaces: SecurityGroupsClientAPI)

// Package mock_securitygroups is a generated GoMock package.
package mock_securitygroups

import (
	context "context"
	reflect "reflect"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-08-01/network"
	gomock "github.com/golang/mock/gomock"
)

// MockSecurityGroupsClientAPI is a mock of SecurityGroupsClientAPI interface
type MockSecurityGroupsClientAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSecurityGroupsClientAPIMockRecorder
}

// MockSecurityGroupsClientAPIMockRecorder is the mock recorder for MockSecurityGroupsClientAPI
type MockSecurityGroupsClientAPIMockRecorder struct {
	mock *MockSecurityGroupsClientAPI
}

// NewMockSecurityGroupsClientAPI creates a new mock instance
func NewMockSecurityGroupsClientAPI(ctrl *gomock.Controller) *MockSecurityGroupsClientAPI {
	mock := &MockSecurityGroupsClientAPI{ctrl: ctrl}
	mock.recorder = &MockSecurityGroupsClientAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecurityGroupsClientAPI) EXPECT() *MockSecurityGroupsClientAPIMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method
func (m *MockSecurityGroupsClientAPI) CreateOrUpdate(arg0 context.Context, arg1, arg2 string, arg3 network.SecurityGroup) (network.SecurityGroupsCreateOrUpdateFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.SecurityGroupsCreateOrUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockSecurityGroupsClientAPIMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockSecurityGroupsClientAPI)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3)
}

// Delete mocks base method
func (m *MockSecurityGroupsClientAPI) Delete(arg0 context.Context, arg1, arg2 string) (network.SecurityGroupsDeleteFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.SecurityGroupsDeleteFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockSecurityGroupsClientAPIMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSecurityGroupsClientAPI)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockSecurityGroupsClientAPI) Get(arg0 context.Context, arg1, arg2, arg3 string) (network.SecurityGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.SecurityGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSecurityGroupsClientAPIMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSecurityGroupsClientAPI)(nil).Get), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockSecurityGroupsClientAPI) List(arg0 context.Context, arg1 string) (network.SecurityGroupListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(network.SecurityGroupListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockSecurityGroupsClientAPIMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecurityGroupsClientAPI)(nil).List), arg0, arg1)
}

// ListAll mocks base method
func (m *MockSecurityGroupsClientAPI) ListAll(arg0 context.Context) (network.SecurityGroupListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.SecurityGroupListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll
func (mr *MockSecurityGroupsClientAPIMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockSecurityGroupsClientAPI)(nil).ListAll), arg0)
}

// UpdateTags mocks base method
func (m *MockSecurityGroupsClientAPI) UpdateTags(arg0 context.Context, arg1, arg2 string, arg3 network.TagsObject) (network.SecurityGroupsUpdateTagsFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTags", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.SecurityGroupsUpdateTagsFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTags indicates an expected call of UpdateTags
func (mr *MockSecurityGroupsClientAPIMockRecorder) UpdateTags(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTags", reflect.TypeOf((*MockSecurityGroupsClientAPI)(nil).UpdateTags), arg0, arg1, arg2, arg3)
}
