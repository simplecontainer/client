package commands

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/ps"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

const HELP_PS string = "Eg: smr ps ['', watch]"

func Ps() {
	Commands = append(Commands, Command{
		name: "ps",
		condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest(mgr.Context)
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) > 2 {
					switch os.Args[2] {
					case "watch":
						ps.Ps(mgr.Context, true)
						break
					default:
						fmt.Println(HELP_PS)
					}
				} else {
					ps.Ps(mgr.Context, false)
				}
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
