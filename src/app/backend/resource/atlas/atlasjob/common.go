// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
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

package atlasjob

import (
	"strings"
)

type EnvVar struct {
	Name  string `yaml:"key";json:"name"`
	Value string `yaml:"value";json:"value"`
}

type Data struct {
	Pvcname  string `yaml:"pvcname";json:"pvcname"`
	Destpath string `yaml:"destpath";json:"destpath"`
}
type (
	JobName  = string
	JobImage = string
	//JobGpuCount = int
	JobEnvs       = []EnvVar
	JobWorkingDir = string
	JobDataDir    = string
	JobCommand    = string
	JobRetry      = int
	JobDataSet    = []Data
	JobDataDirs   = []DataDirVolume
	// DataDirVolume
	JobDataVolumeHostPath      = string
	JobDataVolumeContainerPath = string
	JobDataVolumeName          = string
)

//Just for distribution job because of read params from yaml templatefile
//Base
type JobCommonArgs struct {
	Name  JobName  `yaml:"name"`
	Image JobImage `yaml:"image";json:"image"`
	//GPUCount   JobGpuCount   `yaml:"gpuCount";json:"gpuConunt"`
	Envs       JobEnvs       `yaml:"mpijobenvs";json:"mpijobenvs"`
	WorkingDir JobWorkingDir `yaml:"workingDir";json:"workingDir"`
	Command    JobCommand    `yaml:"command";json:"command"`

	Retry    JobRetry    `yaml:"retry";json:"retry"`
	DataDir  JobDataDir  `yaml:"dataDir";json:"dataDir"` // --dataDir
	DataSet  JobDataSet  `yaml:"DataSet";json:"dataSet"`
	DataDirs JobDataDirs `yaml:"DataDirs";json:"dataDirs"`
}

type DataDirVolume struct {
	HostPath      JobDataVolumeHostPath      `yaml:"hostPath";json:"hostPath"`
	ContainerPath JobDataVolumeContainerPath `yaml:"containerPath";json:"containerPath"`
	Name          JobDataVolumeName          `yaml:"name";json:"name"`
}

func TransformSliceToMap(sets []string, split string) (valuesMap map[string]string) {
	valuesMap = map[string]string{}
	for _, member := range sets {
		splits := strings.SplitN(member, split, 2)
		if len(splits) == 2 {
			valuesMap[splits[0]] = splits[1]
		}
	}

	return valuesMap
}
