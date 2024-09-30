package commands

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/definitions"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

const HELP_DEFINITIONS string = `
Definitions command used to see loaded definitions on the agent.
Eg: smr definitions ['',... any output from the smr definitions]"
`

func Definitions() {
	Commands = append(Commands, Command{
		name:      "definitions",
		condition: func(mgr *manager.Manager) bool { return mgr.Context.ConnectionTest() },
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				switch len(os.Args) {
				case 2:
					definitions.Definitions(mgr.Context, "")
					break
				case 3:
					definitions.Definitions(mgr.Context, os.Args[2])
					break
				default:
					fmt.Println(HELP_DEFINITIONS)
				}
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
