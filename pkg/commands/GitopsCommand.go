package commands

import (
	"fmt"
	"github.com/qdnqn/smr-client/pkg/commands/gitops"
	"github.com/qdnqn/smr-client/pkg/manager"
	"os"
)

const HELP_GITOPS string = "Eg: smr gitops [describe, list, sync]"

func GitopsCommand() {
	Commands = append(Commands, Command{
		name: "gitops",
		condition: func(*manager.Manager) bool {
			return true
		},
		functions: []func(*manager.Manager, []string){
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
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
