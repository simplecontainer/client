package remove

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
)

func Remove(context *context.Context, jsonData []byte) {
	response := network.Send(context.Client, fmt.Sprintf("%s/api/v1/delete", context.ApiURL), http.MethodPost, jsonData)

	if response.Explanation != "" {
		fmt.Println(response.Explanation)
	}

	if response.Error {
		fmt.Println(response.ErrorExplanation)
	}
}
