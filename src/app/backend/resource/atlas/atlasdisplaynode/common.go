package atlasdisplaynode

const (
	masterLabelRole = "node-role.kubernetes.io/master"

	gangSchdName = "kube-batchd"

	loadBalancerWidth = 16

	// labelNodeRolePrefix is a label prefix for node roles
	// It's copied over to here until it's merged in core: https://github.com/kubernetes/kubernetes/pull/39112
	labelNodeRolePrefix = "node-role.kubernetes.io/"

	// nodeLabelRole specifies the role of a node
	nodeLabelRole = "kubernetes.io/role"
)
