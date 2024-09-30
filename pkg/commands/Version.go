package commands

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/version"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

func Version() {
	Commands = append(Commands, Command{
		name: "version",
		condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				version.Version(mgr.VersionClient, mgr.Context)
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if mgr.Context == nil {
					fmt.Println("no active context found - please add least one context")
					os.Exit(1)
				}
			},
		},
	})
}
