package test

import (
	"github.com/ghodss/yaml"
	"github.com/kubernetes/dashboard/src/app/backend/resource/atlas/atlasjob"
	"github.com/kubernetes/dashboard/src/app/backend/resource/atlas/atlasjob/mpijob"
	"k8s.io/helm/pkg/chartutil"
	"log"
)

// handleFrontValues handle receive arguments from post by frontend code
func HandleFrontValues() (chartutil.Values, error) {
	// transfer data from frontend

	// construct necessary info
	mpiJobArgs := ConstructMpiArgs()

	// transfer the map to chartutil.Values
	data, err := yaml.Marshal(mpiJobArgs)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	helmValues, err := chartutil.ReadValues(data)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	return helmValues, nil
}

func constructMpiTensorboard() atlasjob.TensorboardArgs {
	tensorBoardArgs := atlasjob.TensorboardArgs{}

	tensorBoardArgs.UseTensorboard = false
	//tensorBoardArgs.HostLogPath = ""
	//tensorBoardArgs.IsLocalLogging = false
	tensorBoardArgs.TensorboardImage = "registry.cn-zhangjiakou.aliyuncs.com/tensorflow-samples/tensorflow:1.5.0-devel"
	//tensorBoardArgs.TrainingLogdir = "/training_logs"

	return tensorBoardArgs
}

func constructSyncCodeArgs() atlasjob.SyncCodeArgs {
	syncArgs := atlasjob.SyncCodeArgs{}

	syncArgs.SyncMode = "git"
	syncArgs.SyncSource = ""
	syncArgs.SyncImage = ""

	return syncArgs
}

func constructCommonFlags() atlasjob.JobCommonArgs {

	jobCommonArgs := atlasjob.JobCommonArgs{}

	//jobCommonArgs
	//jobCommonArgs.Mode = "MPIJob"

	jobCommonArgs.Image = ""
	//jobCommonArgs.GPUCount = 0
	//jobCommonArgs.WorkerCount = 1
	jobCommonArgs.Retry = 0

	jobCommonArgs.WorkingDir = "/root"

	// DataDirVolume
	jobCommonArgs.DataDirs = []atlasjob.DataDirVolume{}

	// DataSet
	//if len(atlasjob.DataSet) > 0 {
	//	err := atlasutil.ValidateDatasets(atlasjob.DataSet)
	//	if err != nil {
	//		return err
	//	}
	//	s.DataSet = transformSliceToMap(mpijobdataset, ":")
	//}

	// Env
	//jobCommonArgs.Envs =

	//

	return jobCommonArgs
}

func ConstructMpiArgs() mpijob.MPIJobArgs {
	mpiArgs := mpijob.MPIJobArgs{}

	mpiArgs.Cpu = 5
	mpiArgs.Memory = 100
	//mpiArgs.GPUCount = 10
	//mpiArgs.GPUS = 5
	//
	//mpiArgs.NodeLabels = []string{"l", "node labels"}
	//
	//
	//// Tensorboard
	//mpiArgs.TensorboardArgs = constructMpiTensorboard()
	//
	//// SyncCodeArgs
	//mpiArgs.SyncCodeArgs = constructSyncCodeArgs()
	//
	//// CommonArgs
	//mpiArgs.JobCommonArgs = constructCommonFlags()

	return mpiArgs
}
