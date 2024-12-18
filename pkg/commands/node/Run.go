package node

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/simplecontainer/smr/pkg/configuration"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms/engines/docker"
	"github.com/spf13/viper"
	"os"
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

		fmt.Println("********************************************************************")
		fmt.Println("Running entrypoint and args")
		fmt.Println(agent.Entrypoint)
		fmt.Println(agent.Args)
		fmt.Println("********************************************************************")

		fmt.Println("The smr agent started with success!")
		fmt.Println(fmt.Sprintf("Node persistance and configuration location: %s", smrDir))
		fmt.Println(container.Names)
		fmt.Println(container.Status)
		fmt.Println("====================================================================")

		if viper.GetBool("wait") {
			for c, _ := agent.Get(); c.State != "exited"; c, _ = agent.Get() {
				time.Sleep(1 * time.Second)
			}
		}
	}
}
