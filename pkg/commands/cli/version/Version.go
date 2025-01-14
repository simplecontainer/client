package version

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/context"
)

func Version(version string, context *context.Context) {
	fmt.Print(fmt.Sprintf("Client version: %s", version))
}
