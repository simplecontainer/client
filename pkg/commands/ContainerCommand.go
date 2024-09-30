package commands

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/containers"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

const HELP_CONTAINERS string = "Eg: smr configuration [describe, edit, get, list, view]"

func ContainersCommand() {
	Commands = append(Commands, Command{
		name: "container",
		condition: func(mgr *manager.Manager) bool {

			return mgr.Context.ConnectionTest()
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) > 2 {
					switch os.Args[2] {
					case "describe":
						containers.Describe(mgr.Context)
						break
					case "list":
						containers.List(mgr.Context)
						break
					case "get":
						if len(os.Args) > 4 {
							containers.Get(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CONTAINERS)
						}
						break
					case "view":
						if len(os.Args) > 4 {
							containers.View(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CONTAINERS)
						}
						break
					case "edit":
						if len(os.Args) > 4 {
							containers.Edit(mgr.Context, os.Args[3], os.Args[4])
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
				if mgr.Context == nil {
					fmt.Println("no active context found - please add least one context")
					os.Exit(1)
				}
			},
		},
	})
}
