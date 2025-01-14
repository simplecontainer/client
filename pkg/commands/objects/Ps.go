package objects

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/objects/ps"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

const HELP_PS string = "Eg: smr ps ['', watch]"

func Ps() contracts.Command {
	return command.Command{
		Name: "ps",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
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
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
