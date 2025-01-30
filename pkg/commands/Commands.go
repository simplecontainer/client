package commands

import (
	"fmt"
	"github.com/rodaine/table"
	"github.com/simplecontainer/client/pkg/commands/cli"
	"github.com/simplecontainer/client/pkg/commands/cluster"
	"github.com/simplecontainer/client/pkg/commands/control"
	"github.com/simplecontainer/client/pkg/commands/objects"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

var Commands []contracts.Command

func PreloadCommands() {
	Commands = append(Commands, cli.Context())
	Commands = append(Commands, cli.Users())
	Commands = append(Commands, cli.Version())
	Commands = append(Commands, cli.Platform())

	Commands = append(Commands, objects.Apply())
	Commands = append(Commands, objects.Remove())
	Commands = append(Commands, objects.Ps())
	Commands = append(Commands, objects.Debug())
	Commands = append(Commands, objects.Logs())

	Commands = append(Commands, cluster.Node())

	Commands = append(Commands, control.Secret())
	Commands = append(Commands, control.Container())
	Commands = append(Commands, control.Containers())
	Commands = append(Commands, control.Gitops())
	Commands = append(Commands, control.Configuration())
	Commands = append(Commands, control.Resource())
	Commands = append(Commands, control.CertKey())
	Commands = append(Commands, control.HttpAuth())
}

func Run(mgr *manager.Manager) {
	for _, comm := range Commands {
		for k, arg := range os.Args {
			if comm.GetName() == arg && k == 1 {
				if comm.GetCondition(mgr) {
					for _, fn := range comm.GetDependsOn() {
						fn(mgr, os.Args)
					}

					for _, fn := range comm.GetFunctions() {
						fn(mgr, os.Args)
					}
				}

				return
			}
		}
	}

	tbl := table.New("Command", "Help")

	for _, comm := range Commands {
		tbl.AddRow(comm.GetName(), fmt.Sprintf("smr %s help", comm.GetName()))
	}

	fmt.Print("Available Commands: \n\n")
	tbl.Print()
}
