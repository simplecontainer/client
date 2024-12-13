package node

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/flannel"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/network"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

func Start(mgr *manager.Manager) {
	response := network.SendPost(mgr.Context.Client, fmt.Sprintf("%s/cluster/start", mgr.Context.ApiURL), map[string]any{
		"join": viper.GetString("join"),
		"node": viper.GetString("node"),
	})

	if response.Success {
		fmt.Println(response.Explanation)

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
