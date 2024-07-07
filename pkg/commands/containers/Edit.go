package containers

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/apply"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/network"
)

func Edit(context *context.Context, group string, identifier string) {
	response := network.SendOperator(context.Client, fmt.Sprintf("%s/api/v1/operators/containers/Get", context.ApiURL),
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

	data, changed, err := helpers.TmpEditor(bytes)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if changed {
		apply.Apply(context, string(data))
	}
}
