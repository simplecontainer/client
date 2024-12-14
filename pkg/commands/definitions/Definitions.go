package definitions

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
	"net/http"
	"os"
)

func Definitions(context *context.Context, definition string) {
	response := network.SendRequest(context.Client, fmt.Sprintf("%s/api/v1/definitions/%s", context.ApiURL, definition), http.MethodGet, nil)

	if definition == "" {
		bytes, err := response.Data.MarshalJSON()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var plugins []string
		err = json.Unmarshal(bytes, &plugins)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, plugin := range plugins {
			fmt.Println(plugin)
		}
	} else {
		if response != nil {
			if response.Data != nil {
				bytes, err := response.Data.MarshalJSON()

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				var data string
				err = json.Unmarshal(bytes, &data)

				fmt.Printf("%s\n", data)
			} else {
				fmt.Println("definition is not found")
			}
		}
	}
}
