package restore

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
)

func Restore(context *context.Context) {
	response := network.SendOperator(context.Client, fmt.Sprintf("%s/restore", context.ApiURL), nil)

	bytes, err := json.MarshalIndent(response.Data, "", "  ")

	if err != nil {
		fmt.Println("invalid response sent from the smr-agent")
		return
	}

	fmt.Println(string(bytes))
}
