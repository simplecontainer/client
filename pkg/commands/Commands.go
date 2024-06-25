package commands

import (
	"github.com/qdnqn/smr-client/pkg/manager"
	"os"
)

var Commands []Command

func PreloadCommands() {
	Context()
	Apply()
	Delete()
	Ps()

	ContainersCommand()
	GitopsCommand()
	ConfigurationCommand()
	ResourceCommand()
	CertKeyCommand()
	HttpAuthCommand()
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
