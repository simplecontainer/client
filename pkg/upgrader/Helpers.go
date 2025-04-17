package upgrader

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/simplecontainer/client/pkg/cluster"
	"github.com/simplecontainer/client/pkg/commands/cluster/upgrade"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/node"
	"github.com/simplecontainer/smr/pkg/controler"
)

func handleDrain(mgr *manager.Manager) error {
	currentNode, err := node.New(mgr.Configuration.Node, mgr.Configuration)

	if err != nil {
		return fmt.Errorf("failed to create node for drain: %w", err)
	}

	currentNode.Stop()

	return nil
}

func handleUpgrade(mgr *manager.Manager, ctrl controler.Control) error {
	mgr.Configuration.Args = "start"

	n1, err := node.New(mgr.Configuration.Node, mgr.Configuration)

	if err != nil {
		return fmt.Errorf("failed to initialize current node: %w", err)
	}

	mgr.Configuration.Image = ctrl.Upgrade.Image
	mgr.Configuration.Tag = ctrl.Upgrade.Tag

	n2, err := node.New(mgr.Configuration.Node, mgr.Configuration)

	if err != nil {
		return fmt.Errorf("failed to initialize upgraded node: %w", err)
	}

	if err := upgrade.Upgrader(n1, n2); err != nil {
		return fmt.Errorf("upgrade failed: %w", err)
	}

	glog.Info("Node started again — attempting to rejoin cluster after health check")

	if err := mgr.Context.Connect(true); err != nil {
		return fmt.Errorf("failed to reconnect: %w", err)
	}

	cluster.ReJoin(mgr)
	return nil
}
