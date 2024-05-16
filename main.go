package main

import (
	"github.com/qdnqn/smr-client/pkg/bootstrap"
	"github.com/qdnqn/smr-client/pkg/commands"
	_ "github.com/qdnqn/smr-client/pkg/commands"
	"github.com/qdnqn/smr-client/pkg/context"
	"github.com/qdnqn/smr-client/pkg/manager"
	"github.com/qdnqn/smr/pkg/config"
	"github.com/qdnqn/smr/pkg/logger"
	"github.com/qdnqn/smr/pkg/runtime"
	"github.com/spf13/viper"
)

func main() {
	logger.Log = logger.NewLogger()

	conf := config.NewConfig()
	conf.ReadFlags()

	manager := &manager.Manager{}
	manager.Config = conf

	viper.Set("project", "smr")
	manager.Runtime = runtime.GetRuntimeInfo()

	bootstrap.CreateDirectoryTree(manager.Runtime.PROJECTDIR)

	manager.Context = context.LoadContext(manager.Runtime.PROJECTDIR)

	if manager.Context == nil {
		logger.Log.Fatal("please first connect to one smr-agent instance")
	}

	commands.PreloadCommands()
	commands.Run(manager)
}
