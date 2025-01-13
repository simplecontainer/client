package objects

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/network"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"strconv"
)

const HelpLogs string = "Eg: smr debug {identifier}"

func Logs() contracts.Command {
	return command.Command{
		Name: "logs",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) > 3 {
					resp, err := network.Raw(mgr.Context.Client, fmt.Sprintf("%s/api/v1/logs/%s/%s/%s", mgr.Context.ApiURL, os.Args[3], os.Args[4], strconv.FormatBool(viper.GetBool("f"))), http.MethodGet, nil)

					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					var bytes int
					buff := make([]byte, 512)

					for {
						bytes, err = resp.Body.Read(buff)

						if bytes == 0 || err == io.EOF {
							err = resp.Body.Close()

							if err != nil {
								fmt.Println(err)
								os.Exit(1)
							}

							fmt.Println("server closed connection")
							break
						}

						fmt.Print(string(buff[:bytes]))
					}

					fmt.Print(string(buff[:bytes]))
				} else {
					fmt.Println(HelpLogs)
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
