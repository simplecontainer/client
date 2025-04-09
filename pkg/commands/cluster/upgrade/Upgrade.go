package upgrade

import (
	"github.com/simplecontainer/client/pkg/cluster"
	"github.com/simplecontainer/client/pkg/manager"
)

func Upgrade(mgr *manager.Manager, node uint64, image string, tag string) {
	cluster.Upgrade(mgr, node, image, tag)
}
