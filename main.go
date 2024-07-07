package main

import (
	"github.com/simplecontainer/client/pkg/bootstrap"
	"github.com/simplecontainer/client/pkg/commands"
	_ "github.com/simplecontainer/client/pkg/commands"
	"github.com/simplecontainer/client/pkg/configuration"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/startup"
	"github.com/simplecontainer/smr/pkg/logger"
)

func main() {
	logger.Log = logger.NewLogger()

	config := configuration.NewConfig()
	startup.Load(config, config.Root)

	manager := &manager.Manager{}
	manager.Configuration = config

	bootstrap.CreateDirectoryTree(manager.Configuration.Environment.PROJECTDIR)

	manager.Context = context.LoadContext(manager.Configuration.Environment.PROJECTDIR)

	if manager.Context == nil {
		logger.Log.Fatal("please first connect to one smr-agent instance")
	}

	commands.PreloadCommands()
	commands.Run(manager)
}
