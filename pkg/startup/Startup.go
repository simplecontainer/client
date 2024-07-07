package startup

import (
	"flag"
	"github.com/simplecontainer/client/pkg/configuration"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Load(configObj *configuration.Configuration, projectDir string) {
	configObj.Environment = GetEnvironmentInfo()
	ReadFlags(configObj)
}

func ReadFlags(configObj *configuration.Configuration) {
	/* Operation mode */
	flag.String("context", "", "Context name")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	configObj.Flags.Context = viper.GetString("context")
}
