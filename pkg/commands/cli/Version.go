package cli

import (
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/cli/version"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/manager"
)

func Version() contracts.Command {
	return command.Command{
		Name: "version",
		Condition: func(mgr *manager.Manager) bool {
			return true
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				version.Version(mgr.VersionClient, mgr.Context)
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
