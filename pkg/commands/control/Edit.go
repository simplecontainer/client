package control

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/control/control"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

func Edit() contracts.Command {
	return command.Command{
		Name: "edit",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				format, err := helpers.BuildFormat(helpers.GrabArg(2), mgr.Configuration.Flags.G)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				get, err := control.Edit(mgr.Context, format.GetPrefix(), format.GetVersion(), format.GetCategory(), format.GetKind(), format.GetGroup(), format.GetName())

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(string(get))
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
