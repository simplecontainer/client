package main

import (
	"smr/pkg/commands"
	_ "smr/pkg/commands"
	"smr/pkg/config"
	"smr/pkg/logger"
	"smr/pkg/manager"
	"smr/pkg/runtime"
)

func main() {
	logger.Log = logger.NewLogger()

	conf := config.NewConfig()
	conf.ReadFlags()

	manager := &manager.Manager{}
	manager.Config = conf
	manager.Runtime = runtime.GetRuntimeInfo()

	commands.PreloadCommands()
	commands.Run(manager)
}
