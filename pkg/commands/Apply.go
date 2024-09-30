package commands

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/apply"
	"github.com/simplecontainer/client/pkg/definitions"
	"github.com/simplecontainer/client/pkg/manager"
	"net/url"
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
				u, err := url.ParseRequestURI(args[2])

				definition := ""

				if err != nil {
					definition, err = definitions.ReadFile(args[2])
				} else {
					definition, err = definitions.DownloadFile(u)
				}

				if err != nil {
					fmt.Println(err)
				} else {
					if definition != "" {
						apply.Apply(mgr.Context, definition)
					} else {
						fmt.Println("specified file/url is not valid definition")
					}
				}
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if mgr.Context != nil {

				} else {
					fmt.Println("no active context found - please add least one context")
					os.Exit(1)
				}
			},
		},
	})
}
