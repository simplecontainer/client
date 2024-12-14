package node

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/network"
	"net/http"
)

func Get(mgr *manager.Manager) {
	response := network.SendRequest(mgr.Context.Client, fmt.Sprintf("%s/cluster", mgr.Context.ApiURL), http.MethodGet, nil)
	fmt.Println(response)
}
