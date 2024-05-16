package commands

import (
	"fmt"
	"github.com/qdnqn/smr-client/pkg/commands/apply"
	"github.com/qdnqn/smr-client/pkg/manager"
	"github.com/qdnqn/smr/pkg/definitions"
	"os"
)

func Apply() {
	Commands = append(Commands, Command{
		name: "apply",
		condition: func(*manager.Manager) bool {
			if len(os.Args) > 2 {
				return true
			} else {
				fmt.Println("try to specify a file")
				return false
			}
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				definition := definitions.ReadFile(args[2])
				apply.Apply(mgr.Context, definition)
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
