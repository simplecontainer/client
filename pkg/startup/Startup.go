package startup

import (
	"flag"
	"github.com/simplecontainer/client/pkg/configuration"
	"github.com/simplecontainer/smr/pkg/static"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func Load(configObj *configuration.Configuration, projectDir string) {
	configObj.Environment = GetEnvironmentInfo()
	ReadFlags(configObj)
}

func ReadFlags(configObj *configuration.Configuration) {
	/* Operation mode */
	flag.String("context", "", "Context name")
	flag.Bool("y", false, "Say yes to everything")

	HOMEDIR, _ := os.UserHomeDir()

	flag.Bool("wait", false, "Wait for container to exit")

	flag.String("domains", "localhost", "Comma separated list of the domains")
	flag.String("ips", "127.0.0.1", "Comma separated list of the IPs")
	flag.String("homedir", HOMEDIR, "Host homedir")

	flag.String("image", "simplecontainermanager/smr", "The smr image repo")
	flag.String("tag", "latest", "The smr image tag")
	flag.String("entrypoint", "/opt/smr/smr", "Entrypoint for the smr")
	flag.String("args", "create smr --agent smr-agent", "args")
	flag.String("hostport", "1443", "Expose smr on hostport")
	flag.String("overlayport", "", "Expose overlay on port")

	flag.String("platform", static.PLATFORM_DOCKER, "Container engine name. Supported: [docker]")
	flag.String("agent", "smr-agent", "Name of the smr agent container")

	flag.String("node", "", "Node ID in the cluster - must be unique")
	flag.String("url", "", "Reachable Node https://URL:PORT combination")
	flag.String("cluster", "", "Reachable list of peers; comma separated - eg. https://URL:PORT,...")
	flag.String("etcdport", "2379", "Etcd client port listen and advertise")
	flag.Bool("join", false, "Join the cluster")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	configObj.Flags.Context = viper.GetString("context")
	configObj.Flags.Y = viper.GetBool("y")
}
