package nodes

import (
	"github.com/simplecontainer/client/pkg/node"
)

func Stop(n *node.Node) error {
	return n.Stop()
}
