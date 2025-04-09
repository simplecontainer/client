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
	control := controler.New()
	control.SetStart(controler.NewStart(mgr.Configuration.API, mgr.Configuration.Flannel.CIDR, mgr.Configuration.Flannel.Backend))

	data, err := control.ToJSON()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response := network.Send(mgr.Context.Client, fmt.Sprintf("%s/api/v1/cluster/start", mgr.Context.ApiURL), http.MethodPost, data)

	if response.HttpStatus == http.StatusOK {
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

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		done := make(chan error, 1)

		go func() {
			fmt.Println("starting flannel")
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
}

func ReJoin(mgr *manager.Manager) {
	control := controler.New()
	control.SetStart(controler.NewStart(mgr.Configuration.API, mgr.Configuration.Flannel.CIDR, mgr.Configuration.Flannel.Backend))

	data, err := control.ToJSON()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response := network.Send(mgr.Context.Client, fmt.Sprintf("%s/api/v1/cluster/start", mgr.Context.ApiURL), http.MethodPost, data)

	if response.HttpStatus == http.StatusOK {
		fmt.Println(response.Explanation)
	} else {
		fmt.Println(response.ErrorExplanation)
	}
}

func Leave(mgr *manager.Manager, node uint64) {
	control := controler.New()
	control.SetDrain(controler.NewDrain(node))
	control.SetUpgrade(nil)

	data, err := control.ToJSON()

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
	control.SetStart(nil)
	control.Time()

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
