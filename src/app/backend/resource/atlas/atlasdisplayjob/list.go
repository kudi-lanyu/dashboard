// Copyright 2017 The Kubernetes Authors.
//
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

package atlasdisplayjob

import (
	"log"
  "github.com/kubernetes/dashboard/src/app/backend/api"
  "github.com/kubernetes/dashboard/src/app/backend/errors"
  "github.com/kubernetes/dashboard/src/app/backend/resource/common"
  "github.com/kubernetes/dashboard/src/app/backend/resource/dataselect"
  "github.com/unisound-ail/atlasctl/pkg/mpi-operator/client/clientset/versioned"
  v1alpha1 "github.com/unisound-ail/atlasctl/pkg/mpi-operator/apis/kubeflow/v1alpha1"
  "k8s.io/client-go/kubernetes"
)

// ATlasctlList contains a list of MPI Jobs in the cluster.
type AtlasctlList struct {
	ListMeta api.ListMeta `json:"listMeta"`

	// Unordered list of RbacRoles
	Items []Atlasctl `json:"items"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}

type MPIJobLauncherStatusType string

// These are valid launcher statuses of an MPIJob.
const (
	// LauncherActive means the MPIJob launcher is actively running.
	LauncherActive MPIJobLauncherStatusType = "Active"
	// LauncherSucceeded means the MPIJob launcher has succeeded.
	LauncherSucceeded MPIJobLauncherStatusType = "Succeeded"
	// LauncherFailed means the MPIJob launcher has failed its execution.
	LauncherFailed MPIJobLauncherStatusType = "Failed"
)

type MPIJobStatus struct {
	// Current status of the launcher job.
	// +optional
	LauncherStatus MPIJobLauncherStatusType `json:"launcherStatus,omitempty"`

	// The number of available worker replicas.
	// +optional
	WorkerReplicas int32 `json:"workerReplicas,omitempty"`
}


// RbacRole provides the simplified, combined presentation layer view of Kubernetes' RBAC Roles and ClusterRoles.
// ClusterRoles will be referred to as Roles for the namespace "all namespaces".
type Atlasctl struct {
	ObjectMeta    api.ObjectMeta                   `json:"objectMeta"`
  TypeMeta      api.TypeMeta                     `json:"typeMeta"`
  Spec          AtlasctlSpec                     `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
  trainerType  string                            `json:"trainerType"`
  status       MPIJobLauncherStatusType          `json:"status"`
}



func (mj *Atlasctl) Trainer() string {
	return mj.trainerType
}

// GetRbacRoleList returns a list of all RBAC Roles in the cluster.
func GetAtlasctlList(mpijobclient versioned.Interface, client kubernetes.Interface, dsQuery *dataselect.DataSelectQuery) (*AtlasctlList, error) {
	log.Println("Getting list of MPI Job")
	channels := &common.ResourceChannels{
    AtlasctlList:        common.GetAtlasctlListChannel(mpijobclient, client, 1),

	}
  log.Println("Print channels.AtlasctlList 11: ", channels.AtlasctlList)
	return GetAtlasctlListFromChannels(channels, dsQuery)
}

// GetAtlasctlListFromChannels returns a list of all MPI Job in the cluster reading required resource list once from the channels.
func GetAtlasctlListFromChannels(channels *common.ResourceChannels, dsQuery *dataselect.DataSelectQuery) (*AtlasctlList, error) {
	atlasctls := <-channels.AtlasctlList.List
	err := <-channels.AtlasctlList.Error
	nonCriticalErrors, criticalError := errors.HandleError(err)
	if criticalError != nil {
		return nil, criticalError
  }

	result := toAtlasctlLists(atlasctls.Items, nonCriticalErrors, dsQuery)
	return result, nil
}

func toAtlasctlLists(atlasctls []v1alpha1.MPIJob, nonCriticalErrors []error,
	dsQuery *dataselect.DataSelectQuery) *AtlasctlList {
	result := &AtlasctlList{
		Items:    make([]Atlasctl, 0),
		ListMeta: api.ListMeta{TotalItems: len(atlasctls)},
		Errors:   nonCriticalErrors,
	}

	atlasctlCells, filteredTotal := dataselect.GenericDataSelectWithFilter(toCells(atlasctls), dsQuery)
	atlasctls = fromCells(atlasctlCells)
	result.ListMeta = api.ListMeta{TotalItems: filteredTotal}

	for _, item := range atlasctls {
		result.Items = append(result.Items,
			Atlasctl{
        ObjectMeta:    api.NewObjectMeta(item.ObjectMeta),
        TypeMeta:      api.NewTypeMeta(api.ResourceKindAtlasctl),
        trainerType:       "mpijob",


      })
    log.Printf("Print toAtlasctlLists ######888: %d", item.Spec.GPUs)
    log.Printf("Print toAtlasctlLists ######999: %s", item.Status)
    log.Printf("Print toAtlasctlLists ######000: %v", item.ObjectMeta)

  }
  log.Printf("Print toAtlasctlLists ######666: %+v", atlasctls)
  log.Printf("Print toAtlasctlLists ######777: %+v", result.Items)
  log.Printf("Print toAtlasctlLists ######5555: %v", result.ListMeta)
  log.Printf("Print results ######5566: %v", result)
	return result
}
