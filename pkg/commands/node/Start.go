package node

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/network"
	"github.com/spf13/viper"
)

func Start(mgr *manager.Manager) {
	response := network.SendPost(mgr.Context.Client, fmt.Sprintf("%s/cluster/start", mgr.Context.ApiURL), map[string]any{
		"cluster": viper.GetString("cluster"),
		"url":     viper.GetString("url"),
		"node":    viper.GetString("node"),
	})

	fmt.Println(response)
}
