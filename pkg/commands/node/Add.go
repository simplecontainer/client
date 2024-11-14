package node

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/network"
	"github.com/spf13/viper"
)

func Add(mgr *manager.Manager) {
	response := network.SendPost(mgr.Context.Client, fmt.Sprintf("%s/cluster/node", mgr.Context.ApiURL, viper.GetString("node")), map[string]any{
		"url":  viper.GetString("url"),
		"node": viper.GetString("node"),
	})

	fmt.Println(response)
}
