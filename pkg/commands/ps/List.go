package ps

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/network"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms/engines/docker"
	"github.com/simplecontainer/smr/pkg/static"
	"os"
	"os/exec"
	"time"
)

func Ps(context *context.Context, watch bool) {
	for {
		if watch {
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			err := c.Run()

			if err != nil {
				return
			}
		}

		response := network.SendPs(context.Client, fmt.Sprintf("%s/api/v1/ps", context.ApiURL))

		if response == nil {
			return
		}

		var display = make(map[string][]ContainerInformation)
		var states = make(map[string]map[string]any)

		err := json.Unmarshal(response, &states)

		if err != nil {
			fmt.Println("invalid response from the server")
			return
		}

		for groupName, group := range states {
			for _, state := range group {
				var container = make(map[string]interface{})

				var bytes []byte
				bytes, err = json.Marshal(state)

				if err != nil {
					continue
				}

				err = json.Unmarshal(bytes, &container)

				switch container["Type"].(string) {
				case static.PLATFORM_DOCKER:
					ghost := &platforms.Container{
						Platform: &docker.Docker{},
						General:  &platforms.General{},
						Type:     static.PLATFORM_DOCKER,
					}

					bytes, err = json.Marshal(container)

					if err != nil {
						continue
					}

					err = json.Unmarshal(bytes, ghost)
					if err != nil {
						continue
					}

					info := ContainerInformation{
						Group:         ghost.Platform.GetGroup(),
						Name:          ghost.Platform.GetName(),
						GeneratedName: ghost.Platform.GetGeneratedName(),
						Image:         ghost.Platform.(*docker.Docker).Image,
						Tag:           ghost.Platform.(*docker.Docker).Tag,
						IPs:           "",
						Ports:         "",
						Dependencies:  "",
						DockerState:   "",
						SmrState:      ghost.General.Status.State.State,
					}

					for _, port := range ghost.Platform.(*docker.Docker).Ports.Ports {
						if port.Host != "" {
							info.Ports += fmt.Sprintf("%s:%s, ", port.Host, port.Container)
						} else {
							info.Ports += fmt.Sprintf("%s, ", port.Container)
						}
					}

					if info.Ports == "" {
						info.Ports = "-"
					}

					for _, network := range ghost.Platform.(*docker.Docker).Networks.Networks {
						if network.Docker.IP != "" {
							info.IPs += fmt.Sprintf("%s (%s), ", network.Docker.IP, network.Reference.Name)
						}
					}

					for _, u := range ghost.Platform.(*docker.Docker).Definition.Spec.Container.Dependencies {
						info.Dependencies += fmt.Sprintf("%s.%s ", u.Group, u.Name)
					}

					if info.Dependencies == "" {
						info.Dependencies = "-"
					}

					info.DockerState = fmt.Sprintf("%s (%s)", ghost.Platform.(*docker.Docker).DockerState, static.PLATFORM_DOCKER)

					info.LastUpdate = time.Since(ghost.GetStatus().LastUpdate).Round(time.Second)

					info.NodeIP = ghost.General.Runtime.NodeIP
					info.NodeName = ghost.General.Runtime.Agent

					display[groupName] = append(display[groupName], info)
				}
			}
		}

		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("NODE", "GROUP", "NAME", "DOCKER NAME", "IMAGE", "IP", "PORTS", "DEPS", "ENGINE STATE", "SMR STATE")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for _, group := range display {
			for _, container := range group {
				tbl.AddRow(
					fmt.Sprintf("%s (%s)", container.NodeName, container.NodeIP),
					helpers.CliRemoveComa(container.Group),
					helpers.CliRemoveComa(container.Name),
					helpers.CliRemoveComa(container.GeneratedName),
					fmt.Sprintf("%s:%s", container.Image, container.Tag),
					helpers.CliRemoveComa(container.IPs),
					helpers.CliRemoveComa(container.Ports),
					helpers.CliRemoveComa(container.Dependencies),
					container.DockerState,
					fmt.Sprintf("%s (%s)", container.SmrState, container.LastUpdate),
				)
			}
		}

		tbl.Print()

		if watch {
			time.Sleep(1 * time.Second)
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
		} else {
			return
		}
	}
}
