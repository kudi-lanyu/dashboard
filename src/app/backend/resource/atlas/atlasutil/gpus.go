package atlasutil

import (
	"github.com/kubernetes/dashboard/src/app/backend/resource/atlas/atlaspod"
	"github.com/kubernetes/dashboard/src/app/backend/resource/atlas/common"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

// calculate the GPU count of each node
func CalculateNodeGPU(client kubernetes.Interface, node v1.Node) (totalGPU int64, allocatedGPU int64) {
	totalGPU = GpuInNode(node)

	pods, err := atlaspod.GetAllPodsFromNode(client, node)
	if err != nil {
		return totalGPU, 0

	}

	for _, pod := range pods {
		allocatedGPU += GpuInPod(pod)
	}

	return totalGPU, allocatedGPU
}

func GetGpuUsage(allocatedGpu, totalGpu int64) float64 {
	var gpuUsage float64 = 0
	if totalGpu > 0 {
		gpuUsage = float64(allocatedGpu) / float64(totalGpu) * 100
	}
	return gpuUsage
}

// The way to get GPU Count of Node: nvidia.com/gpu
func GpuInNode(node v1.Node) int64 {
	val, ok := node.Status.Capacity[common.NVIDIAGPUResourceName]

	if !ok {
		return gpuInNodeDeprecated(node)
	}

	return val.Value()
}

// The way to get GPU Count of Node: alpha.kubernetes.io/nvidia-gpu
func gpuInNodeDeprecated(node v1.Node) int64 {
	val, ok := node.Status.Capacity[common.DeprecatedNVIDIAGPUResourceName]

	if !ok {
		return 0
	}

	return val.Value()
}

func GpuInPod(pod v1.Pod) (gpuCount int64) {
	containers := pod.Spec.Containers
	for _, container := range containers {
		gpuCount += gpuInContainer(container)
	}

	return gpuCount
}

func gpuInContainer(container v1.Container) int64 {
	val, ok := container.Resources.Limits[common.NVIDIAGPUResourceName]

	if !ok {
		return gpuInContainerDeprecated(container)
	}

	return val.Value()
}

func gpuInContainerDeprecated(container v1.Container) int64 {
	val, ok := container.Resources.Limits[common.DeprecatedNVIDIAGPUResourceName]

	if !ok {
		return 0
	}

	return val.Value()
}

func GpuInActivePod(pod v1.Pod) (gpuCount int64) {
	if pod.Status.Phase == v1.PodSucceeded || pod.Status.Phase == v1.PodFailed {
		return 0
	}

	containers := pod.Spec.Containers
	for _, container := range containers {
		gpuCount += gpuInContainer(container)
	}

	return gpuCount
}
