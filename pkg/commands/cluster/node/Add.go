package node

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/network"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func Add(mgr *manager.Manager) {
	data, err := json.Marshal(map[string]any{
		"node": viper.GetString("node"),
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response := network.Send(mgr.Context.Client, fmt.Sprintf("%s/api/v1/cluster/node", mgr.Context.ApiURL), http.MethodPost, data)

	fmt.Println(response)
}
