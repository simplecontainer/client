package node

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/network"
	"github.com/spf13/viper"
	"net/http"
)

func Delete(mgr *manager.Manager) {
	network.SendRequest(mgr.Context.Client, fmt.Sprintf("%s/cluster/node/%s", mgr.Context.ApiURL, viper.GetString("node")), http.MethodDelete, nil)
}
