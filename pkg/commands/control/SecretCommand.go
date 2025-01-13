package control

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/control/secret"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

const HELP_SECRET string = "Eg: smr secret [set, get, list]"

func Secret() contracts.Command {
	return command.Command{
		Name: "secret",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
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
					case "set":
						if len(os.Args) > 4 {
							secret.Set(mgr.Context, os.Args[3], []byte(os.Args[4]))
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
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
