package definitions

import (
	"encoding/base64"
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
	"log"
)

func Definitions(context *context.Context, definition string) {
	response := network.SendGet(context.Client, fmt.Sprintf("%s/api/v1/definitions/%s", context.ApiURL, definition))

	if definition == "" {
		plugins := response.Data.([]interface{})

		for _, plugin := range plugins {
			fmt.Println(plugin.(string))
		}
	} else {
		if response != nil {
			if response.Data != nil {
				data, err := base64.StdEncoding.DecodeString(response.Data.(string))
				if err != nil {
					log.Fatal("error:", err)
				}

				fmt.Printf("%s\n", data)
			} else {
				fmt.Println("definition is not found")
			}
		}
	}
}
