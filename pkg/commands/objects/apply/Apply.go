package apply

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/smr/pkg/contracts"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
)

func Apply(context *context.Context, jsonData []byte) *contracts.Response {
	return network.Send(context.Client, fmt.Sprintf("%s/api/v1/propose/apply", context.ApiURL), http.MethodPost, jsonData)
}
