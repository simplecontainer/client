package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/simplecontainer/client/pkg/bootstrap"
	"github.com/simplecontainer/client/pkg/commands"
	_ "github.com/simplecontainer/client/pkg/commands"
	"github.com/simplecontainer/client/pkg/configuration"
	"github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/logger"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/startup"
	"github.com/simplecontainer/smr/pkg/static"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = static.DEFAULT_LOG_LEVEL
	}

	startup.ReadFlags()
	config, err := startup.Load(viper.GetString("node"), configuration.GetEnvironmentInfo())

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			startup.LoadFromFlags(config)
		} else {
			log.Fatal(err)
		}
	} else {
		startup.LoadFromFlagsDynamic(config)
	}

	bootstrap.CreateDirectoryTree(config.Environment.ClientDirectory)

	managerObj := &manager.Manager{}
	managerObj.VersionClient = SMR_VERSION
	managerObj.Configuration = config

	logger.Log = logger.NewLogger(config.Environment.LogsDirectory, logLevel)
	logger.LogFlannel = logger.NewLoggerFlannel(config.Environment.LogsDirectory, logLevel)

	f, err := os.OpenFile(fmt.Sprintf("%s/flannel.log", config.Environment.LogsDirectory), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	managerObj.Context = context.NewContext(managerObj.Configuration.Environment.ClientDirectory)
	managerObj.Context.LoadContext()

	commands.PreloadCommands()
	commands.Run(managerObj)
}
