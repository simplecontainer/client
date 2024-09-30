package version

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/network"
)

func Version(version string, context *context.Context) {
	response := network.SendGet(context.Client, fmt.Sprintf("%s/version", context.ApiURL))

	fmt.Print(fmt.Sprintf("Server version: %s", response.Data.(map[string]any)["ServerVersion"]))
	fmt.Print(fmt.Sprintf("Client version: %s", version))
}
