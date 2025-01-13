package certkey

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
)

func Get(context *context.Context, group string, identifier string) {
	response := network.Send(context.Client, fmt.Sprintf("%s/api/v1/control/certkey/get", context.ApiURL), http.MethodGet, nil)

	bytes, err := json.MarshalIndent(response.Data, "", "  ")

	if err != nil {
		fmt.Println("invalid response sent from the smr-agent")
		return
	}

	fmt.Println(string(bytes))
}
