package ps

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/qdnqn/smr-client/pkg/context"
	"github.com/qdnqn/smr-client/pkg/network"
	"github.com/rodaine/table"
	"sort"
)

func Ps(context *context.Context) {
	containers := network.SendPs(context.Client, fmt.Sprintf("%s/api/v1/ps", context.ApiURL))

	if containers == nil {
		return
	}

	keys := make([]string, 0, len(containers))

	for k := range containers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("GROUP", "NAME", "DOCKER NAME", "IMAGE", "IP", "PORTS", "DEPS", "STATE")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, k := range keys {
		for _, v := range containers[k] {
			ips := ""
			ports := ""
			deps := ""
			status := ""

			for _, x := range v.Static.MappingPorts {
				ports += fmt.Sprintf("%s:%s ", x.Host, x.Container)
			}

			for _, u := range v.Runtime.Networks {
				ips += fmt.Sprintf("%s(%s) ", u.IP, u.NetworkName)
			}

			for _, u := range v.Static.Definition.Spec.Container.Dependencies {
				deps += fmt.Sprintf("%s ", u.Name)
			}

			if v.Status.DependsSolved {
				status += fmt.Sprintf("%s ", "Healthy")
			} else {
				status += fmt.Sprintf("%s ", "Starting")
			}

			tbl.AddRow(
				v.Static.Group,
				v.Static.Name,
				v.Static.GeneratedName,
				fmt.Sprintf("%s:%s", v.Static.Image, v.Static.Tag),
				ips,
				ports,
				deps,
				status)
		}
	}

	tbl.Print()
}
