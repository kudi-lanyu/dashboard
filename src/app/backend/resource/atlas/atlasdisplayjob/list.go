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
	"github.com/kubernetes/dashboard/src/app/backend/api"
	//dashboardapi "github.com/kubernetes/dashboard/src/app/backend/client/api"
	"github.com/kubernetes/dashboard/src/app/backend/resource/atlas/atlasutil"
	"github.com/unisound-ail/atlasctl/pkg/mpi-operator/apis/kubeflow/v1alpha1"
	kubeflowV1alpha1 "github.com/unisound-ail/atlasctl/pkg/mpi-operator/client/clientset/versioned/typed/kubeflow/v1alpha1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

type AtlasMpiJobList struct {
	ListMeta api.ListMeta `json:"listMeta"`
	v1alpha1.MPIJobList
	Errors []error `json:"errors"`
}

type AtlasMpiJob struct {
	v1alpha1.MPIJob
	Trainer      string   `json:"trainer"`
	Pods         []v1.Pod `json:"pods"`     // all the pods including statefulset and job
	ChiefPod     v1.Pod   `json:"chiefPod"` // the chief pod
	RequestedGPU int64    `json:"requestedGpu"`
	AllocatedGPU int64    `json:"allocatedGpu"`
	TrainerType  string   `json:"trainerType"` // return trainer type: TENSORFLOW
}

func NewAtlasMpiJobClient(config *rest.Config, namespace string) (kubeflowV1alpha1.MPIJobInterface, error) {
	KubeflowV1alpha1Client, err := kubeflowV1alpha1.NewForConfig(config)
	if err != nil {
		log.Println("kubeflow client init failed. ", err)
		return nil, err
	}

	mpijobclient := KubeflowV1alpha1Client.MPIJobs(namespace)

	return mpijobclient, err
}

func GetAtlasMpiJob(mpiClient kubeflowV1alpha1.MPIJobInterface, k8sClient kubernetes.Interface, namespace string, jobname string) (*AtlasMpiJob, error) {
	// get mpi job info
	mpiJob, err := mpiClient.Get(jobname, metav1.GetOptions{})
	if err != nil {
		log.Println("mpijob is not found.")
		return nil, err
	}
	atlasmpijob := AtlasMpiJob{}
	atlasmpijob.Kind = mpiJob.Kind
	atlasmpijob.APIVersion = mpiJob.APIVersion
	atlasmpijob.ResourceVersion = mpiJob.ResourceVersion
	atlasmpijob.Namespace = mpiJob.Namespace
	atlasmpijob.Spec = mpiJob.Spec

	atlasmpijob.Status = mpiJob.Status

	atlasmpijob.CreationTimestamp = mpiJob.CreationTimestamp
	atlasmpijob.Annotations = mpiJob.Annotations
	atlasmpijob.ObjectMeta = mpiJob.ObjectMeta
	atlasmpijob.TypeMeta = mpiJob.TypeMeta
	atlasmpijob.OwnerReferences = mpiJob.OwnerReferences
	atlasmpijob.DeletionTimestamp = mpiJob.DeletionTimestamp
	atlasmpijob.Labels = mpiJob.Labels

	atlasmpijob.TrainerType = "mpijob"

	pods, chiefpod, err := getMpiJobPods(k8sClient, mpiJob.Name, namespace)
	if err != nil {
		log.Println("Get Mpi job Pods failed.")
		return nil, err
	} else {
		atlasmpijob.Pods = pods
		atlasmpijob.ChiefPod = chiefpod
	}

	// gpu
	if atlasmpijob.RequestedGPU == 0 {
		for _, pod := range atlasmpijob.Pods {
			atlasmpijob.RequestedGPU += atlasutil.GpuInPod(pod)
		}
	}

	if atlasmpijob.AllocatedGPU == 0 {
		for _, pod := range atlasmpijob.Pods {
			atlasmpijob.AllocatedGPU += atlasutil.GpuInActivePod(pod)
		}
	}
	return &atlasmpijob, err
}

func GetAtlasMpiJobList(mpiClient kubeflowV1alpha1.MPIJobInterface, namespace string) *AtlasMpiJobList {
	mpiJobList, err := mpiClient.List(metav1.ListOptions{})

	totalItems := 0
	if err == nil {
		totalItems = len(mpiJobList.Items)
	}

	atlasMpiJobList := &AtlasMpiJobList{
		ListMeta: api.ListMeta{TotalItems: totalItems},
	}

	atlasMpiJobList.Items = mpiJobList.Items
	atlasMpiJobList.Kind = mpiJobList.Kind
	atlasMpiJobList.APIVersion = mpiJobList.APIVersion
	atlasMpiJobList.ResourceVersion = mpiJobList.ResourceVersion
	atlasMpiJobList.Continue = mpiJobList.Continue

	if err != nil {
		log.Println("mpiJobList get failed.")
		atlasMpiJobList.Errors = append(atlasMpiJobList.Errors, err)
	}
	return atlasMpiJobList
}

func getMpiJobPods(client kubernetes.Interface, jobname, namespace string) (pods []v1.Pod, chiefPod v1.Pod, err error) {
	podList, err := client.CoreV1().Pods(namespace).List(metav1.ListOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ListOptions",
			APIVersion: "v1",
		},
	})

	if err != nil {
		return pods, chiefPod, err
	}

	for _, item := range podList.Items {
		if !isMPIPod(jobname, namespace, item) {
			continue
		}
		if isChiefPod(item) {
			chiefPod = item
		}
		pods = append(pods, item)
	}
	return pods, chiefPod, err
}

func isChiefPod(item v1.Pod) bool {

	if val, ok := item.Labels["mpi_role_type"]; ok && (val == "launcher") {
		log.Println("the mpijob %s with labels %s", item.Name, val)
	} else {
		return false
	}

	return true
}

func isMPIPod(name, ns string, item v1.Pod) bool {
	if val, ok := item.Labels["release"]; ok && (val == name) {
		log.Println("the mpijob %s with labels %s", item.Name, val)
	} else {
		return false
	}

	if val, ok := item.Labels["app"]; ok && (val == "mpijob") {
		log.Println("the mpijob %s with labels %s is found.", item.Name, val)
	} else {
		return false
	}

	if val, ok := item.Labels["group_name"]; ok && (val == "kubeflow.org") {
		log.Println("the mpijob %s with labels %s is found.", item.Name, val)
	} else {
		return false
	}

	if item.Namespace != ns {
		return false
	}
	return true
}

func isMPIJobSucceeded(status v1alpha1.MPIJobStatus) bool {
	// status.MPIJobLauncherStatusType

	return status.LauncherStatus == v1alpha1.LauncherSucceeded
}

func isMPIJobFailed(status v1alpha1.MPIJobStatus) bool {
	return status.LauncherStatus == v1alpha1.LauncherFailed
}

func isMPIJobPending(status v1alpha1.MPIJobStatus) bool {
	return false
}
