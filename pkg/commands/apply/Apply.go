package apply

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
	"net/http"
)

func Apply(context *context.Context, jsonData string) {
	response := network.SendRequest(context.Client, fmt.Sprintf("%s/api/v1/apply", context.ApiURL), http.MethodPost, jsonData)

	if response.Explanation != "" {
		fmt.Println(response.Explanation)
	}

	if response.Error {
		fmt.Println(response.ErrorExplanation)
	}
}
