package commands

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/simplecontainer/client/pkg/commands/restore"
	"github.com/simplecontainer/client/pkg/logger"
	"github.com/simplecontainer/client/pkg/manager"
)

func Restore() {
	Commands = append(Commands, Command{
		name: "restore",
		condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest(mgr.Context)
		},
		functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				prompt := promptui.Select{
					Label: "Are you sure you want to restore state from the key-value store?",
					Items: []string{"no", "yes"},
				}

				if !mgr.Configuration.Flags.Y {
					_, result, err := prompt.Run()

					if err != nil {
						logger.Log.Fatal("failed to select from list of contexts")
					}

					if result == "yes" {
						restore.Restore(mgr.Context)
					} else {
						fmt.Println("restore canceled")
					}
				} else {
					restore.Restore(mgr.Context)
				}
			},
		},
		depends_on: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	})
}
