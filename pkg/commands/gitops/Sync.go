package gitops

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
)

func Sync(context *context.Context, group string, identifier string) {
	data := map[string]any{
		"group":      group,
		"identifier": identifier,
	}

	response := network.SendOperator(context.Client, fmt.Sprintf("%s/api/v1/operators/gitops/Sync", context.ApiURL), data)

	if response.Explanation != "" {
		fmt.Println(response.Explanation)
	}

	if response.Error {
		fmt.Println(response.ErrorExplanation)
	}
}
