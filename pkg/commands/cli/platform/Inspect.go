package platform

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/smr/pkg/configuration"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms/engines/docker"
	"github.com/simplecontainer/smr/pkg/static"
	"github.com/spf13/viper"
	"os"
)

func Inspect(config *configuration.Configuration, definition *v1.ContainerDefinition, platform string) {
	var agent platforms.IPlatform
	var err error

	switch platform {
	case static.PLATFORM_DOCKER:
		agent, err = docker.New(viper.GetString("name"), definition)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		inspect, err := docker.DockerInspect(agent.GetGeneratedName())

		bytes, err := json.Marshal(inspect)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(string(bytes))
		break
	default:
		fmt.Println("unsupported platform selected")
		os.Exit(1)
		return
	}
}
