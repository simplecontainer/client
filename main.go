package main

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/bootstrap"
	"github.com/simplecontainer/client/pkg/commands"
	_ "github.com/simplecontainer/client/pkg/commands"
	"github.com/simplecontainer/client/pkg/configuration"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/logger"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/startup"
	"github.com/simplecontainer/smr/pkg/static"
	"log"
	"os"
)

func main() {
	config := configuration.NewConfig()
	startup.Load(config, config.Root)

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = static.DEFAULT_LOG_LEVEL
	}

	managerObj := &manager.Manager{}
	managerObj.VersionClient = SMR_VERSION
	managerObj.Configuration = config

	bootstrap.CreateDirectoryTree(managerObj.Configuration.Environment.ROOTDIR)

	logger.Log = logger.NewLogger(config.Environment.LOGDIR, logLevel)
	logger.LogFlannel = logger.NewLoggerFlannel(config.Environment.LOGDIR, logLevel)

	f, err := os.OpenFile(fmt.Sprintf("%s/flannel.log", config.Environment.LOGDIR), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	managerObj.Context = context.NewContext(managerObj.Configuration.Environment.ROOTDIR)
	managerObj.Context.LoadContext()

	commands.PreloadCommands()
	commands.Run(managerObj)
}
