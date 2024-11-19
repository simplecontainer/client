package node

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/network"
)

func Get(mgr *manager.Manager) {
	response := network.SendGet(mgr.Context.Client, fmt.Sprintf("%s/cluster", mgr.Context.ApiURL))
	fmt.Println(response)
}
