package version

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
	"net/http"
	"os"
)

func Version(version string, context *context.Context) {
	response := network.SendRequest(context.Client, fmt.Sprintf("%s/version", context.ApiURL), http.MethodGet, nil)

	if response != nil && response.Success {
		bytes, err := response.Data.MarshalJSON()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var ver map[string]string
		err = json.Unmarshal(bytes, &ver)

		fmt.Print(fmt.Sprintf("Server version: %s", ver["ServerVersion"]))
	}

	fmt.Print(fmt.Sprintf("Client version: %s", version))
}
