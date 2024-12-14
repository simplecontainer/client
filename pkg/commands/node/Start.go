package node

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/flannel"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/network"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"net/http"
	"os"
)

func Start(mgr *manager.Manager) {
	response := network.SendRequest(mgr.Context.Client, fmt.Sprintf("%s/cluster/start", mgr.Context.ApiURL), http.MethodPost, map[string]any{
		"join": viper.GetString("join"),
		"node": viper.GetString("node"),
	})

	if response.Success {
		fmt.Println(response.Explanation)

		var data map[string]string
		bytes, err := response.Data.MarshalJSON()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = json.Unmarshal(bytes, &data)

		ctx := context.Background()
		err = flannel.Run(ctx, mgr.Context, mgr.Configuration, data["agent"])

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("flannel exited")
	} else {
		fmt.Println(response.ErrorExplanation)
	}
}
