package control

import (
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/control/control"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/formaters"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/static"
	"os"
)

func List() contracts.Command {
	return command.Command{
		Name: "list",
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

				var objects []json.RawMessage

				switch format.GetKind() {
				case static.KIND_GITOPS:
					objects, err = control.ListKind(mgr.Context, format.GetPrefix(), format.GetVersion(), static.CATEGORY_STATE, format.GetKind())
					formaters.Gitops(objects)
					break
				case static.KIND_CONTAINER:
					objects, err = control.ListKind(mgr.Context, format.GetPrefix(), format.GetVersion(), static.CATEGORY_STATE, format.GetKind())
					formaters.Container(objects)
					break
				default:
					objects, err = control.ListKind(mgr.Context, format.GetPrefix(), format.GetVersion(), static.CATEGORY_KIND, format.GetKind())
					formaters.Default(objects)
					break
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
