package ps

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/network"
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
		var IContainers = make(map[string]map[string]Container)
		var MContainers = make(map[string]map[string]any)

		err := json.Unmarshal(response, &MContainers)

		if err != nil {
			fmt.Println("invalid response from the server")
			return
		}

		err = json.Unmarshal(response, &IContainers)

		if err != nil {
			fmt.Println("invalid response from the server")
			return
		}

		for group, groupContainers := range IContainers {
			for index, containerObj := range groupContainers {
				switch containerObj.Type {
				case static.PLATFORM_DOCKER:
					docker := &ContainerDocker{}

					var bytes []byte
					bytes, err = json.Marshal(MContainers[group][index])

					err = json.Unmarshal(bytes, docker)

					info := ContainerInformation{
						Group:         docker.Platform.Group,
						Name:          docker.Platform.Name,
						GeneratedName: docker.Platform.GeneratedName,
						Image:         docker.Platform.Image,
						Tag:           docker.Platform.Tag,
						IPs:           "",
						Ports:         "",
						Dependencies:  "",
						DockerState:   "",
						SmrState:      containerObj.General.Status.State.State,
					}

					for _, port := range docker.Platform.Ports.Ports {
						if port.Host != "" {
							info.Ports += fmt.Sprintf("%s:%s, ", port.Host, port.Container)
						} else {
							info.Ports += fmt.Sprintf("%s, ", port.Container)
						}
					}

					if info.Ports == "" {
						info.Ports = "-"
					}

					for _, network := range docker.Platform.Networks.Networks {
						if network.Docker.IP != "" {
							info.IPs += fmt.Sprintf("%s (%s), ", network.Docker.IP, network.Reference.Name)
						}
					}

					for _, u := range docker.Platform.Definition.Spec.Container.Dependencies {
						info.Dependencies += fmt.Sprintf("%s.%s ", u.Group, u.Name)
					}

					if info.Dependencies == "" {
						info.Dependencies = "-"
					}

					info.DockerState = fmt.Sprintf("%s (%s)", docker.Platform.DockerState, static.PLATFORM_DOCKER)

					info.LastUpdate = time.Since(containerObj.General.Status.LastUpdate).Round(time.Second)

					display[group] = append(display[group], info)

					break
				}
			}
		}

		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("GROUP", "NAME", "DOCKER NAME", "IMAGE", "IP", "PORTS", "DEPS", "ENGINE STATE", "SMR STATE")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for _, group := range display {
			for _, container := range group {
				tbl.AddRow(
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
