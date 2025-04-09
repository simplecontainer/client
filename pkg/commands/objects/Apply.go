package objects

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/packer"
	"github.com/simplecontainer/smr/pkg/relations"
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
					fmt.Println("try to specify a path/url")
					return
				}

				u, err := url.ParseRequestURI(args[2])

				var pack = packer.New()

				if err != nil || !u.IsAbs() {
					var stat os.FileInfo
					stat, err = os.Stat(args[2])

					if os.IsNotExist(err) {
						fmt.Println("path does not exist")
						return
					}

					if stat.IsDir() {
						kinds := relations.NewDefinitionRelationRegistry()
						kinds.InTree()

						pack, err = packer.Read(args[2], kinds)

						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					} else {
						var definitions []byte
						definitions, err = packer.ReadYAMLFile(args[2])

						if err != nil {
							fmt.Println(err.Error())
							os.Exit(1)
						}

						pack.Definitions, err = packer.Parse(definitions)

						if err != nil {
							fmt.Println(err.Error())
							os.Exit(1)
						}
					}
				} else {
					var definitions []byte
					definitions, err = packer.Download(u)

					if err != nil {
						fmt.Println(err.Error())
						os.Exit(1)
					}

					pack.Definitions, err = packer.Parse(definitions)

					if err != nil {
						fmt.Println(err.Error())
						os.Exit(1)
					}
				}

				if len(pack.Definitions) != 0 {
					for _, definition := range pack.Definitions {
						err = definition.ProposeApply(mgr.Context.Client, mgr.Context.ApiURL)

						if err != nil {
							fmt.Println(err)
						} else {
							fmt.Println(fmt.Sprintf("object applied: %s", definition.Definition.GetKind()))
						}
					}
				} else {
					fmt.Println("specified file/url is not valid definition")
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
