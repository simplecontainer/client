package secret

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
)

func Set(context *context.Context, identifier string, value []byte) {
	response := network.Send(context.Client, fmt.Sprintf("%s/api/v1/secrets/propose/secret/%s", context.ApiURL, identifier), http.MethodPost,
		value,
	)

	if response.Success {
		fmt.Println("secret is stored")
	} else {
		fmt.Println("failed to create a secret")
	}
}
