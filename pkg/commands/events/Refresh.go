package events

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/control/control"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/events/events"
	"github.com/simplecontainer/smr/pkg/static"
	"os"
)

func Refresh() contracts.Command {
	return command.Command{
		Name: "refresh",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				format, err := helpers.BuildFormat(helpers.GrabArg(2), mgr.Configuration.G)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				event := events.New(events.EVENT_REFRESH, static.KIND_GITOPS, static.SMR_PREFIX, static.KIND_GITOPS, format.GetGroup(), format.GetName(), nil)

				var bytes []byte
				bytes, err = event.ToJSON()

				control.Event(mgr.Context, format.GetPrefix(), format.GetVersion(), static.CATEGORY_EVENT, format.GetKind(), format.GetGroup(), format.GetName(), bytes)
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
