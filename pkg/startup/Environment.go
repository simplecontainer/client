package startup

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/configuration"
	"github.com/simplecontainer/smr/pkg/static"
	"os"
)

func GetEnvironmentInfo() *configuration.Environment {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}

	sudoUser := os.Getenv("SUDO_USER")

	if home == "/root" && sudoUser != "" {
		home = fmt.Sprintf("/home/%s", sudoUser)
	}

	return &configuration.Environment{
		Home:            home,
		ClientDirectory: fmt.Sprintf("%s/%s/%s", home, static.ROOTDIR, static.ROOTSMR),
		LogsDirectory:   fmt.Sprintf("%s/%s/%s/logs", home, static.ROOTDIR, static.ROOTSMR),
	}
}
