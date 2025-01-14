package control

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/control/container"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

const HELP_CONTAINER string = "Eg: smr configuration [describe, delete, edit, get, list, restart, view]"

func Container() contracts.Command {
	return command.Command{
		Name: "container",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) > 2 {
					switch os.Args[2] {
					case "describe":
						container.Describe(mgr.Context)
						break
					case "list":
						container.List(mgr.Context)
						break
					case "get":
						if len(os.Args) > 4 {
							container.Get(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CONTAINER)
						}
						break
					case "view":
						if len(os.Args) > 4 {
							container.View(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CONTAINER)
						}
						break
					case "edit":
						if len(os.Args) > 4 {
							container.Edit(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CONTAINER)
						}
						break
					case "restart":
						if len(os.Args) > 4 {
							container.Restart(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CONTAINER)
						}
						break
					case "delete":
						if len(os.Args) > 4 {
							container.Delete(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CONTAINER)
						}
						break
					default:
						fmt.Println(HELP_CONTAINER)
					}
				} else {
					fmt.Println(HELP_CONTAINER)
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
