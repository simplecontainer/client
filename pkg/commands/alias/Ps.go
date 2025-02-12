package alias

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/control"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
	"slices"
)

func Ps() contracts.Command {
	return command.Command{
		Name: "ps",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if helpers.GrabArg(1) == "ps" {
					os.Args[1] = "list"
				}

				if len(os.Args) == 2 {
					os.Args = append(os.Args, "container")
				} else {
					if !slices.Contains([]string{"containers", "gitops"}, helpers.GrabArg(2)) {
						os.Args = append(os.Args[:3], os.Args[2:]...)
						os.Args[2] = "container"
					}
				}

				fmt.Println(os.Args)

				comm := control.List()

				for _, fn := range comm.GetFunctions() {
					fn(mgr, os.Args)
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
