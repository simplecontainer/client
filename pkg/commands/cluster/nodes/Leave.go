package nodes

import (
	"github.com/simplecontainer/client/pkg/cluster"
	"github.com/simplecontainer/client/pkg/manager"
)

func Leave(mgr *manager.Manager, nodeID uint64) {
	cluster.Leave(mgr, nodeID)
}
