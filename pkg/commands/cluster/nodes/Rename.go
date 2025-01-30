package nodes

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/node"
)

func Rename(n *node.Node, name string) error {
	err := n.Directory(n.Name, n.Home)

	if err != nil {
		return err
	}

	return n.Rename(fmt.Sprintf("%s", name))
}
