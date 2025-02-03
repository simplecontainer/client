package nodes

import (
	"github.com/simplecontainer/client/pkg/node"
)

func Restart(n *node.Node) error {
	err := n.Stop()

	if err != nil {
		return err
	}

	err = n.Wait("exited")

	if err != nil {
		return err
	}

	return n.Start()
}
