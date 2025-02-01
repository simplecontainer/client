package control

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/control/control"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/static"
	"os"
)

func Remove() contracts.Command {
	return command.Command{
		Name: "get",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				format, err := helpers.BuildFormat(helpers.GrabArg(2), mgr.Configuration.Startup.G)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				err = control.Remove(mgr.Context, format.GetPrefix(), format.GetCategory(), format.GetKind(), format.GetGroup(), format.GetName())

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(static.STATUS_RESPONSE_DELETED)
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
