package certkey

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/apply"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/network"
	"net/http"
)

func Edit(context *context.Context, group string, identifier string) {
	response := network.SendRequest(context.Client, fmt.Sprintf("%s/api/v1/operators/certkey/Get/%s/%s", context.ApiURL, group, identifier), http.MethodPut, nil)
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
