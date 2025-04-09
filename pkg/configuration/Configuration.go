package configuration

import (
	"fmt"
	"github.com/simplecontainer/smr/pkg/static"
	"os"
)

func NewConfig() *Configuration {
	return &Configuration{
		Environment: GetEnvironmentInfo(),
	}
}

func GetEnvironmentInfo() *Environment {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}

	sudoUser := os.Getenv("SUDO_USER")

	if home == "/root" && sudoUser != "" {
		home = fmt.Sprintf("/home/%s", sudoUser)
	}

	return &Environment{
		Home:            home,
		ClientDirectory: fmt.Sprintf("%s/%s", home, static.ROOTSMR),
		LogsDirectory:   fmt.Sprintf("%s/%s/logs", home, static.ROOTSMR),
	}
}
