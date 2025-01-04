package node

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/simplecontainer/smr/pkg/configuration"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms/engines/docker"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

func RunDocker(config *configuration.Configuration, definition *v1.ContainerDefinition) {
	agent, err := docker.New(viper.GetString("agent"), config, definition)
	agent.Delete()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		smrDir := fmt.Sprintf("%s/.%s", config.Environment.HOMEDIR, viper.GetString("agent"))
		err = os.MkdirAll(smrDir, 0750)

		if err != nil {
			fmt.Println(err)
			return
		}

		var container *types.Container
		container, err = agent.RunRaw()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Details about starting up node:")
		fmt.Println(fmt.Sprintf("Entrypoint: %s", agent.Entrypoint))
		fmt.Println(fmt.Sprintf("Arguments: %s", agent.Args))

		if viper.GetBool("wait") {
			for c, _ := agent.Get(); c.State != "exited"; c, _ = agent.Get() {
				fmt.Println("Waiting for the node controller to finish....")
				time.Sleep(1 * time.Second)
			}

			agentInspect, err := docker.DockerInspect(agent.DockerID)

			if err != nil {
				fmt.Println(fmt.Sprintf("failed to inspect container %s", agent.DockerID))
				os.Exit(1)
			}

			if agentInspect.State.ExitCode != 0 {
				fmt.Println("The smr node controller finished with error!")
				os.Exit(agentInspect.State.ExitCode)
			} else {
				fmt.Println("The smr node controller finished with success!")
			}
		} else {
			fmt.Println(fmt.Sprintf("Node persistance and configuration location: %s", smrDir))
			fmt.Println(fmt.Sprintf("Container name: %s", container.Names))
			fmt.Println(fmt.Sprintf("Status: %s", container.Status))

			// Implement readiness checking - abort if not ready
		}

		fmt.Println(strings.Repeat("*", 40))
	}
}
