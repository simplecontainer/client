package upgrade

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/node"
)

func Upgrade(n *node.Node, image string, tag string) error {
	err := n.Rename(fmt.Sprintf("%s-backup", n.Name))

	if err != nil {
		return err
	}

	err = n.Stop()

	if err != nil {
		return err
	}

	return n.Start()
}
