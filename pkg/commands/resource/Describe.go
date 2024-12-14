package resource

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
	"net/http"
)

func Describe(context *context.Context) {
	response := network.SendRequest(context.Client, fmt.Sprintf("%s/api/v1/operators/gitops", context.ApiURL), http.MethodGet, nil)
	fmt.Println(response.Data)
}
