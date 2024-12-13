package node

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/network"
	"github.com/spf13/viper"
)

func Add(mgr *manager.Manager) {
	response := network.SendPost(mgr.Context.Client, fmt.Sprintf("%s/cluster/node", mgr.Context.ApiURL), map[string]any{
		"node": viper.GetString("node"),
	})

	fmt.Println(response)
}
