package control

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/control/gitops"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

const HELP_GITOPS string = "Eg: smr gitops [describe, delete, edit, get, list, sync]"

func Gitops() contracts.Command {
	return command.Command{
		Name: "gitops",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) > 2 {
					switch os.Args[2] {
					case "describe":
						gitops.Describe(mgr.Context)
						break
					case "list":
						gitops.List(mgr.Context)
						break
					case "get":
						if len(os.Args) > 4 {
							gitops.Get(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_GITOPS)
						}
						break
					case "edit":
						if len(os.Args) > 4 {
							gitops.Edit(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_GITOPS)
						}
						break
					case "delete":
						if len(os.Args) > 4 {
							gitops.Delete(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_GITOPS)
						}
						break
					case "refresh":
						if len(os.Args) > 4 {
							gitops.Refresh(mgr.Context, os.Args[3], os.Args[4])
						}
					case "sync":
						if len(os.Args) > 4 {
							gitops.Sync(mgr.Context, os.Args[3], os.Args[4])
						}
					default:
						fmt.Println(HELP_GITOPS)
					}
				} else {
					fmt.Println(HELP_GITOPS)
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
