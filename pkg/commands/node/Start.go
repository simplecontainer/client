package node

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/network"
	"github.com/spf13/viper"
	"strconv"
)

func Start(mgr *manager.Manager) {
	fmt.Println(fmt.Sprintf("%s/cluster/start/%s", mgr.Context.ApiURL, strconv.FormatBool(viper.GetBool("join"))))

	response := network.SendPost(mgr.Context.Client, fmt.Sprintf("%s/cluster/start/%s", mgr.Context.ApiURL, strconv.FormatBool(viper.GetBool("join"))), map[string]any{
		"cluster": viper.GetString("cluster"),
		"url":     viper.GetString("url"),
		"node":    viper.GetString("node"),
	})

	fmt.Println(response)
}
