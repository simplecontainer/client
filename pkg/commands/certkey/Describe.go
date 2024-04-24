package certkey

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
)

func Describe(context *context.Context) {
	response := network.SendOperator(context.Client, fmt.Sprintf("%s/api/v1/operators/certkey", context.ApiURL), nil)

	if !response.Error && len(response.Data) > 0 {
		for _, x := range response.Data["SupportedOperations"].([]interface{}) {
			fmt.Println(x.(string))
		}
	} else {
		fmt.Println(response.Explanation)
	}
}
