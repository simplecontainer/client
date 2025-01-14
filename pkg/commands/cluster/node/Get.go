package node

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
)

func Get(mgr *manager.Manager) {
	response := network.Send(mgr.Context.Client, fmt.Sprintf("%s/api/v1/cluster", mgr.Context.ApiURL), http.MethodGet, nil)
	fmt.Println(response)
}
