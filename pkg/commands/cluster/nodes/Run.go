package nodes

import (
	"github.com/simplecontainer/client/pkg/node"
)

func Start(n *node.Node) error {
	err := n.Directory(n.Name, n.Home)

	if err != nil {
		return err
	}

	return n.Run()
}
