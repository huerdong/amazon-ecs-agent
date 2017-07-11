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
	"fmt"

	"github.com/aws/amazon-ecs-agent/agent/api"
	"github.com/aws/amazon-ecs-agent/agent/config"
	"github.com/cihub/seelog"
	docker "github.com/fsouza/go-dockerclient"
)

// acquireNetworkMetadata parses the docker.NetworkSettings struct and
// packages the desired metadata for JSON marshaling
func acquireNetworkMetadata(settings *docker.NetworkSettings) (NetworkMetadata, error) {
	var err error
	if settings == nil {
		err = fmt.Errorf("Failed to acquire network metadata: Network metadata not available or does not exist")
		return NetworkMetadata{}, err
	}

	// Get Port bindings from docker settings
	ports, err := api.PortBindingFromDockerPortBinding(settings.Ports)
	if err != nil {
		ports = nil
	}

	// This metadata is available in two different places in NetworkSettings
	gateway := settings.Gateway
	iPAddress := settings.IPAddress
	iPv6Gateway := settings.IPv6Gateway

	// Assume there is at most one network mode (And if no mode, network is "none")
	networkModeFromContainer := ""
	if len(settings.Networks) == 1 {
		for modeFromSettings, containerNetwork := range settings.Networks {
			networkModeFromContainer = modeFromSettings
			gateway = containerNetwork.Gateway
			iPAddress = containerNetwork.IPAddress
			iPv6Gateway = containerNetwork.IPv6Gateway
		}
	} else if len(settings.Networks) == 0 {
		networkModeFromContainer = "none"
	}
	networkMD := NetworkMetadata{
		ports:       ports,
		networkMode: networkModeFromContainer,
		gateway:     gateway,
		iPAddress:   iPAddress,
		iPv6Gateway: iPv6Gateway,
	}
	return networkMD, err
}

// acquireDockerContainerMetadata parses the metadata in a docker container
// and packages this data for JSON marshaling
func acquireDockerContainerMetadata(container *docker.Container) (DockerContainerMD, error) {
	var err error
	if container == nil {
		err = fmt.Errorf("Failed to acquire container metadata: Container metadata not available or does not exist")
		return DockerContainerMD{}, err
	}

	imageNameFromConfig := ""
	if container.Config != nil {
		imageNameFromConfig = container.Config.Image
	} else {
		// This case should never happen in real world cases because valid containers
		// cannot lack a Config
		err = fmt.Errorf("Failed to acquire container metadata: container has no configuration")
		return DockerContainerMD{}, err
	}

	networkMD, err := acquireNetworkMetadata(container.NetworkSettings)

	dockerContainerMD := DockerContainerMD{
		status:        container.State.StateString(),
		containerID:   container.ID,
		containerName: container.Name,
		imageID:       container.Image,
		imageName:     imageNameFromConfig,
		networkInfo:   networkMD,
	}
	return dockerContainerMD, err
}

// acquireTaskMetadata parses metadata in the AWS configuration and task
// and packages this data for JSON marshaling
func acquireTaskMetadata(client dockerDummyClient, cfg *config.Config, task *api.Task) TaskMetadata {
	// Get docker version from client. May block metadata file updates so the file changes
	// should be made in a goroutine as this does a docker client call
	version, err := client.Version()
	if err != nil {
		version = ""
		seelog.Errorf("Failed to get docker version number: %s", err.Error())
	}

	clusterArnFromConfig := ""
	if cfg != nil {
		clusterArnFromConfig = cfg.Cluster
	} else {
		// This error should not happen in real world use cases (Or occurs somewhere else long
		// before this should matter)
		seelog.Errorf("Failed to get cluster Arn: Invalid configuration")
	}

	taskArnFromConfig := ""
	if task != nil {
		taskArnFromConfig = task.Arn
	} else {
		// This error should not happen in real world use cases (Or occurs somewhere else long
		// before this should matter)
		seelog.Errorf("Failed to get task Arn: Invalid task")
	}

	return TaskMetadata{
		version:    version,
		clusterArn: clusterArnFromConfig,
		taskArn:    taskArnFromConfig,
	}
}

// acquireMetadata gathers metadata from a docker container, and task
// configuration and data then packages it for JSON Marshaling
func acquireMetadata(client dockerDummyClient, container *docker.Container, cfg *config.Config, task *api.Task) Metadata {
	dockerMD, err := acquireDockerContainerMetadata(container)
	if err != nil {
		seelog.Errorf("Error in getting metadata from docker: %s", err.Error())
	}
	taskMD := acquireTaskMetadata(client, cfg, task)
	return Metadata{
		Version:       taskMD.version,
		Status:        dockerMD.status,
		ContainerID:   dockerMD.containerID,
		ContainerName: dockerMD.containerName,
		ImageID:       dockerMD.imageID,
		ImageName:     dockerMD.imageName,
		ClusterArn:    taskMD.clusterArn,
		TaskArn:       taskMD.taskArn,
		Ports:         dockerMD.networkInfo.ports,
		NetworkMode:   dockerMD.networkInfo.networkMode,
		Gateway:       dockerMD.networkInfo.gateway,
		IPAddress:     dockerMD.networkInfo.iPAddress,
		IPv6Gateway:   dockerMD.networkInfo.iPv6Gateway,
	}
}

// acquireMetadataAtContainerCreate gathers metadata from task and cluster configurations
// then packages it for JSON Marshaling. We use this version to get data
// available prior to container creation
func acquireMetadataAtContainerCreate(client dockerDummyClient, cfg *config.Config, task *api.Task) Metadata {
	taskMD := acquireTaskMetadata(client, cfg, task)
	return Metadata{
		Version:    taskMD.version,
		ClusterArn: taskMD.clusterArn,
		TaskArn:    taskMD.taskArn,
	}
}
