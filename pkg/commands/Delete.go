package commands

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/remove"
	"github.com/simplecontainer/client/pkg/definitions"
	"github.com/simplecontainer/client/pkg/manager"
	"net/url"
	"os"
)

func Delete() {
	Commands = append(Commands, Command{
		name: "delete",
		condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) < 2 {
					fmt.Println("try to specify a file")
					return
				}

				u, err := url.ParseRequestURI(args[2])

				definition := ""

				if err != nil || !u.IsAbs() {
					definition, err = definitions.ReadFile(args[2])
				} else {
					definition, err = definitions.DownloadFile(u)
				}

				if err != nil {
					fmt.Println(err)
				} else {
					if definition != "" {
						remove.Remove(mgr.Context, definition)
					} else {
						fmt.Println("specified file/url is not valid definition")
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
