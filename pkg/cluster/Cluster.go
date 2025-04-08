package cluster

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/flannel"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/controler"
	"github.com/simplecontainer/smr/pkg/network"
	"golang.org/x/net/context"
	"net/http"
	"os"
)

func Join(mgr *manager.Manager) {
	data, err := json.Marshal(map[string]any{
		"join":     mgr.Configuration.Join,
		"node":     mgr.Configuration.API,
		"nodeName": mgr.Configuration.Node,
		"overlay":  mgr.Configuration.Flannel.CIDR,
		"backend":  mgr.Configuration.Flannel.Backend,
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
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			done := make(chan error, 1)

			go func() {
				err := flannel.Run(ctx, mgr.Context, mgr.Configuration, data["name"])

				if err != nil {
					fmt.Println("flannel error:", err)
				}

				fmt.Println("flannel exited - try to recover")
				done <- err
			}()

			select {
			case <-ctx.Done():
				fmt.Println("context canceled")
			case err := <-done:
				if err != nil {
					fmt.Println("flannel returned with error:", err)
				}
			}
		} else {
			fmt.Println(response.ErrorExplanation)
		}
	} else {
		fmt.Println(response.ErrorExplanation)
	}
}

func Leave(mgr *manager.Manager) {
	data, err := json.Marshal(map[string]any{
		"join":     mgr.Configuration.Join,
		"node":     mgr.Configuration.API,
		"nodeName": mgr.Configuration.Node,
		"overlay":  mgr.Configuration.Flannel.CIDR,
		"backend":  mgr.Configuration.Flannel.Backend,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response := network.Send(mgr.Context.Client, fmt.Sprintf("%s/api/v1/cluster/drain", mgr.Context.ApiURL), http.MethodPost, data)

	if response.Success {
		fmt.Println(response.Explanation)
	} else {
		fmt.Println(response.ErrorExplanation)
	}
}

func Upgrade(mgr *manager.Manager, node uint64, image string, tag string) {
	control := controler.New()
	control.SetDrain(controler.NewDrain(node))
	control.SetUpgrade(controler.NewUpgrade(image, tag))

	data, err := control.ToJSON()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response := network.Send(mgr.Context.Client, fmt.Sprintf("%s/api/v1/cluster/upgrade", mgr.Context.ApiURL), http.MethodPost, data)

	if response.Success {
		fmt.Println(response.Explanation)
	} else {
		fmt.Println(response.ErrorExplanation)
	}
}
