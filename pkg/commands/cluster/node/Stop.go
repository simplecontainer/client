package node

import (
	"fmt"
	"github.com/simplecontainer/smr/pkg/configuration"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms/engines/docker"
	"github.com/simplecontainer/smr/pkg/static"
	"github.com/spf13/viper"
	"os"
	"time"
)

func Stop(config *configuration.Configuration, definition *v1.ContainerDefinition, platform string) {
	var agent platforms.IPlatform
	var err error

	switch platform {
	case static.PLATFORM_DOCKER:
		agent, err = docker.New(viper.GetString("name"), config, definition)
		break
	default:
		fmt.Println("unsupported platform selected")
		os.Exit(1)
		return
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		agent.Stop(static.SIGTERM)

		id := agent.GetId()
		err = agent.Rename(fmt.Sprintf("%s-%s", agent.GetGeneratedName(), id))

		switch platform {
		case static.PLATFORM_DOCKER:
			if viper.GetBool("wait") {
				for c, _ := docker.DockerGet(agent.GetGeneratedName()); c.State != "exited"; c, _ = docker.DockerGet(agent.GetGeneratedName()) {
					fmt.Println("Waiting for the node controller to finish....")
					time.Sleep(1 * time.Second)
				}

				agentInspect, err := docker.DockerInspect(agent.GetGeneratedName())

				if err != nil {
					fmt.Println(fmt.Sprintf("failed to inspect container %s", agent.GetGeneratedName()))
					os.Exit(1)
				}

				if agentInspect.State.ExitCode != 0 {
					fmt.Println("The smr node controller exited with error!")
					os.Exit(agentInspect.State.ExitCode)
				} else {
					fmt.Println("The smr node controller exited with success!")
				}
			}

			break
		default:
			fmt.Println("unsupported platform selected")
			os.Exit(1)
			return
		}

		fmt.Println("container is stopped and renamed for backup purpose")
		fmt.Println(fmt.Sprintf("New name: %s-%d", agent.GetGeneratedName(), id))
	}
}
