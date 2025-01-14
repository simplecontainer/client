package container

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
)

func Describe(context *context.Context) {
	response := network.Send(context.Client, fmt.Sprintf("%s/api/v1/control/container", context.ApiURL), http.MethodGet, nil)
	fmt.Println(string(response.Data))
}
