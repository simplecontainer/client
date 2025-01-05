package node

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/smr/pkg/configuration"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms/engines/docker"
	"github.com/spf13/viper"
	"os"
)

func InspectDocker(config *configuration.Configuration, definition *v1.ContainerDefinition) {
	agent, err := docker.New(viper.GetString("agent"), config, definition)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = agent.Get()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	agentInspect, err := docker.DockerInspect(agent.DockerID)

	bytes, err := json.Marshal(agentInspect)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bytes))
}
