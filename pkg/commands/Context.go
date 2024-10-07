package commands

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/context"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

func Context() {
	Commands = append(Commands, Command{
		name: "context",
		condition: func(*manager.Manager) bool {
			return true
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) < 3 {
					if mgr.Context != nil {
						fmt.Println(fmt.Sprintf("active context is %s", mgr.Context.Name))
					} else {
						fmt.Println("no active context found - please add least one context")
					}
				} else {
					switch os.Args[2] {
					case "connect":
						if len(os.Args) > 4 {
							context.Connect(os.Args[3], os.Args[4], mgr.Configuration.Environment.PROJECTDIR)
						} else {
							fmt.Println("Try this: smr context connect https://API_URL:1443 PATH_TO_CERT.PEM --context NAME_YOU_WANT")
						}
						break
					case "switch":
						contextName := ""
						if len(os.Args) > 3 {
							contextName = os.Args[3]
						}

						context.Switch(contextName, mgr.Context)
						break
					default:
						fmt.Println("Available commands are: connect, switch")
					}
				}
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
