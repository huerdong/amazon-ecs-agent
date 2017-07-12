// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-agent/agent/engine/containermetadata (interfaces: MetadataManager)

package mock_containermetadata

import (
	api "github.com/aws/amazon-ecs-agent/agent/api"
	gomock "github.com/golang/mock/gomock"
)

// Mock of MetadataManager interface
type MockMetadataManager struct {
	ctrl     *gomock.Controller
	recorder *_MockMetadataManagerRecorder
}

// Recorder for MockMetadataManager (not exported)
type _MockMetadataManagerRecorder struct {
	mock *MockMetadataManager
}

func NewMockMetadataManager(ctrl *gomock.Controller) *MockMetadataManager {
	mock := &MockMetadataManager{ctrl: ctrl}
	mock.recorder = &_MockMetadataManagerRecorder{mock}
	return mock
}

func (_m *MockMetadataManager) EXPECT() *_MockMetadataManagerRecorder {
	return _m.recorder
}

func (_m *MockMetadataManager) CleanTaskMetadata(_param0 *api.Task) error {
	ret := _m.ctrl.Call(_m, "CleanTaskMetadata", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMetadataManagerRecorder) CleanTaskMetadata(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CleanTaskMetadata", arg0)
}

func (_m *MockMetadataManager) CreateMetadata(_param0 []string, _param1 *api.Task, _param2 *api.Container) ([]string, error) {
	ret := _m.ctrl.Call(_m, "CreateMetadata", _param0, _param1, _param2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMetadataManagerRecorder) CreateMetadata(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateMetadata", arg0, arg1, arg2)
}

func (_m *MockMetadataManager) UpdateMetadata(_param0 string, _param1 *api.Task, _param2 *api.Container) error {
	ret := _m.ctrl.Call(_m, "UpdateMetadata", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMetadataManagerRecorder) UpdateMetadata(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateMetadata", arg0, arg1, arg2)
}