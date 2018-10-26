package atlasdisplayjob

import (
  "k8s.io/apimachinery/pkg/api/resource"
)

type AtlasctlSpec struct {
  // Name of the application.
  Name string `json:"name"`

  // Docker image path for the application.
  ContainerImage string `json:"containerImage"`

  // The name of an image pull secret in case of a private docker repository.
  ImagePullSecret *string `json:"imagePullSecret"`

  // Command that is executed instead of container entrypoint, if specified.
  ContainerCommand *string `json:"containerCommand"`

  // Arguments for the specified container command or container entrypoint (if command is not
  // specified here).
  ContainerCommandArgs *string `json:"containerCommandArgs"`

  // Number of replicas of the image to maintain.
  Replicas int32 `json:"replicas"`

  // Target namespace of the application.
  Namespace string `json:"namespace"`

  // Optional memory requirement for the container.
  MemoryRequirement *resource.Quantity `json:"memoryRequirement"`

  // Optional CPU requirement for the container.
  CpuRequirement *resource.Quantity `json:"cpuRequirement"`

  // Labels that will be defined on Pods/RCs/Services
  Labels []Label `json:"labels"`

  // Whether to run the container as privileged user (essentially equivalent to root on the host).
  RunAsPrivileged bool `json:"runAsPrivileged"`
}

type Label struct {
  // Label key
  Key string `json:"key"`

  // Label value
  Value string `json:"value"`
}
