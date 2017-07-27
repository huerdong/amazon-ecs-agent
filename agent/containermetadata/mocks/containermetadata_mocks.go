// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/containermetadata (interfaces: MetadataManager)

package mock_containermetadata

import (
	api "github.com/aws/amazon-ecs-agent/agent/api"
	gomock "github.com/golang/mock/gomock"
)

// MockMetadataManager is a mock of MetadataManager interface
type MockMetadataManager struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataManagerMockRecorder
}

// MockMetadataManagerMockRecorder is the mock recorder for MockMetadataManager
type MockMetadataManagerMockRecorder struct {
	mock *MockMetadataManager
}

// NewMockMetadataManager creates a new mock instance
func NewMockMetadataManager(ctrl *gomock.Controller) *MockMetadataManager {
	mock := &MockMetadataManager{ctrl: ctrl}
	mock.recorder = &MockMetadataManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockMetadataManager) EXPECT() *MockMetadataManagerMockRecorder {
	return _m.recorder
}

// CleanTaskMetadata mocks base method
func (_m *MockMetadataManager) CleanTaskMetadata(_param0 *api.Task) error {
	ret := _m.ctrl.Call(_m, "CleanTaskMetadata", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanTaskMetadata indicates an expected call of CleanTaskMetadata
func (_mr *MockMetadataManagerMockRecorder) CleanTaskMetadata(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CleanTaskMetadata", arg0)
}

// CreateMetadata mocks base method
func (_m *MockMetadataManager) CreateMetadata(_param0 []string, _param1 *api.Task, _param2 *api.Container) ([]string, error) {
	ret := _m.ctrl.Call(_m, "CreateMetadata", _param0, _param1, _param2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMetadata indicates an expected call of CreateMetadata
func (_mr *MockMetadataManagerMockRecorder) CreateMetadata(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateMetadata", arg0, arg1, arg2)
}

// SetContainerInstanceArn mocks base method
func (_m *MockMetadataManager) SetContainerInstanceArn(_param0 string) {
	_m.ctrl.Call(_m, "SetContainerInstanceArn", _param0)
}

// SetContainerInstanceArn indicates an expected call of SetContainerInstanceArn
func (_mr *MockMetadataManagerMockRecorder) SetContainerInstanceArn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetContainerInstanceArn", arg0)
}

// UpdateMetadata mocks base method
func (_m *MockMetadataManager) UpdateMetadata(_param0 string, _param1 *api.Task, _param2 *api.Container) error {
	ret := _m.ctrl.Call(_m, "UpdateMetadata", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMetadata indicates an expected call of UpdateMetadata
func (_mr *MockMetadataManagerMockRecorder) UpdateMetadata(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateMetadata", arg0, arg1, arg2)
}
