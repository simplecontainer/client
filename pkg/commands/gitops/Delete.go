package gitops

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
	"net/http"
)

func Delete(context *context.Context, group string, identifier string) {
	response := network.SendRequest(context.Client, fmt.Sprintf("%s/api/v1/control/gitops/remove/%s/%s", context.ApiURL, group, identifier), http.MethodDelete, nil)

	if response.Explanation != "" {
		fmt.Println(response.Explanation)
	}

	if response.Error {
		fmt.Println(response.ErrorExplanation)
	}
}
