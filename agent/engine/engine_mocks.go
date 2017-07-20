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
// Source: github.com/aws/amazon-ecs-agent/agent/engine (interfaces: TaskEngine,DockerClient,ImageManager)

package engine

import (
	time "time"

	api "github.com/aws/amazon-ecs-agent/agent/api"
	dockerclient "github.com/aws/amazon-ecs-agent/agent/engine/dockerclient"
	image "github.com/aws/amazon-ecs-agent/agent/engine/image"
	statechange "github.com/aws/amazon-ecs-agent/agent/statechange"
	statemanager "github.com/aws/amazon-ecs-agent/agent/statemanager"
	go_dockerclient "github.com/fsouza/go-dockerclient"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
)

// MockTaskEngine is a mock of TaskEngine interface
type MockTaskEngine struct {
	ctrl     *gomock.Controller
	recorder *MockTaskEngineMockRecorder
}

// MockTaskEngineMockRecorder is the mock recorder for MockTaskEngine
type MockTaskEngineMockRecorder struct {
	mock *MockTaskEngine
}

// NewMockTaskEngine creates a new mock instance
func NewMockTaskEngine(ctrl *gomock.Controller) *MockTaskEngine {
	mock := &MockTaskEngine{ctrl: ctrl}
	mock.recorder = &MockTaskEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockTaskEngine) EXPECT() *MockTaskEngineMockRecorder {
	return _m.recorder
}

// AddTask mocks base method
func (_m *MockTaskEngine) AddTask(_param0 *api.Task) error {
	ret := _m.ctrl.Call(_m, "AddTask", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTask indicates an expected call of AddTask
func (_mr *MockTaskEngineMockRecorder) AddTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddTask", arg0)
}

// Capabilities mocks base method
func (_m *MockTaskEngine) Capabilities() []string {
	ret := _m.ctrl.Call(_m, "Capabilities")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Capabilities indicates an expected call of Capabilities
func (_mr *MockTaskEngineMockRecorder) Capabilities() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Capabilities")
}

// Disable mocks base method
func (_m *MockTaskEngine) Disable() {
	_m.ctrl.Call(_m, "Disable")
}

// Disable indicates an expected call of Disable
func (_mr *MockTaskEngineMockRecorder) Disable() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Disable")
}

// GetTaskByArn mocks base method
func (_m *MockTaskEngine) GetTaskByArn(_param0 string) (*api.Task, bool) {
	ret := _m.ctrl.Call(_m, "GetTaskByArn", _param0)
	ret0, _ := ret[0].(*api.Task)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetTaskByArn indicates an expected call of GetTaskByArn
func (_mr *MockTaskEngineMockRecorder) GetTaskByArn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTaskByArn", arg0)
}

