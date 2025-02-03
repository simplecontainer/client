package nodes

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/flannel"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/network"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"net/http"
	"os"
)

func Leave(mgr *manager.Manager) {
	data, err := json.Marshal(map[string]any{
		"join":     viper.GetString("join"),
		"node":     viper.GetString("node"),
		"nodeName": viper.GetString(""),
		"overlay":  viper.GetString("fcidr"),
		"backend":  viper.GetString("fbackend"),
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response := network.Send(mgr.Context.Client, fmt.Sprintf("%s/api/v1/cluster/stop", mgr.Context.ApiURL), http.MethodPost, data)

	if response.Success {
		fmt.Println(response.Explanation)

		var bytes []byte
		var data map[string]string

		bytes, err = response.Data.MarshalJSON()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = json.Unmarshal(bytes, &data)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if response.Success {
			for {
				ctx := context.Background()
				err = flannel.Run(ctx, mgr.Context, mgr.Configuration, data["name"])

				if err != nil {
					fmt.Println(err)
				}

				fmt.Println("flannel exited - try to recover")
			}
		} else {
			fmt.Println(response.ErrorExplanation)
		}
	} else {
		fmt.Println(response.ErrorExplanation)
	}
}
