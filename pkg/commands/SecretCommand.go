package commands

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/secret"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

const HELP_SECRET string = "Eg: smr secret [describe, list, get, edit]"

func SecretCommand() {
	Commands = append(Commands, Command{
		name: "secret",
		condition: func(*manager.Manager) bool {
			return true
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) > 2 {
					switch os.Args[2] {
					case "list":
						secret.List(mgr.Context)
						break
					case "get":
						if len(os.Args) > 3 {
							secret.Get(mgr.Context, os.Args[3])
						} else {
							fmt.Println(HELP_SECRET)
						}
						break
					case "create":
						if len(os.Args) > 4 {
							secret.Create(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_SECRET)
						}
						break
					default:
						fmt.Println(HELP_SECRET)
					}
				} else {
					fmt.Println(HELP_SECRET)
				}
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
