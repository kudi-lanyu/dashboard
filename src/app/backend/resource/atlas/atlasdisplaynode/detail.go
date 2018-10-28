package atlasdisplaynode

import (
	"encoding/json"
	"github.com/kubernetes/dashboard/src/app/backend/resource/atlas/atlasutil"
	"k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

type AtlasNodeInfo struct {
	NodeTotalGpuCount     int64   `json:"nodeTotalGpuCount"`
	NodeAllocatedGpuCount int64   `json:"nodeAllocatedGpuCount"`
	NodeGpuUsage          float64 `json:"nodeGpuUsage"`

	NodeRole  []string `json:"nodeRole"`
	IpAddress string   `json:"ipAddress"`
}

func GetNodeInfo(client kubernetes.Interface, nodeName string) *AtlasNodeInfo {
	node, err := client.CoreV1().Nodes().Get(nodeName, metaV1.GetOptions{})

	log.Println("getNodeInfo:[][]]]]]]]]]]]]]]]]]]]]][[[")
	nodeInfo, _ := json.Marshal(node)
	log.Println(string(nodeInfo))

	if err != nil {
		return &AtlasNodeInfo{}
	}
	roles := findNodeRoles(node)
	ipaddress := GetNodeInternalAddress(node)

	allocated, total := atlasutil.CalculateNodeGPU(client, *node)
	log.Println("allocated: ", allocated, "total:", total)
	usage := atlasutil.GetGpuUsage(allocated, total)
	log.Printf("GPU Usage: %f", usage)

	return &AtlasNodeInfo{NodeTotalGpuCount: total | 200,
		NodeAllocatedGpuCount: allocated | 100,
		NodeGpuUsage:          usage,
		NodeRole:              roles,
		IpAddress:             ipaddress}
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
	log.Println("ipaddress: ", address)
	return address
}
