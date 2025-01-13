package secret

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/smr/pkg/network"
	"net/http"
	"strings"
)

func Set(context *context.Context, identifier string, value []byte) {
	response := network.Send(context.Client, fmt.Sprintf("%s/api/v1/secrets/propose/secret/%s", context.ApiURL, identifier), http.MethodPost,
		value,
	)

	bytes, err := json.MarshalIndent(response.Data, "", "  ")

	if err != nil {
		fmt.Println("invalid response sent from the smr-agent")
		return
	}

	if response.Success {
		fmt.Println(strings.Trim(string(bytes), "\""))
	} else {
		fmt.Println("failed to create a secret")
	}
}
