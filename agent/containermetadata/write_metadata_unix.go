// +build !windows

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
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aws/amazon-ecs-agent/agent/api"
	"github.com/aws/amazon-ecs-agent/agent/engine/dockerclient"
)

const (
	mountPoint                     = "/opt/ecs/metadata"
	ContainerMetadataClientVersion = dockerclient.Version_1_21
)

// createMetadataFile initializes the metadata file
func createMetadataFile(metadataDirectoryPath string) error {
	metadataFilePath := filepath.Join(metadataDirectoryPath, metadataFile)
	temp, err := ioutil.TempFile(metadataDirectoryPath, "temp_metadata_file")
	if err != nil {
		return err
	}
	defer temp.Close()
	err = temp.Chmod(readOnlyPerm)
	if err != nil {
		return err
	}
	return os.Rename(temp.Name(), metadataFilePath)
}

// createBinds will do the appropriate formatting to add a new mount in a container's HostConfig
func createBinds(binds []string, dataDirOnHost string, metadataDirectoryPath string, containerName string) []string {
	instanceBind := fmt.Sprintf(`%s/%s:%s/%s`, dataDirOnHost, metadataDirectoryPath, mountPoint, containerName)
	binds = append(binds, instanceBind)
	return binds
}

// injectEnv will add the mount point into the container as an enviornment variable ECS_CONTAINER_METADATA
func injectEnv(env []string, containerName string) []string {
	metadataEnvVariable := fmt.Sprintf("%s=%s/%s", metadataEnvironmentVariable, mountPoint, containerName)
	env = append(env, metadataEnvVariable)
	return env
}

// writeToMetadata puts the metadata into JSON format and writes into
// the metadata file
func writeToMetadataFile(data []byte, task *api.Task, container *api.Container, dataDir string) error {
	metadataFileDir, err := getMetadataFilePath(task, container, dataDir)
	// Boundary case if file path is bad (Such as if task arn is incorrectly formatted)
	if err != nil {
		return fmt.Errorf("write to metadata file for task %s container %s: %v", task, container, err)
	}
	metadataFilePath := filepath.Join(metadataFileDir, metadataFile)

	temp, err := ioutil.TempFile(metadataFileDir, "temp_metadata_file")
	if err != nil {
		return err
	}
	defer temp.Close()
	err = temp.Chmod(readOnlyPerm)
	if err != nil {
		return err
	}
	_, err = temp.Write(data)
	if err != nil {
		return err
	}
	return os.Rename(temp.Name(), metadataFilePath)
}