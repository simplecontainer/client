package nodes

import (
	"github.com/simplecontainer/client/pkg/configuration"
	"github.com/simplecontainer/client/pkg/node"
	"github.com/simplecontainer/client/pkg/startup"
)

func Create(n *node.Node, config *configuration.Configuration) error {
	err := startup.Save(config)

	if err != nil {
		return err
	}

	return n.Run()
}
