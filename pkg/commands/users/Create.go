package users

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
)

func Create(context *context.Context, username string, domain string, externalIP string) {
	response := network.SendPost(context.Client, fmt.Sprintf("%s/api/v1/user/%s/%s/%s", context.ApiURL, username, domain, externalIP), nil)

	if response.Explanation != "" {
		fmt.Println(response.Explanation)
	}

	if response.Error {
		fmt.Println(response.ErrorExplanation)
	}
}
