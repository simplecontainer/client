package control

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/smr/pkg/kinds/common"
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

	request, err := common.NewRequest(kind)

	if err != nil {
		return nil, err
	}

	err = request.Definition.FromJson(data)

	if err != nil {
		return nil, err
	}

	if changed {
		err = request.ProposeApply(context.Client, context.ApiURL)
		return data, err
	}

	return nil, errors.New("nothing changed")
}
