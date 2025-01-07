package node

import (
	"fmt"
	"github.com/simplecontainer/smr/pkg/configuration"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms/engines/docker"
	"github.com/simplecontainer/smr/pkg/static"
	"github.com/spf13/viper"
	"os"
	"time"
)

func StopDocker(config *configuration.Configuration, definition *v1.ContainerDefinition) {
	agent, err := docker.New(viper.GetString("agent"), config, definition)

	if err != nil {
		fmt.Println(err)
	}

	agent.Stop(static.SIGTERM)
	err = agent.Rename(fmt.Sprintf("%s-%s", agent.Name, agent.DockerID))

	if err != nil {
		fmt.Println(err)
	}

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
			fmt.Println("The smr node controller exited with error!")
			os.Exit(agentInspect.State.ExitCode)
		} else {
			fmt.Println("The smr node controller exited with success!")
		}
	}

	fmt.Println("container is stopped and renamed for backup purpose")
	fmt.Println(fmt.Sprintf("New name: %s-%s", agent.Name, agent.DockerID))
}
