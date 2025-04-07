package cluster

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/cluster/nodes"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/node"
	"os"
)

func Node() contracts.Command {
	return command.Command{
		Name: "node",
		Condition: func(mgr *manager.Manager) bool {
			return true
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if helpers.GrabArg(2) == "cluster" {
					switch os.Args[3] {
					case "join":
						nodes.Join(mgr)
						break
					case "leave":
						nodes.Leave(mgr)
						break
					}
				} else {
					n, err := node.New(mgr.Configuration.Setup.Node, mgr.Configuration)

					if err != nil {
						panic(err)
					}

					err = node.Directory(n.Name, mgr.Configuration.Environment.Home)

					if err != nil {
						panic(err)
					}

					switch helpers.GrabArg(2) {
					case "create":
						err = nodes.Create(n, mgr.Configuration)

						if err != nil {
							helpers.ExitWithErr(err)
						}

						if mgr.Configuration.Flags.W != "" {
							err = n.Wait(mgr.Configuration.Flags.W)

							if err != nil {
								helpers.ExitWithErr(err)
							}
						}

						fmt.Println(n.Container.GetId())
						break
					case "run":
						err = nodes.Run(n)

						if err != nil {
							helpers.ExitWithErr(err)
						}

						if mgr.Configuration.Flags.W != "" {
							err = n.Wait(mgr.Configuration.Flags.W)

							if err != nil {
								helpers.ExitWithErr(err)
							}
						}

						fmt.Println(n.Container.GetId())
						break
					case "rename":
						err = nodes.Rename(n, helpers.GrabArg(3))

						if err != nil {
							helpers.ExitWithErr(err)
						}

						if mgr.Configuration.Flags.W != "" {
							err = n.Wait(mgr.Configuration.Flags.W)

							if err != nil {
								helpers.ExitWithErr(err)
							}
						}

						fmt.Println("node renamed")
						break
					case "restart":
						err = nodes.Restart(n)

						if err != nil {
							helpers.ExitWithErr(err)
						}

						if mgr.Configuration.Flags.W != "" {
							err = n.Wait(mgr.Configuration.Flags.W)

							if err != nil {
								helpers.ExitWithErr(err)
							}
						}

						fmt.Println("node restarted")
						break
					case "stop":
						err = nodes.Stop(n)

						if err != nil {
							helpers.ExitWithErr(err)
						}

						if mgr.Configuration.Flags.W != "" {
							err = n.Wait(mgr.Configuration.Flags.W)

							if err != nil {
								helpers.ExitWithErr(err)
							}
						}

						fmt.Println("node stopped")
						break
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
