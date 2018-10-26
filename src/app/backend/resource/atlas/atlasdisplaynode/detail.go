package atlasdisplaynode

import (
  "k8s.io/api/core/v1"
  "k8s.io/apimachinery/pkg/fields"
  "k8s.io/client-go/kubernetes"
  metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AtlasNodeInfo struct {
  NodeTotalGpuCount     int64   `json:"nodeTotalGpuCount"`
  NodeAllocatedGpuCount int64   `json:"nodeAllocatedGpuCount"`
  NodeGpuUsage          float64 `json:"nodeGpuUsage"`

  NodeRole  []string `json:"nodeRole"`
  IpAddress string `json:"ipAddress"`
}

func GetNodeInfo(client kubernetes.Interface, nodeName string) *AtlasNodeInfo {
  node, err := client.CoreV1().Nodes().Get(nodeName, metaV1.GetOptions{})
  if err != nil {
    return &AtlasNodeInfo{}
  }
  roles := findNodeRoles(node)
  ipaddress := GetNodeInternalAddress(node)

  allocated, total := CalculateNodeGPU(client, *node)
  usage := getGpuUsage(allocated, total)

  return &AtlasNodeInfo{NodeTotalGpuCount: total,
                        NodeAllocatedGpuCount: allocated,
                        NodeGpuUsage: usage,
                        NodeRole:roles,
                        IpAddress:ipaddress}
}

// NodeInternalAddress . may handle it in frontend is good idea
func GetNodeInternalAddress(node *v1.Node) string {
  address := "unknown"
  if len(node.Status.Addresses) > 0 {
    // address = nodeInfo.node.Status.Addresses[0].Address
    for _, addr := range node.Status.Addresses {
      if addr.Type == v1.NodeInternalIP {
        address = addr.Address
        break
      }
    }
  }
  return address
}

// getPodsFromNode return pod list belong to node
func getPodsFromNode(client kubernetes.Interface, node v1.Node) ([]v1.Pod, error) {
  fieldSelector, err := fields.ParseSelector("spec.nodeName=" + node.Name +
    ",status.phase!=" + string(v1.PodSucceeded) +
    ",status.phase!=" + string(v1.PodFailed))
  if err != nil {
    return nil, err
  }

  podList, err := client.CoreV1().Pods(v1.NamespaceAll).List(metaV1.ListOptions{
    FieldSelector: fieldSelector.String(),
  })
  if err != nil {
    return nil, err
  }
  return podList.Items, err
}
func getPodsByNodeName(client kubernetes.Interface, nodeName string) ([]v1.Pod, error) {
  fieldSelector, err := fields.ParseSelector("spec.nodeName=" + nodeName +
    ",status.phase!=" + string(v1.PodSucceeded) +
    ",status.phase!=" + string(v1.PodFailed))
  if err != nil {
    return nil, err
  }

  podList, err := client.CoreV1().Pods(v1.NamespaceAll).List(metaV1.ListOptions{
    FieldSelector: fieldSelector.String(),
  })
  if err != nil {
    return nil, err
  }
  return podList.Items, err
}
