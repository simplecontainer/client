package secret

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
	"os"
)

func Edit(context *context.Context, identifier string, value string) {
	data, err := json.Marshal(map[string]any{
		"value": value,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	response := network.Send(context.Client, fmt.Sprintf("%s/api/v1/secrets/propose/secret/%s", context.ApiURL, identifier), http.MethodPut, data)

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
