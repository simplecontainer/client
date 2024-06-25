package resource

import (
	"fmt"
	"github.com/qdnqn/smr-client/pkg/context"
	"github.com/qdnqn/smr-client/pkg/network"
)

func Describe(context *context.Context) {
	response := network.SendOperator(context.Client, fmt.Sprintf("%s/api/v1/operators/gitops", context.ApiURL), nil)

	if !response.Error && len(response.Data) > 0 {
		for _, x := range response.Data["SupportedOperations"].([]interface{}) {
			fmt.Println(x.(string))
		}
	} else {
		fmt.Println(response.Explanation)
	}
}
