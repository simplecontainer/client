package gitops

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
)

func Delete(context *context.Context, group string, identifier string) {
	response := network.SendOperator(context.Client, fmt.Sprintf("%s/api/v1/operators/gitops/Delete", context.ApiURL),
		map[string]any{
			"group":      group,
			"identifier": identifier,
		},
	)

	fmt.Println(response.Explanation)

	if response.Error {
		fmt.Println(response.ErrorExplanation)
	}
}
