package commands

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/remove"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/definitions"
	"os"
)

func Delete() {
	Commands = append(Commands, Command{
		name: "delete",
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
				remove.Remove(mgr.Context, definition)
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
