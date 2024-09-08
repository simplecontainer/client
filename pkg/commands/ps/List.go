package ps

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/network"
	"github.com/simplecontainer/smr/implementations/container/container"
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
			c.Run()
		}

		response := network.SendPs(context.Client, fmt.Sprintf("%s/api/v1/ps", context.ApiURL))

		if response == nil {
			return
		}

		keys := make([]string, 0)
		containers := make(map[string]*container.Container)

		for _, group := range response {
			for containerName, container := range group {
				containers[containerName] = container
				keys = append(keys, containerName)
			}
		}

		sort.Strings(keys)

		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("GROUP", "NAME", "DOCKER NAME", "IMAGE", "IP", "PORTS", "DEPS", "DOCKER STATE", "SMR STATE")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for _, k := range keys {
			ips := ""
			ports := ""
			deps := ""

			for _, x := range containers[k].Static.MappingPorts {
				if x.Host != "" {
					ports += fmt.Sprintf("%s:%s, ", x.Host, x.Container)
				} else {
					ports += fmt.Sprintf("%s, ", x.Container)
				}
			}

			networkKeys := make([]string, 0)
			for networkKey, _ := range containers[k].Runtime.Networks {
				networkKeys = append(networkKeys, networkKey)
			}

			sort.Strings(networkKeys)

			for _, u := range networkKeys {
				if containers[k].Runtime.Networks[u].IP != "" {
					ips += fmt.Sprintf("%s (%s), ", containers[k].Runtime.Networks[u].IP, containers[k].Runtime.Networks[u].NetworkName)
				}
			}

			for _, u := range containers[k].Static.Definition.Spec.Container.Dependencies {
				deps += fmt.Sprintf("%s.%s ", u.Group, u.Name)
			}

			lastUpdate := time.Since(containers[k].Status.LastUpdate).Round(time.Second)

			tbl.AddRow(
				helpers.CliRemoveComa(containers[k].Static.Group),
				helpers.CliRemoveComa(containers[k].Static.Name),
				helpers.CliRemoveComa(containers[k].Static.GeneratedName),
				fmt.Sprintf("%s:%s", containers[k].Static.Image, containers[k].Static.Tag),
				helpers.CliRemoveComa(ips),
				helpers.CliRemoveComa(ports),
				helpers.CliRemoveComa(deps),
				containers[k].Runtime.State,
				fmt.Sprintf("%s (%s)", containers[k].Status.GetState(), lastUpdate.String()),
			)

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
