package remove

import (
	"fmt"
	"github.com/qdnqn/smr-client/pkg/context"
	"github.com/qdnqn/smr-client/pkg/network"
)

func Remove(context *context.Context, jsonData string) {
	response := network.SendFile(context.Client, fmt.Sprintf("%s/api/v1/delete", context.ApiURL), jsonData)

	fmt.Println(response.Explanation)

	if response.Error {
		fmt.Println(response.ErrorExplanation)
	}
}
