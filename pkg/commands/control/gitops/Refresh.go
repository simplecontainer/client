package gitops

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
)

func Refresh(context *context.Context, group string, identifier string) {
	response := network.Send(context.Client, fmt.Sprintf("%s/api/v1/control/gitops/refresh/%s/%s", context.ApiURL, group, identifier), http.MethodGet, nil)

	if response.Explanation != "" {
		fmt.Println(response.Explanation)
	}

	if response.Error {
		fmt.Println(response.ErrorExplanation)
	}
}
