package objects

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/objects/apply"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/definitions"
	"github.com/simplecontainer/client/pkg/manager"
	"net/url"
	"os"
)

func Apply() contracts.Command {
	return command.Command{
		Name: "apply",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) < 2 {
					fmt.Println("try to specify a file")
					return
				}

				u, err := url.ParseRequestURI(args[2])

				var definition []byte

				if err != nil || !u.IsAbs() {
					definition, err = definitions.ReadFile(args[2])
				} else {
					definition, err = definitions.DownloadFile(u)
				}

				if err != nil {
					fmt.Println(err)
				} else {
					if definition != nil {
						response := apply.Apply(mgr.Context, definition)

						if response.Success {
							fmt.Println(response.Explanation)
						} else {
							fmt.Println(response.ErrorExplanation)
						}
					} else {
						fmt.Println("specified file/url is not valid definition")
					}
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
