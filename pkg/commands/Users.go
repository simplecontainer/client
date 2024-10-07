package commands

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/users"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

func Users() {
	Commands = append(Commands, Command{
		name: "users",
		condition: func(*manager.Manager) bool {
			return true
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) > 2 {
					switch os.Args[2] {
					case "create":
						if len(os.Args) > 4 {
							users.Create(mgr.Context, os.Args[3], os.Args[4], os.Args[5])
						} else {
							fmt.Println("Try this: smr users create bob example.com 8.8.8.8")
						}
						break
					default:
						fmt.Println("Available commands are: create")
					}
				}
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
