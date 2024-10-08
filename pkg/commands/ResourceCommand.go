package commands

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/resource"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

const HELP_RESOURCE string = "Eg: smr resource [describe, delete, edit, get, list]"

func ResourceCommand() {
	Commands = append(Commands, Command{
		name: "resource",
		condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest(mgr.Context)
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) > 2 {
					switch os.Args[2] {
					case "describe":
						resource.Describe(mgr.Context)
						break
					case "list":
						resource.List(mgr.Context)
						break
					case "get":
						if len(os.Args) > 4 {
							resource.Get(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_RESOURCE)
						}
						break
					case "edit":
						if len(os.Args) > 4 {
							resource.Edit(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_RESOURCE)
						}
						break
					case "delete":
						if len(os.Args) > 4 {
							resource.Delete(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_RESOURCE)
						}
						break
					default:
						fmt.Println(HELP_RESOURCE)
					}
				} else {
					fmt.Println(HELP_RESOURCE)
				}
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
