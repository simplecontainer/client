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
	"net/http"
	"os"
	"os/exec"
	"sort"
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

		response := network.SendRequest(context.Client, fmt.Sprintf("%s/api/v1/ps", context.ApiURL), http.MethodGet, nil)
		if response == nil {
			return
		}

		var display = make(map[string]map[string]ContainerInformation)
		var sortedGroups = make([]string, 0)
		var sortedContainers = make(map[string][]string)
		var states = make(map[string]map[string]any)

		if response.Data != nil {
			bytes, err := response.Data.MarshalJSON()

			if err != nil {
				fmt.Println("invalid response from the server")
				return
			}

			err = json.Unmarshal(bytes, &states)

			if err != nil {
				fmt.Println("invalid response from the server")
				return
			}
		}

		for groupName, group := range states {
			sortedGroups = append(sortedGroups, groupName)

			for containerName, state := range group {
				var container = make(map[string]interface{})

				sortedContainers[groupName] = append(sortedContainers[groupName], containerName)

				var bytes []byte
				var err error

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
						Recreated:     ghost.General.Status.Recreated,
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

					if ghost.Platform.(*docker.Docker).DockerState != "" {
						info.DockerState = fmt.Sprintf("%s (%s)", ghost.Platform.(*docker.Docker).DockerState, static.PLATFORM_DOCKER)
					} else {
						info.DockerState = "-"
					}

					info.LastUpdate = time.Since(ghost.GetStatus().LastUpdate).Round(time.Second)

					info.NodeIP = ghost.General.Runtime.NodeIP
					info.NodeName = ghost.General.Runtime.Agent

					if display[groupName] == nil {
						display[groupName] = make(map[string]ContainerInformation)
					}

					display[groupName][info.GeneratedName] = info
				}
			}
		}

		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("NODE", "GROUP", "NAME", "DOCKER NAME", "IMAGE", "IP", "PORTS", "DEPS", "ENGINE STATE", "SMR STATE")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		sort.Strings(sortedGroups)

		for _, key := range sortedGroups {
			for _, containerName := range sortedContainers[key] {
				container := display[key][containerName]

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
					fmt.Sprintf("%s%s (%s)", container.SmrState, helpers.CliMask(container.Recreated, " (*)", ""), container.LastUpdate),
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
