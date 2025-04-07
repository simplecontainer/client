package nodes

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
	"os"
)

func Leave(mgr *manager.Manager) {
	data, err := json.Marshal(map[string]any{
		"join":     mgr.Configuration.Dynamic.Join,
		"node":     mgr.Configuration.Dynamic.API,
		"nodeName": mgr.Configuration.Setup.Node,
		"overlay":  mgr.Configuration.Network.Fcidr,
		"backend":  mgr.Configuration.Network.Fbackend,
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
