package streams

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/network"
	"github.com/spf13/viper"
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
				format, err := helpers.BuildFormat(helpers.GrabArg(2), mgr.Configuration.G)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				resp, err := network.Raw(mgr.Context.Client, fmt.Sprintf("%s/api/v1/logs/%s/%s/%s", mgr.Context.ApiURL, format.ToString(), viper.GetString("container"), strconv.FormatBool(viper.GetBool("f"))), http.MethodGet, nil)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				err = helpers.PrintBytes(resp.Body)

				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
