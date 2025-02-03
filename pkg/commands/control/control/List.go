package control

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
)

func ListKind(context *context.Context, prefix string, version string, category string, kind string) ([]json.RawMessage, error) {
	response := network.Send(context.Client, fmt.Sprintf("%s/api/v1/kind/%s/%s/%s/%s", context.ApiURL, prefix, version, category, kind), http.MethodGet, nil)

	objects := make([]json.RawMessage, 0)

	err := json.Unmarshal(response.Data, &objects)

	if err != nil {
		return nil, err
	}

	return objects, nil
}

func ListKindGroup(context *context.Context, prefix string, version string, category string, kind string, group string) ([]json.RawMessage, error) {
	response := network.Send(context.Client, fmt.Sprintf("%s/api/v1/kind/%s/%s/%s/%s/%s", context.ApiURL, prefix, version, category, kind, group), http.MethodGet, nil)

	objects := make([]json.RawMessage, 0)

	err := json.Unmarshal(response.Data, &objects)

	if err != nil {
		return nil, err
	}

	return objects, nil
}
