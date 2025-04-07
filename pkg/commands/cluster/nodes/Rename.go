package nodes

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/node"
)

func Rename(n *node.Node, name string) error {
	return n.Rename(fmt.Sprintf("%s", name))
}
