package commands

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/node"
	"github.com/simplecontainer/client/pkg/definitions"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/configuration"
	"github.com/simplecontainer/smr/pkg/static"
	"github.com/spf13/viper"
	"os"
)

func Node() {
	Commands = append(Commands, Command{
		name: "node",
		condition: func(mgr *manager.Manager) bool {
			return true
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				config := &configuration.Configuration{
					HostHome: viper.GetString("homedir"),
					Environment: &configuration.Environment{
						HOMEDIR: viper.GetString("homedir"),
						AGENTIP: "",
					},
				}

				switch os.Args[2] {
				case "cluster":
					switch os.Args[3] {
					case "start":
						node.Start(mgr)
						break
					case "add":
						node.Add(mgr)
						break
					case "delete":
						node.Delete(mgr)
						break
					case "get":
						node.Get(mgr)
						break
					}
				case "run":
					switch viper.GetString("platform") {
					case static.PLATFORM_DOCKER:
						node.RunDocker(config, definitions.AgentDefinition())
						break
					default:
						fmt.Println("platform is not supported")
						return
					}
					break
				}
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
