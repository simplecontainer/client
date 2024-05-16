package commands

import (
	"github.com/qdnqn/smr-client/pkg/commands/ps"
	"github.com/qdnqn/smr-client/pkg/manager"
)

func Ps() {
	Commands = append(Commands, Command{
		name:      "ps",
		condition: func(*manager.Manager) bool { return true },
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				ps.Ps(mgr.Context)
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
