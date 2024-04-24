package commands

import (
	"fmt"
	"github.com/spf13/viper"
	"smr/pkg/manager"
	"smr/pkg/utils"
)

func Delete() {
	Commands = append(Commands, Command{
		name: "delete",
		condition: func(*manager.Manager) bool {
			return viper.GetString("project") != ""
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if utils.Confirm(fmt.Sprintf("Are you sure? Delete project %s is irreversible?", mgr.Runtime.PROJECT)) {
					mgr.Load()
					mgr.DeleteProject(viper.GetString("project"))
				}
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {

			},
		},
	})
}
