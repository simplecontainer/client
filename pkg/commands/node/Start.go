package node

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/flannel"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/network"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"strconv"
)

func Start(mgr *manager.Manager) {
	fmt.Println(fmt.Sprintf("%s/cluster/start/%s", mgr.Context.ApiURL, strconv.FormatBool(viper.GetBool("join"))))

	response := network.SendPost(mgr.Context.Client, fmt.Sprintf("%s/cluster/start/%s", mgr.Context.ApiURL, strconv.FormatBool(viper.GetBool("join"))), map[string]any{
		"cluster": viper.GetString("cluster"),
		"url":     viper.GetString("url"),
		"node":    viper.GetString("node"),
	})

	if response.Success {
		ctx := context.Background()
		err := flannel.Run(ctx, mgr.Context, mgr.Configuration, response.Data["agent"].(string))

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("flannel exited")
	} else {
		fmt.Println(response.ErrorExplanation)
	}
}
