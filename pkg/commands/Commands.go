package commands

import (
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

var Commands []Command

func PreloadCommands() {
	Context()
	Apply()
	Delete()
	Ps()
	Restore()

	SecretCommand()
	ContainersCommand()
	GitopsCommand()
	ConfigurationCommand()
	ResourceCommand()
	CertKeyCommand()
	HttpAuthCommand()
	LogsCommand()
}

func Run(mgr *manager.Manager) {
	for _, comm := range Commands {
		for k, arg := range os.Args {
			if comm.name == arg && k == 1 {
				if comm.condition(mgr) {
					for _, fn := range comm.depends_on {
						fn(mgr, os.Args)
					}

					for _, fn := range comm.functions {
						fn(mgr, os.Args)
					}
				}
			}
		}
	}
}
