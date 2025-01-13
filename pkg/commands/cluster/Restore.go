package cluster

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/cluster/restore"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/logger"
	"github.com/simplecontainer/client/pkg/manager"
)

func Restore() contracts.Command {
	return command.Command{
		Name: "restore",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
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
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
