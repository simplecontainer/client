package apply

import (
	"fmt"
	"github.com/qdnqn/smr-client/pkg/context"
	"github.com/qdnqn/smr-client/pkg/network"
)

func Apply(context *context.Context, jsonData string) {
	response := network.SendFile(context.Client, fmt.Sprintf("%s/api/v1/apply", context.ApiURL), jsonData)

	fmt.Println(response.Explanation)

	if response.Error {
		fmt.Println(response.ErrorExplanation)
	}
}
