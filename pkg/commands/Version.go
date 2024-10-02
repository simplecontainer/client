package commands

import (
	"github.com/simplecontainer/client/pkg/commands/version"
	"github.com/simplecontainer/client/pkg/manager"
)

func Version() {
	Commands = append(Commands, Command{
		name: "version",
		condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest(mgr.Context)
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				version.Version(mgr.VersionClient, mgr.Context)
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
