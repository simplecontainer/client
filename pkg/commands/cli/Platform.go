package cli

import (
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/cli/platform"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/definitions"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/configuration"
	"github.com/spf13/viper"
	"os"
)

func Platform() contracts.Command {
	return command.Command{
		Name: "cli",
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
				case "inspect":
					platform.Inspect(config, definitions.NodeDefinition(), viper.GetString("platform"))
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
