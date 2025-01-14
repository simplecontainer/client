package control

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/control/certkey"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

const HELP_CERTKEY string = "Eg: smr configuration [describe, delete, edit, get, list]"

func CertKey() contracts.Command {
	return command.Command{
		Name: "certkey",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) > 2 {
					switch os.Args[2] {
					case "describe":
						certkey.Describe(mgr.Context)
						break
					case "list":
						certkey.List(mgr.Context)
						break
					case "get":
						if len(os.Args) > 4 {
							certkey.Get(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CERTKEY)
						}
						break
					case "edit":
						if len(os.Args) > 4 {
							certkey.Edit(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CERTKEY)
						}
						break
					case "delete":
						if len(os.Args) > 4 {
							certkey.Delete(mgr.Context, os.Args[3], os.Args[4])
						} else {
							fmt.Println(HELP_CERTKEY)
						}
						break
					default:
						fmt.Println(HELP_CERTKEY)
					}
				} else {
					fmt.Println(HELP_CERTKEY)
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
