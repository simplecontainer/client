package commands

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/containers"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

const HELP_CONTAINERS string = "Eg: smr configuration [describe, delete, edit, get, list, view]"

func ContainersCommand() {
	Commands = append(Commands, Command{
		name: "containers",
		condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		functions: []func(*manager.Manager, []string){
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
							fmt.Println(HELP_CONTAINERS)
						}
						break
					case "view":
						if len(os.Args) > 4 {
							container.View(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CONTAINERS)
						}
						break
					case "edit":
						if len(os.Args) > 4 {
							container.Edit(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CONTAINERS)
						}
						break
					case "delete":
						if len(os.Args) > 4 {
							container.Delete(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CONTAINERS)
						}
						break
					default:
						fmt.Println(HELP_CONTAINERS)
					}
				} else {
					fmt.Println(HELP_CONTAINERS)
				}
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
