package nodes

import (
	"github.com/simplecontainer/client/pkg/node"
)

func Run(n *node.Node) error {
	return n.Run()
}
