package commands

import (
	"fmt"
	"github.com/qdnqn/smr-client/pkg/commands/gitops"
	"github.com/qdnqn/smr-client/pkg/manager"
	"os"
)

func Gitops() {
	Commands = append(Commands, Command{
		name: "gitops",
		condition: func(*manager.Manager) bool {
			return true
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				switch os.Args[2] {
				case "describe":
					gitops.Describe(mgr.Context)
					break
				case "list":
					gitops.List(mgr.Context)
					break
				case "sync":
					if len(os.Args) > 4 {
						gitops.Sync(mgr.Context, os.Args[3], os.Args[4])
					}
				default:
					fmt.Println("Available commands are: list, sync")
				}
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
