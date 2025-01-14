package node

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/network"
	"github.com/spf13/viper"
	"net/http"
)

func Delete(mgr *manager.Manager) {
	network.Send(mgr.Context.Client, fmt.Sprintf("%s/api/v1/cluster/node/%s", mgr.Context.ApiURL, viper.GetString("node")), http.MethodDelete, nil)
}
