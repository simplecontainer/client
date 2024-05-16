package gitops

import (
	"fmt"
	"github.com/qdnqn/smr-client/pkg/context"
	"github.com/qdnqn/smr-client/pkg/network"
)

func Describe(context *context.Context) {
	gitops := network.SendOperator(context.Client, fmt.Sprintf("%s/api/v1/operators/gitops", context.ApiURL), nil)

	for _, x := range gitops.Data["SupportedOperations"].([]interface{}) {
		fmt.Println(x.(string))
	}
}
