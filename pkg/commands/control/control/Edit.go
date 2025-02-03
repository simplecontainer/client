package control

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/simplecontainer/client/pkg/commands/objects/apply"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
)

func Edit(context *context.Context, prefix string, version string, category string, kind string, group string, name string) (json.RawMessage, error) {
	response := network.Send(context.Client, fmt.Sprintf("%s/api/v1/kind/%s/%s/%s/%s/%s/%s", context.ApiURL, prefix, version, category, kind, group, name), http.MethodGet, nil)

	object := json.RawMessage{}

	err := json.Unmarshal(response.Data, &object)

	if err != nil {
		return nil, err
	}

	data, changed, err := helpers.TmpEditor(object)

	if err != nil {
		return nil, err
	}

	if changed {
		response = apply.Apply(context, data)

		if !response.Success {
			return nil, errors.New(response.ErrorExplanation)
		}

		return data, nil
	}

	return nil, errors.New("nothing changed")
}
