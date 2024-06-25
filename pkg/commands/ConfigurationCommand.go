package commands

import (
	"fmt"
	"github.com/qdnqn/smr-client/pkg/commands/configuration"
	"github.com/qdnqn/smr-client/pkg/manager"
	"os"
)

const HELP_CONFIGURATION string = "Eg: smr configuration [describe, list, get]"

func ConfigurationCommand() {
	Commands = append(Commands, Command{
		name: "configuration",
		condition: func(*manager.Manager) bool {
			return true
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) > 2 {
					switch os.Args[2] {
					case "describe":
						configuration.Describe(mgr.Context)
						break
					case "list":
						configuration.List(mgr.Context)
						break
					case "get":
						if len(os.Args) > 4 {
							configuration.Get(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CONFIGURATION)
						}
						break
					case "edit":
						if len(os.Args) > 4 {
							configuration.Edit(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CONFIGURATION)
						}
						break
					default:
						fmt.Println(HELP_CONFIGURATION)
					}
				} else {
					fmt.Println(HELP_CONFIGURATION)
				}
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
