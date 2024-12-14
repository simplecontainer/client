package certkey

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
	"net/http"
)

func Get(context *context.Context, group string, identifier string) {
	response := network.SendRequest(context.Client, fmt.Sprintf("%s/api/v1/control/certkey/Get", context.ApiURL), http.MethodGet,
		map[string]any{
			"group":      group,
			"identifier": identifier,
		},
	)

	bytes, err := json.MarshalIndent(response.Data, "", "  ")

	if err != nil {
		fmt.Println("invalid response sent from the smr-agent")
		return
	}

	fmt.Println(string(bytes))
}
