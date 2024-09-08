package commands

import (
	"github.com/simplecontainer/client/pkg/commands/restore"
	"github.com/simplecontainer/client/pkg/manager"
)

func Restore() {
	Commands = append(Commands, Command{
		name: "restore",
		condition: func(*manager.Manager) bool {
			return true
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				restore.Restore(mgr.Context)
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
