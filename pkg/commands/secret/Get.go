package secret

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
)

func Get(context *context.Context, identifier string) {
	response := network.SendOperator(context.Client, fmt.Sprintf("%s/api/v1/secrets/get/%s", context.ApiURL, identifier),
		nil,
	)

	bytes, err := json.MarshalIndent(response.Data, "", "  ")

	if err != nil {
		fmt.Println("invalid response sent from the smr-agent")
		return
	}

	if response.Success {
		fmt.Println(string(bytes))
	} else {
		fmt.Println("failed to get a secret")
	}
}