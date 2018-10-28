package atlasdisplaynode

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"log"
	"strings"
)

// findNodeRoles returns the roles of a given node.
// The roles are determined by looking for:
// * a node-role.kubernetes.io/<role>="" label
// * a kubernetes.io/role="<role>" label
func findNodeRoles(node *v1.Node) []string {
	roles := sets.NewString()
	for k, v := range node.Labels {
		switch {
		case strings.HasPrefix(k, labelNodeRolePrefix):
			if role := strings.TrimPrefix(k, labelNodeRolePrefix); len(role) > 0 {
				roles.Insert(role)
			}

		case k == nodeLabelRole && v != "":
			roles.Insert(v)
		}
	}
	log.Println("role list:", roles.List())
	return roles.List()
}
