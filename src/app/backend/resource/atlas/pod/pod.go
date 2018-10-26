package pod

import (
  "k8s.io/client-go/kubernetes"
  "k8s.io/api/core/v1"
  "k8s.io/apimachinery/pkg/fields"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

  "log"
)

// acquire all active pods from all namespaces
func acquireAllActivePods(client kubernetes.Interface) ([]v1.Pod, error) {
  allPods := []v1.Pod{}

  fieldSelector, err := fields.ParseSelector("status.phase!=" + string(v1.PodSucceeded) + ",status.phase!=" + string(v1.PodFailed))
  if err != nil {
    return allPods,err
  }
  nodeNonTerminatedPodsList, err := client.CoreV1().Pods(metav1.NamespaceAll).List(metav1.ListOptions{FieldSelector: fieldSelector.String()})
  if err!= nil {
    return allPods,nil
  }

  for _, pod := range nodeNonTerminatedPodsList.Items {
    allPods = append(allPods, pod)
  }
  return allPods, nil
}

func acquireAllPods(client kubernetes.Interface) ([]v1.Pod, error) {
  allPods := []v1.Pod{}
  podList, err := client.CoreV1().Pods(metav1.NamespaceAll).List(metav1.ListOptions{})
  if err != nil {
    return allPods, err
  }
  for _, pod := range podList.Items {
    allPods = append(allPods, pod)
  }
  return allPods, nil
}


func GetPodsFromNode(client kubernetes.Interface,node v1.Node) []v1.Pod {
  pods := []v1.Pod{}

  allPods,err := acquireAllActivePods(client)
  if err!= nil {
    log.Println("acquire pod fail.")
    return nil
  }

  for _, pod := range allPods {
    if pod.Spec.NodeName == node.Name {
      pods = append(pods, pod)
    }
  }
  return pods
}
