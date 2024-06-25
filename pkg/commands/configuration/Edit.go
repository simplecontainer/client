package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/qdnqn/smr-client/pkg/commands/apply"
	"github.com/qdnqn/smr-client/pkg/context"
	"github.com/qdnqn/smr-client/pkg/helpers"
	"github.com/qdnqn/smr-client/pkg/network"
)

func Edit(context *context.Context, group string, identifier string) {
	response := network.SendOperator(context.Client, fmt.Sprintf("%s/api/v1/operators/configuration/Get", context.ApiURL),
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

	data, err := helpers.TmpEditor(bytes)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	apply.Apply(context, string(data))
}
