package nodes

import (
	"github.com/simplecontainer/client/pkg/cluster"
	"github.com/simplecontainer/client/pkg/manager"
)

func Leave(mgr *manager.Manager) {
	cluster.Leave(mgr)
}
