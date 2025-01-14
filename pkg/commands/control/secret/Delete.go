package secret

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
)

func Delete(context *context.Context, identifier string) {
	response := network.Send(context.Client, fmt.Sprintf("%s/api/v1/secrets/delete/%s", context.ApiURL, identifier), http.MethodDelete, nil)

	if response.Explanation != "" {
		fmt.Println(response.Explanation)
	}

	if response.Error {
		fmt.Println(response.ErrorExplanation)
	}
}
