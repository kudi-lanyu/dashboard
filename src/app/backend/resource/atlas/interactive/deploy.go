package interactive

import (
  client "k8s.io/client-go/kubernetes"
  "log"
  "encoding/json"
)

func DeployApp(spec *DeploySpec, client client.Interface) error {
  data, err := json.Marshal(spec)
  if err != nil {
    log.Println("json marshal failed.", err)
  }
  // print spec info
  log.Println(string(data))
  return err
}

// DeployApp deploys an app based on the given configuration. The app is deployed using the given
// client. App deployment consists of a deployment and an optional service. Both of them
//// share common labels.
//func DeployApp(spec *AppDeploymentSpec, client client.Interface) error {
//  log.Printf("Deploying %s application into %s namespace", spec.Name, spec.Namespace)
//
//  annotations := map[string]string{}
//  if spec.Description != nil {
//    annotations[DescriptionAnnotationKey] = *spec.Description
//  }
//  labels := getLabelsMap(spec.Labels)
//  objectMeta := metaV1.ObjectMeta{
//    Annotations: annotations,
//    Name:        spec.Name,
//    Labels:      labels,
//  }
//
//  containerSpec := api.Container{
//    Name:  spec.Name,
//    Image: spec.ContainerImage,
//    SecurityContext: &api.SecurityContext{
//      Privileged: &spec.RunAsPrivileged,
//    },
//    Resources: api.ResourceRequirements{
//      Requests: make(map[api.ResourceName]resource.Quantity),
//    },
//    Env: convertEnvVarsSpec(spec.Variables),
//  }
//
//  if spec.ContainerCommand != nil {
//    containerSpec.Command = []string{*spec.ContainerCommand}
//  }
//  if spec.ContainerCommandArgs != nil {
//    containerSpec.Args = []string{*spec.ContainerCommandArgs}
//  }
//
//  if spec.CpuRequirement != nil {
//    containerSpec.Resources.Requests[api.ResourceCPU] = *spec.CpuRequirement
//  }
//  if spec.MemoryRequirement != nil {
//    containerSpec.Resources.Requests[api.ResourceMemory] = *spec.MemoryRequirement
//  }
//  podSpec := api.PodSpec{
//    Containers: []api.Container{containerSpec},
//  }
//  if spec.ImagePullSecret != nil {
//    podSpec.ImagePullSecrets = []api.LocalObjectReference{{Name: *spec.ImagePullSecret}}
//  }
//
//  podTemplate := api.PodTemplateSpec{
//    ObjectMeta: objectMeta,
//    Spec:       podSpec,
//  }
//
//  deployment := &apps.Deployment{
//    ObjectMeta: objectMeta,
//    Spec: apps.DeploymentSpec{
//      Replicas: &spec.Replicas,
//      Template: podTemplate,
//      Selector: &metaV1.LabelSelector{
//        // Quoting from https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#selector:
//        // In API version apps/v1beta2, .spec.selector and .metadata.labels no longer default to
//        // .spec.template.metadata.labels if not set. So they must be set explicitly.
//        // Also note that .spec.selector is immutable after creation of the Deployment in apps/v1beta2.
//        MatchLabels: labels,
//      },
//    },
//  }
//  _, err := client.AppsV1beta2().Deployments(spec.Namespace).Create(deployment)
//
//  if err != nil {
//    // TODO(bryk): Roll back created resources in case of error.
//    return err
//  }
//
//  if len(spec.PortMappings) > 0 {
//    service := &api.Service{
//      ObjectMeta: objectMeta,
//      Spec: api.ServiceSpec{
//        Selector: labels,
//      },
//    }
//
//    if spec.IsExternal {
//      service.Spec.Type = api.ServiceTypeLoadBalancer
//    } else {
//      service.Spec.Type = api.ServiceTypeClusterIP
//    }
//
//    for _, portMapping := range spec.PortMappings {
//      servicePort :=
//        api.ServicePort{
//          Protocol: portMapping.Protocol,
//          Port:     portMapping.Port,
//          Name:     generatePortMappingName(portMapping),
//          TargetPort: intstr.IntOrString{
//            Type:   intstr.Int,
//            IntVal: portMapping.TargetPort,
//          },
//        }
//      service.Spec.Ports = append(service.Spec.Ports, servicePort)
//    }
//
//    _, err = client.CoreV1().Services(spec.Namespace).Create(service)
//
//    // TODO(bryk): Roll back created resources in case of error.
//    return err
//  }
//
//  return nil
//}
