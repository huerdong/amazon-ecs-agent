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

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/utils/oswrapper (interfaces: File,OS)

package mock_oswrapper

import (
	os "os"

	oswrapper "github.com/aws/amazon-ecs-agent/agent/utils/oswrapper"
	gomock "github.com/golang/mock/gomock"
)

// MockFile is a mock of File interface
type MockFile struct {
	ctrl     *gomock.Controller
	recorder *MockFileMockRecorder
}

// MockFileMockRecorder is the mock recorder for MockFile
type MockFileMockRecorder struct {
	mock *MockFile
}

// NewMockFile creates a new mock instance
func NewMockFile(ctrl *gomock.Controller) *MockFile {
	mock := &MockFile{ctrl: ctrl}
	mock.recorder = &MockFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockFile) EXPECT() *MockFileMockRecorder {
	return _m.recorder
}

// Chmod mocks base method
func (_m *MockFile) Chmod(_param0 os.FileMode) error {
	ret := _m.ctrl.Call(_m, "Chmod", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chmod indicates an expected call of Chmod
func (_mr *MockFileMockRecorder) Chmod(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Chmod", arg0)
}

// Close mocks base method
func (_m *MockFile) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockFileMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Name mocks base method
func (_m *MockFile) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (_mr *MockFileMockRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

// Sync mocks base method
func (_m *MockFile) Sync() error {
	ret := _m.ctrl.Call(_m, "Sync")
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync
func (_mr *MockFileMockRecorder) Sync() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Sync")
}

// Write mocks base method
func (_m *MockFile) Write(_param0 []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Write", _param0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (_mr *MockFileMockRecorder) Write(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0)
}

// MockOS is a mock of OS interface
type MockOS struct {
	ctrl     *gomock.Controller
	recorder *MockOSMockRecorder
}

// MockOSMockRecorder is the mock recorder for MockOS
type MockOSMockRecorder struct {
	mock *MockOS
}

// NewMockOS creates a new mock instance
func NewMockOS(ctrl *gomock.Controller) *MockOS {
	mock := &MockOS{ctrl: ctrl}
	mock.recorder = &MockOSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockOS) EXPECT() *MockOSMockRecorder {
	return _m.recorder
}

// Create mocks base method
func (_m *MockOS) Create(_param0 string) (oswrapper.File, error) {
	ret := _m.ctrl.Call(_m, "Create", _param0)
	ret0, _ := ret[0].(oswrapper.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (_mr *MockOSMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0)
}

// IsNotExist mocks base method
func (_m *MockOS) IsNotExist(_param0 error) bool {
	ret := _m.ctrl.Call(_m, "IsNotExist", _param0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNotExist indicates an expected call of IsNotExist
func (_mr *MockOSMockRecorder) IsNotExist(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsNotExist", arg0)
}

// MkdirAll mocks base method
func (_m *MockOS) MkdirAll(_param0 string, _param1 os.FileMode) error {
	ret := _m.ctrl.Call(_m, "MkdirAll", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MkdirAll indicates an expected call of MkdirAll
func (_mr *MockOSMockRecorder) MkdirAll(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MkdirAll", arg0, arg1)
}

// OpenFile mocks base method
func (_m *MockOS) OpenFile(_param0 string, _param1 int, _param2 os.FileMode) (oswrapper.File, error) {
	ret := _m.ctrl.Call(_m, "OpenFile", _param0, _param1, _param2)
	ret0, _ := ret[0].(oswrapper.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenFile indicates an expected call of OpenFile
func (_mr *MockOSMockRecorder) OpenFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OpenFile", arg0, arg1, arg2)
}

// RemoveAll mocks base method
func (_m *MockOS) RemoveAll(_param0 string) error {
	ret := _m.ctrl.Call(_m, "RemoveAll", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAll indicates an expected call of RemoveAll
func (_mr *MockOSMockRecorder) RemoveAll(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveAll", arg0)
}

// Rename mocks base method
func (_m *MockOS) Rename(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "Rename", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rename indicates an expected call of Rename
func (_mr *MockOSMockRecorder) Rename(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rename", arg0, arg1)
}
