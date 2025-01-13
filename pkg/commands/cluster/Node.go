package cluster

import (
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/cluster/node"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/definitions"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/configuration"
	"github.com/spf13/viper"
	"os"
)

func Node() contracts.Command {
	return command.Command{
		Name: "node",
		Condition: func(mgr *manager.Manager) bool {
			return true
		},
		Functions: []func(*manager.Manager, []string){
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
					case "restore":
						node.Restore(mgr)
						break
					}
				case "run":
					node.Run(config, definitions.AgentDefinition(), viper.GetString("platform"))
					break
				case "stop":
					node.Stop(config, definitions.AgentDefinition(), viper.GetString("platform"))
					break
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
