package objects

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/control/control"
	"github.com/simplecontainer/client/pkg/commands/objects/remove"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/manager"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
	common "github.com/simplecontainer/smr/pkg/kinds/common"
	"os"
)

func Remove() contracts.Command {
	return command.Command{
		Name: "remove",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				format, err := helpers.BuildFormat(helpers.GrabArg(2), mgr.Configuration.Startup.G)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				get, err := control.Get(mgr.Context, format.GetPrefix(), format.GetVersion(), format.GetCategory(), format.GetKind(), format.GetGroup(), format.GetName())

				fmt.Println(string(get))

				c := v1.CommonDefinition{}

				err = json.Unmarshal(get, &c)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				request, err := common.NewRequest(c.GetKind())

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				err = request.Definition.FromJson(get)

				if err != nil {
					fmt.Println(err)
				} else {
					var bytes []byte
					bytes, err = request.Definition.ToJsonForUser()

					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					remove.Remove(mgr.Context, bytes)
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