// Init mocks base method
func (_m *MockTaskEngine) Init(_param0 context.Context) error {
	ret := _m.ctrl.Call(_m, "Init", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (_mr *MockTaskEngineMockRecorder) Init(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init", arg0)
}

// ListTasks mocks base method
func (_m *MockTaskEngine) ListTasks() ([]*api.Task, error) {
	ret := _m.ctrl.Call(_m, "ListTasks")
	ret0, _ := ret[0].([]*api.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasks indicates an expected call of ListTasks
func (_mr *MockTaskEngineMockRecorder) ListTasks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListTasks")
}

// MarshalJSON mocks base method
func (_m *MockTaskEngine) MarshalJSON() ([]byte, error) {
	ret := _m.ctrl.Call(_m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON
func (_mr *MockTaskEngineMockRecorder) MarshalJSON() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MarshalJSON")
}

// MustInit mocks base method
func (_m *MockTaskEngine) MustInit(_param0 context.Context) {
	_m.ctrl.Call(_m, "MustInit", _param0)
}

// MustInit indicates an expected call of MustInit
func (_mr *MockTaskEngineMockRecorder) MustInit(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MustInit", arg0)
}

// SetSaver mocks base method
func (_m *MockTaskEngine) SetSaver(_param0 statemanager.Saver) {
	_m.ctrl.Call(_m, "SetSaver", _param0)
}

// SetSaver indicates an expected call of SetSaver
func (_mr *MockTaskEngineMockRecorder) SetSaver(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSaver", arg0)
}

// StateChangeEvents mocks base method
func (_m *MockTaskEngine) StateChangeEvents() <-chan statechange.Event {
	ret := _m.ctrl.Call(_m, "StateChangeEvents")
	ret0, _ := ret[0].(<-chan statechange.Event)
	return ret0
}

// StateChangeEvents indicates an expected call of StateChangeEvents
func (_mr *MockTaskEngineMockRecorder) StateChangeEvents() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StateChangeEvents")
}

// UnmarshalJSON mocks base method
func (_m *MockTaskEngine) UnmarshalJSON(_param0 []byte) error {
	ret := _m.ctrl.Call(_m, "UnmarshalJSON", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnmarshalJSON indicates an expected call of UnmarshalJSON
func (_mr *MockTaskEngineMockRecorder) UnmarshalJSON(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnmarshalJSON", arg0)
}

// Version mocks base method
func (_m *MockTaskEngine) Version() (string, error) {
	ret := _m.ctrl.Call(_m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (_mr *MockTaskEngineMockRecorder) Version() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Version")
}

// MockDockerClient is a mock of DockerClient interface
type MockDockerClient struct {
	ctrl     *gomock.Controller
	recorder *MockDockerClientMockRecorder
}

// MockDockerClientMockRecorder is the mock recorder for MockDockerClient
type MockDockerClientMockRecorder struct {
	mock *MockDockerClient
}

// NewMockDockerClient creates a new mock instance
func NewMockDockerClient(ctrl *gomock.Controller) *MockDockerClient {
	mock := &MockDockerClient{ctrl: ctrl}
	mock.recorder = &MockDockerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockDockerClient) EXPECT() *MockDockerClientMockRecorder {
	return _m.recorder
}

// ClientVersion mocks base method
func (_m *MockDockerClient) ClientVersion() dockerclient.DockerVersion {
	ret := _m.ctrl.Call(_m, "ClientVersion")
	ret0, _ := ret[0].(dockerclient.DockerVersion)
	return ret0
}

// ClientVersion indicates an expected call of ClientVersion
func (_mr *MockDockerClientMockRecorder) ClientVersion() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClientVersion")
}

// ContainerEvents mocks base method
func (_m *MockDockerClient) ContainerEvents(_param0 context.Context) (<-chan DockerContainerChangeEvent, error) {
	ret := _m.ctrl.Call(_m, "ContainerEvents", _param0)
	ret0, _ := ret[0].(<-chan DockerContainerChangeEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerEvents indicates an expected call of ContainerEvents
func (_mr *MockDockerClientMockRecorder) ContainerEvents(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerEvents", arg0)
}

// CreateContainer mocks base method
func (_m *MockDockerClient) CreateContainer(_param0 *go_dockerclient.Config, _param1 *go_dockerclient.HostConfig, _param2 string, _param3 time.Duration) DockerContainerMetadata {
	ret := _m.ctrl.Call(_m, "CreateContainer", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(DockerContainerMetadata)
	return ret0
}

// CreateContainer indicates an expected call of CreateContainer
func (_mr *MockDockerClientMockRecorder) CreateContainer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateContainer", arg0, arg1, arg2, arg3)
}

// DescribeContainer mocks base method
func (_m *MockDockerClient) DescribeContainer(_param0 string) (api.ContainerStatus, DockerContainerMetadata) {
	ret := _m.ctrl.Call(_m, "DescribeContainer", _param0)
	ret0, _ := ret[0].(api.ContainerStatus)
	ret1, _ := ret[1].(DockerContainerMetadata)
	return ret0, ret1
}

// DescribeContainer indicates an expected call of DescribeContainer
func (_mr *MockDockerClientMockRecorder) DescribeContainer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeContainer", arg0)
}

// InspectContainer mocks base method
func (_m *MockDockerClient) InspectContainer(_param0 string, _param1 time.Duration) (*go_dockerclient.Container, error) {
	ret := _m.ctrl.Call(_m, "InspectContainer", _param0, _param1)
	ret0, _ := ret[0].(*go_dockerclient.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectContainer indicates an expected call of InspectContainer
func (_mr *MockDockerClientMockRecorder) InspectContainer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InspectContainer", arg0, arg1)
}

// InspectImage mocks base method
func (_m *MockDockerClient) InspectImage(_param0 string) (*go_dockerclient.Image, error) {
	ret := _m.ctrl.Call(_m, "InspectImage", _param0)
	ret0, _ := ret[0].(*go_dockerclient.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectImage indicates an expected call of InspectImage
func (_mr *MockDockerClientMockRecorder) InspectImage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InspectImage", arg0)
}

// KnownVersions mocks base method
func (_m *MockDockerClient) KnownVersions() []dockerclient.DockerVersion {
	ret := _m.ctrl.Call(_m, "KnownVersions")
	ret0, _ := ret[0].([]dockerclient.DockerVersion)
	return ret0
}

// KnownVersions indicates an expected call of KnownVersions
func (_mr *MockDockerClientMockRecorder) KnownVersions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KnownVersions")
}

// ListContainers mocks base method
func (_m *MockDockerClient) ListContainers(_param0 bool, _param1 time.Duration) ListContainersResponse {
	ret := _m.ctrl.Call(_m, "ListContainers", _param0, _param1)
	ret0, _ := ret[0].(ListContainersResponse)
	return ret0
}

// ListContainers indicates an expected call of ListContainers
func (_mr *MockDockerClientMockRecorder) ListContainers(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListContainers", arg0, arg1)
}

// PullImage mocks base method
func (_m *MockDockerClient) PullImage(_param0 string, _param1 *api.RegistryAuthenticationData) DockerContainerMetadata {
	ret := _m.ctrl.Call(_m, "PullImage", _param0, _param1)
	ret0, _ := ret[0].(DockerContainerMetadata)
	return ret0
}

// PullImage indicates an expected call of PullImage
func (_mr *MockDockerClientMockRecorder) PullImage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PullImage", arg0, arg1)
}

// RemoveContainer mocks base method
func (_m *MockDockerClient) RemoveContainer(_param0 string, _param1 time.Duration) error {
	ret := _m.ctrl.Call(_m, "RemoveContainer", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContainer indicates an expected call of RemoveContainer
func (_mr *MockDockerClientMockRecorder) RemoveContainer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveContainer", arg0, arg1)
}

// RemoveImage mocks base method
func (_m *MockDockerClient) RemoveImage(_param0 string, _param1 time.Duration) error {
	ret := _m.ctrl.Call(_m, "RemoveImage", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveImage indicates an expected call of RemoveImage
func (_mr *MockDockerClientMockRecorder) RemoveImage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveImage", arg0, arg1)
}

// StartContainer mocks base method
func (_m *MockDockerClient) StartContainer(_param0 string, _param1 time.Duration) DockerContainerMetadata {
	ret := _m.ctrl.Call(_m, "StartContainer", _param0, _param1)
	ret0, _ := ret[0].(DockerContainerMetadata)
	return ret0
}

// StartContainer indicates an expected call of StartContainer
func (_mr *MockDockerClientMockRecorder) StartContainer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartContainer", arg0, arg1)
}

// Stats mocks base method
func (_m *MockDockerClient) Stats(_param0 string, _param1 context.Context) (<-chan *go_dockerclient.Stats, error) {
	ret := _m.ctrl.Call(_m, "Stats", _param0, _param1)
	ret0, _ := ret[0].(<-chan *go_dockerclient.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stats indicates an expected call of Stats
func (_mr *MockDockerClientMockRecorder) Stats(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stats", arg0, arg1)
}

// StopContainer mocks base method
func (_m *MockDockerClient) StopContainer(_param0 string, _param1 time.Duration) DockerContainerMetadata {
	ret := _m.ctrl.Call(_m, "StopContainer", _param0, _param1)
	ret0, _ := ret[0].(DockerContainerMetadata)
	return ret0
}

// StopContainer indicates an expected call of StopContainer
func (_mr *MockDockerClientMockRecorder) StopContainer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopContainer", arg0, arg1)
}

// SupportedVersions mocks base method
func (_m *MockDockerClient) SupportedVersions() []dockerclient.DockerVersion {
	ret := _m.ctrl.Call(_m, "SupportedVersions")
	ret0, _ := ret[0].([]dockerclient.DockerVersion)
	return ret0
}

// SupportedVersions indicates an expected call of SupportedVersions
func (_mr *MockDockerClientMockRecorder) SupportedVersions() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SupportedVersions")
}

// Version mocks base method
func (_m *MockDockerClient) Version() (string, error) {
	ret := _m.ctrl.Call(_m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (_mr *MockDockerClientMockRecorder) Version() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Version")
}

// WithVersion mocks base method
func (_m *MockDockerClient) WithVersion(_param0 dockerclient.DockerVersion) DockerClient {
	ret := _m.ctrl.Call(_m, "WithVersion", _param0)
	ret0, _ := ret[0].(DockerClient)
	return ret0
}

// WithVersion indicates an expected call of WithVersion
func (_mr *MockDockerClientMockRecorder) WithVersion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WithVersion", arg0)
}

// MockImageManager is a mock of ImageManager interface
type MockImageManager struct {
	ctrl     *gomock.Controller
	recorder *MockImageManagerMockRecorder
}

// MockImageManagerMockRecorder is the mock recorder for MockImageManager
type MockImageManagerMockRecorder struct {
	mock *MockImageManager
}

// NewMockImageManager creates a new mock instance
func NewMockImageManager(ctrl *gomock.Controller) *MockImageManager {
	mock := &MockImageManager{ctrl: ctrl}
	mock.recorder = &MockImageManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockImageManager) EXPECT() *MockImageManagerMockRecorder {
	return _m.recorder
}

// AddAllImageStates mocks base method
func (_m *MockImageManager) AddAllImageStates(_param0 []*image.ImageState) {
	_m.ctrl.Call(_m, "AddAllImageStates", _param0)
}

// AddAllImageStates indicates an expected call of AddAllImageStates
func (_mr *MockImageManagerMockRecorder) AddAllImageStates(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddAllImageStates", arg0)
}

// GetImageStateFromImageName mocks base method
func (_m *MockImageManager) GetImageStateFromImageName(_param0 string) *image.ImageState {
	ret := _m.ctrl.Call(_m, "GetImageStateFromImageName", _param0)
	ret0, _ := ret[0].(*image.ImageState)
	return ret0
}

// GetImageStateFromImageName indicates an expected call of GetImageStateFromImageName
func (_mr *MockImageManagerMockRecorder) GetImageStateFromImageName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetImageStateFromImageName", arg0)
}

// RecordContainerReference mocks base method
func (_m *MockImageManager) RecordContainerReference(_param0 *api.Container) error {
	ret := _m.ctrl.Call(_m, "RecordContainerReference", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordContainerReference indicates an expected call of RecordContainerReference
func (_mr *MockImageManagerMockRecorder) RecordContainerReference(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RecordContainerReference", arg0)
}

// RemoveContainerReferenceFromImageState mocks base method
func (_m *MockImageManager) RemoveContainerReferenceFromImageState(_param0 *api.Container) error {
	ret := _m.ctrl.Call(_m, "RemoveContainerReferenceFromImageState", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContainerReferenceFromImageState indicates an expected call of RemoveContainerReferenceFromImageState
func (_mr *MockImageManagerMockRecorder) RemoveContainerReferenceFromImageState(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveContainerReferenceFromImageState", arg0)
}

// SetSaver mocks base method
func (_m *MockImageManager) SetSaver(_param0 statemanager.Saver) {
	_m.ctrl.Call(_m, "SetSaver", _param0)
}

// SetSaver indicates an expected call of SetSaver
func (_mr *MockImageManagerMockRecorder) SetSaver(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSaver", arg0)
}

// StartImageCleanupProcess mocks base method
func (_m *MockImageManager) StartImageCleanupProcess(_param0 context.Context) {
	_m.ctrl.Call(_m, "StartImageCleanupProcess", _param0)
}

// StartImageCleanupProcess indicates an expected call of StartImageCleanupProcess
func (_mr *MockImageManagerMockRecorder) StartImageCleanupProcess(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartImageCleanupProcess", arg0)
}
