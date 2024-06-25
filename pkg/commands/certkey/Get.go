package certkey

import (
	"encoding/json"
	"fmt"
	"github.com/qdnqn/smr-client/pkg/context"
	"github.com/qdnqn/smr-client/pkg/network"
)

func Get(context *context.Context, group string, identifier string) {
	response := network.SendOperator(context.Client, fmt.Sprintf("%s/api/v1/operators/certkey/Get", context.ApiURL),
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
