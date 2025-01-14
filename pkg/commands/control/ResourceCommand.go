package control

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/control/resource"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

const HELP_RESOURCE string = "Eg: smr resource [describe, delete, edit, get, list]"

func Resource() contracts.Command {
	return command.Command{
		Name: "resource",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
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
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
