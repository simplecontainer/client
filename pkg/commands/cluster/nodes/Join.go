package nodes

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/flannel"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/network"
	"golang.org/x/net/context"
	"net/http"
	"os"
)

func Join(mgr *manager.Manager) {
	data, err := json.Marshal(map[string]any{
		"join":     mgr.Configuration.Startup.Join,
		"node":     mgr.Configuration.Startup.Node,
		"nodeName": mgr.Configuration.Startup.Name,
		"overlay":  mgr.Configuration.Startup.FlannelCIDR,
		"backend":  mgr.Configuration.Startup.FlannelBackend,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response := network.Send(mgr.Context.Client, fmt.Sprintf("%s/api/v1/cluster/start", mgr.Context.ApiURL), http.MethodPost, data)

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
