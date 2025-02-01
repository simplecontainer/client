package control

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
)

func Remove(context *context.Context, prefix string, version string, category string, kind string, group string, name string) error {
	response := network.Send(context.Client, fmt.Sprintf("%s/api/v1/kind/%s/%s/%s/%s/%s/%s", context.ApiURL, prefix, version, category, kind, group, name), http.MethodDelete, nil)

	if response.Error {
		return errors.New(response.ErrorExplanation)
	} else {
		return nil
	}
}
