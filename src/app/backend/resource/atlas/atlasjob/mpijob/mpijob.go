package mpijob

import (
	"github.com/kubernetes/dashboard/src/app/backend/resource/atlas/atlasjob"
)

type (
	Gpus   = int
	Cpu    = int
	Memory = int
	//MpiJobNodeLabels = []string
)

type MPIJobArgs struct {
	// necessary
	Name    atlasjob.JobName    `yaml:"name";json:"name"`
	Image   atlasjob.JobImage   `yaml:"image";json:"image"`
	GPUS    Gpus                `yaml:"gpus";json:"gpus"`
	Command atlasjob.JobCommand `yaml:"command";json:"command"`

	// optional
	Cpu    Cpu    `yaml:"cpu";json:"cpu"`
	Memory Memory `yaml:"memory";json:"memory"`

	// for common optional
	Envs       atlasjob.JobEnvs       `yaml:"mpijobenvs";json:"mpijobenvs"`
	WorkingDir atlasjob.JobWorkingDir `yaml:"workingDir";json:"workingDir"`
	Retry      atlasjob.JobRetry      `yaml:"retry";json:"retry"`

	// dashboard base optional should complete
	DataSet  atlasjob.JobDataSet  `yaml:"mpijobdataset";json:"mpijobdataset"`
	DataDirs atlasjob.JobDataDirs `yaml:"dataDirs";json:"dataDirs"`

	// for tensorboard
	// optional
	UseTensorboard atlasjob.UseTensorboard `yaml:"useTensorboard";json:"useTensorboard"`
	// if useTensorboard is true set blow field with default value, else
	HostLogPath    atlasjob.HostLogPath    `yaml:"hostLogPath";json:"hostLogPath"`
	TrainingLogdir atlasjob.TrainingLogdir `yaml:"trainingLogdir";json:"trainingLogdir"`
	// default true
	IsLocalLogging atlasjob.IsLocalLogging `yaml:"isLocalLogging";json:"isLocalLogging"`
	//TensorboardImage atlasjob.TensorboardImage `yaml:"tensorboardImage";json:"tensorboardImage"`

	// for sync up source code
	SyncMode           atlasjob.SyncMode           `yaml:"syncMode";json:"syncMode"`
	SyncSource         atlasjob.SyncSource         `yaml:"syncSource";json:"syncSource"`
	SyncImage          atlasjob.SyncImage          `yaml:"syncImage";json:"syncImage"`
	SyncGitProjectName atlasjob.SyncGitProjectName `yaml:"syncGitProjectName";json:"syncGitProjectName"`
}

// AppDeploymentFromFileResponse is a specification for deployment from file
type MpiJobDeployResponse struct {
	// Name of the file
	Name string `json:"name"`

	// Error after create resource
	Error string `json:"error"`
}
