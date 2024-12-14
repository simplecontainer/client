package commands

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/network"
	"os"
)

const HELP_LOGS string = "Eg: smr logs {group} {identifier}"

func LogsCommand() {
	Commands = append(Commands, Command{
		name: "logs",
		condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) > 3 {
					network.TailLogs(mgr.Context.Client, fmt.Sprintf("%s/api/v1/logs/%s/%s/%s", mgr.Context.ApiURL, os.Args[2], os.Args[3], os.Args[4]))
				} else {
					fmt.Println(HELP_LOGS)
				}
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
