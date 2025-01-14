package control

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/control/httpauth"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

const HELP_HTTPAUTH string = "Eg: smr httpauth [describe, delete, edit, get, list]"

func HttpAuth() contracts.Command {
	return command.Command{
		Name: "httpauth",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) > 2 {
					switch os.Args[2] {
					case "describe":
						httpauth.Describe(mgr.Context)
						break
					case "list":
						httpauth.List(mgr.Context)
						break
					case "get":
						if len(os.Args) > 4 {
							httpauth.Get(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_HTTPAUTH)
						}
						break
					case "edit":
						if len(os.Args) > 4 {
							httpauth.Edit(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_HTTPAUTH)
						}
						break
					case "delete":
						if len(os.Args) > 4 {
							httpauth.Delete(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_HTTPAUTH)
						}
						break
					default:
						fmt.Println(HELP_HTTPAUTH)
					}
				} else {
					fmt.Println(HELP_HTTPAUTH)
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
