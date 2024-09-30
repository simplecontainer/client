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

	managerObj := &manager.Manager{}
	managerObj.VersionClient = SMR_VERSION
	managerObj.Configuration = config

	bootstrap.CreateDirectoryTree(managerObj.Configuration.Environment.PROJECTDIR)

	managerObj.Context = context.LoadContext(managerObj.Configuration.Environment.PROJECTDIR)

	commands.PreloadCommands()
	commands.Run(managerObj)
}
