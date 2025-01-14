package node

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/simplecontainer/smr/pkg/authentication"
	"github.com/simplecontainer/smr/pkg/client"
	"github.com/simplecontainer/smr/pkg/configuration"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
	"github.com/simplecontainer/smr/pkg/dns"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms/engines/docker"
	"github.com/simplecontainer/smr/pkg/static"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

func Run(config *configuration.Configuration, definition *v1.ContainerDefinition, platform string) {
	var agent platforms.IPlatform
	var err error

	switch platform {
	case static.PLATFORM_DOCKER:
		agent, err = docker.New(viper.GetString("agent"), config, definition)
		break
	default:
		fmt.Println("unsupported platform selected")
		os.Exit(1)
		return
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		smrDir := fmt.Sprintf("%s/.%s", config.Environment.HOMEDIR, viper.GetString("agent"))
		err = os.MkdirAll(smrDir, 0750)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		_, err = agent.Run(config, &client.Http{}, &dns.Records{}, &authentication.User{})

		if err != nil {
			if err.Error() != "agent container name is empty" {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		fmt.Println("Details about starting up node:")
		fmt.Println(fmt.Sprintf("Entrypoint: %s", agent.GetDefinition().(*v1.ContainerDefinition).Spec.Container.Entrypoint))
		fmt.Println(fmt.Sprintf("Arguments: %s", agent.GetDefinition().(*v1.ContainerDefinition).Spec.Container.Args))

		if viper.GetBool("wait") {
			switch platform {
			case static.PLATFORM_DOCKER:
				for c, _ := docker.DockerGet(agent.GetGeneratedName()); c.State != "exited"; c, _ = docker.DockerGet(agent.GetGeneratedName()) {
					fmt.Println("Waiting for the node controller to finish....")
					time.Sleep(1 * time.Second)
				}

				var agentInspect types.ContainerJSON
				agentInspect, err = docker.DockerInspect(agent.GetGeneratedName())

				if err != nil {
					fmt.Println(fmt.Sprintf("failed to inspect container %s", agent.GetGeneratedName))
					os.Exit(1)
				}

				if agentInspect.State.ExitCode != 0 {
					fmt.Println("The smr node controller finished with error!")
					os.Exit(agentInspect.State.ExitCode)
				} else {
					fmt.Println("The smr node controller finished with success!")
				}

				id := agent.GetId()
				err = agent.Rename(fmt.Sprintf("%s-%s", agent.GetGeneratedName(), id))

				if err != nil {
					fmt.Println("failed to rename container")
					os.Exit(1)
				}
				break
			default:
				fmt.Println("unsupported platform selected")
				os.Exit(1)
				return
			}
		} else {
			fmt.Println(fmt.Sprintf("Node persistance and configuration location: %s", smrDir))
			fmt.Println(fmt.Sprintf("Container name: %s", agent.GetGeneratedName()))

			// TODO: Implement readiness checking - abort if not ready
		}

		fmt.Println(strings.Repeat("*", 40))
	}
}
