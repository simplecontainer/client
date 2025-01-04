package secret

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
	"net/http"
)

func Edit(context *context.Context, identifier string, value string) {
	response := network.SendRequest(context.Client, fmt.Sprintf("%s/api/v1/secrets/propose/secret/%s", context.ApiURL, identifier), http.MethodPut,
		map[string]any{
			"value": value,
		},
	)

	bytes, err := json.MarshalIndent(response.Data, "", "  ")

	if err != nil {
		fmt.Println("invalid response sent from the smr-agent")
		return
	}

	if response.Success {
		fmt.Println(string(bytes))
	} else {
		fmt.Println("failed to create a secret")
	}
}
