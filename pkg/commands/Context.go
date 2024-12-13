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
						fmt.Println(mgr.Context.Name)
					} else {
						fmt.Println("no active context found - please add least one context")
					}
				} else {
					switch os.Args[2] {
					case "connect":
						if len(os.Args) > 4 {
							context.Connect(os.Args[3], os.Args[4], mgr.Configuration.Environment.ROOTDIR)
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
					case "export":
						contextName := ""
						if len(os.Args) > 3 {
							contextName = os.Args[3]
						}

						API := ""
						_, err := fmt.Scan(&API)

						if err != nil {
							fmt.Println("failed to read API URL, please specify API url")
							os.Exit(1)
						}

						context.Export(contextName, mgr.Context, mgr.Configuration.Environment.ROOTDIR, API)
					case "import":
						encrypted := ""
						if len(os.Args) > 3 {
							encrypted = os.Args[3]
						}

						key := ""
						_, err := fmt.Scan(&key)

						if err != nil {
							fmt.Println("failed to read decryption key, please specify key in stdin")
							os.Exit(1)
						}

						context.Import(encrypted, mgr.Context, mgr.Configuration.Environment.ROOTDIR, key)
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
