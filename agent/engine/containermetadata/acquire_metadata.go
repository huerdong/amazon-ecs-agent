// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package containermetadata

import (
	"github.com/aws/amazon-ecs-agent/agent/api"
	"github.com/aws/amazon-ecs-agent/agent/config"

	"github.com/cihub/seelog"
	docker "github.com/fsouza/go-dockerclient"
)

// acquireNetworkMetadata parses the docker.NetworkSettings struct and
// packages the desired metadata for JSON marshaling
// Since we accept incomplete metadata fields, we should not return
// errors here and handle them at this stage.
func acquireNetworkMetadata(settings *docker.NetworkSettings) NetworkMetadata {
	if settings == nil {
		return NetworkMetadata{}
	}

	// This metadata is available in two different places in NetworkSettings
	ipv4Address := settings.IPAddress
	ipv4Gateway := settings.Gateway

	// Network mode is not available for Docker API versions 1.17-1.20
	networkModeFromContainer := ""
	if len(settings.Networks) == 1 {
		for modeFromSettings, containerNetwork := range settings.Networks {
			networkModeFromContainer = modeFromSettings
			ipv4Address = containerNetwork.IPAddress
			ipv4Gateway = containerNetwork.Gateway
		}
	}

	networkMD := NetworkMetadata{
		networkMode: networkModeFromContainer,
		ipv4Address: ipv4Address,
		ipv4Gateway: ipv4Gateway,
	}
	return networkMD
}

// acquireDockerContainerMetadata parses the metadata in a docker container
// and packages this data for JSON marshaling
// Since we accept incomplete metadata fields, we should not return
// errors here and handle them at this stage.
func acquireDockerContainerMetadata(container *docker.Container) DockerContainerMD {
	if container == nil {
		seelog.Warn("Failed to acquire container metadata: Container metadata not available or does not exist")
		return DockerContainerMD{}
	}

	imageNameFromConfig := ""
	// In most cases a container should never lack a config but we check regardless to avoid
	// nil pointer exceptions (Could occur if there is some error in the docker api call, if the
	// container we receive has incomplete information)
	if container.Config != nil {
		imageNameFromConfig = container.Config.Image
	} else {
		seelog.Warn("Failed to acquire container metadata: container has no configuration")
		return DockerContainerMD{}
	}

	// Get Port bindings from docker configurations
	var ports []api.PortBinding
	var err error
	// In most cases a container should never lack a host config but we check just in case to avoid
	// nil pointer exception (Could occur if there is some error in the docker api call)
	if container.HostConfig != nil {
		ports, err = api.PortBindingFromDockerPortBinding(container.HostConfig.PortBindings)
		if err != nil {
			seelog.Warnf("Failed to acquire port binding metadata: %v", err)
		}
	} else {
		seelog.Warn("Failed to acquire container metadata: container has no host configuration")
		return DockerContainerMD{}
	}

	networkMD := acquireNetworkMetadata(container.NetworkSettings)

	dockerContainerMD := DockerContainerMD{
		containerID:   container.ID,
		containerName: container.Name,
		imageID:       container.Image,
		imageName:     imageNameFromConfig,
		ports:         ports,
		networkInfo:   networkMD,
	}
	return dockerContainerMD
}

// acquireTaskMetadata parses metadata in the AWS configuration and task
// and packages this data for JSON marshaling
// Since we accept incomplete metadata fields, we should not return
// errors here and handle them at this stage.
func acquireTaskMetadata(cfg *config.Config, task *api.Task) TaskMetadata {
	clusterFromConfig := ""
	if cfg != nil {
		clusterFromConfig = cfg.Cluster
	} else {
		// This error should not happen in most use cases as the agent should not be missing
		// a configuration. This is only here as a safety in tests where the config may not
		// valid to avoid nil pointer dereference
		seelog.Warn("Failed to get cluster Arn: invalid configuration")
	}

	taskArnFromConfig := ""
	if task != nil {
		taskArnFromConfig = task.Arn
	} else {
		// This error should not happen in most use cases. This check is mostly for current or
		// future tests having mocked tasks that are nil
		seelog.Warn("Failed to get task Arn: invalid task")
	}

	return TaskMetadata{
		cluster: clusterFromConfig,
		taskArn: taskArnFromConfig,
	}
}

// acquireMetadata gathers metadata from a docker container, and task
// configuration and data then packages it for JSON Marshaling
// Since we accept incomplete metadata fields, we should not return
// errors here and handle them at this stage.
func (manager *metadataManager) acquireMetadata(createTime string, updateTime string, container *docker.Container, task *api.Task) Metadata {
	taskMD := acquireTaskMetadata(manager.cfg, task)
	dockerMD := acquireDockerContainerMetadata(container)
	return Metadata{
		taskMetadata:            taskMD,
		dockerContainerMetadata: dockerMD,
		containerInstance:       manager.containerInstanceArn,
		createTime:              createTime,
		updateTime:              updateTime,
	}
}

// acquireMetadataAtContainerCreate gathers metadata from task and cluster configurations
// then packages it for JSON Marshaling. We use this version to get data
// available prior to container creation
// Since we accept incomplete metadata fields, we should not return
// errors here and handle them at this stage.
func (manager *metadataManager) acquireMetadataAtContainerCreate(task *api.Task) Metadata {
	taskMD := acquireTaskMetadata(manager.cfg, task)
	return Metadata{
		taskMetadata:      taskMD,
		containerInstance: manager.containerInstanceArn,
	}
}
