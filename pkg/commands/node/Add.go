package node

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/network"
	"github.com/spf13/viper"
	"net/http"
)

func Add(mgr *manager.Manager) {
	response := network.SendRequest(mgr.Context.Client, fmt.Sprintf("%s/cluster/node", mgr.Context.ApiURL), http.MethodPost, map[string]any{
		"node": viper.GetString("node"),
	})

	fmt.Println(response)
}
