package startup

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/configuration"
	"github.com/simplecontainer/smr/pkg/static"
	"net"
	"os"
)

func GetEnvironmentInfo() *configuration.Environment {
	HOMEDIR, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}

	sudoUser := os.Getenv("SUDO_USER")

	if HOMEDIR == "/root" && sudoUser != "" {
		HOMEDIR = fmt.Sprintf("/home/%s", sudoUser)
	}

	return &configuration.Environment{
		HOMEDIR:    HOMEDIR,
		ROOTDIR:    fmt.Sprintf("%s/%s/%s", HOMEDIR, static.ROOTDIR, static.ROOTSMR),
		LOGDIR:     fmt.Sprintf("%s/%s/%s/logs", HOMEDIR, static.ROOTDIR, static.ROOTSMR),
		FLANNELDIR: fmt.Sprintf("%s/%s/%s/flannel", HOMEDIR, static.ROOTDIR, static.ROOTSMR),
		CLIENTIP:   GetOutboundIP().String(),
	}
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
