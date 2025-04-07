package startup

import (
	"flag"
	"fmt"
	"github.com/simplecontainer/client/pkg/configuration"
	"github.com/simplecontainer/smr/pkg/static"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadFromFlagsDynamic(configObj *configuration.Configuration) {
	_ = UnmarshalFields(viper.GetViper(), []string{
		"api", "join", "node", "config",
	}, &configObj.Dynamic)
}

func LoadFromFlags(configObj *configuration.Configuration) {
	_ = UnmarshalFields(viper.GetViper(), []string{
		"y", "f", "o", "w", "g",
	}, &configObj.Flags)

	_ = UnmarshalFields(viper.GetViper(), []string{
		"platform", "node", "context", "log", "domains", "ips", "image", "tag", "entrypoint", "args", "hostport", "overlayport", "etcdport",
	}, &configObj.Setup)

	_ = UnmarshalFields(viper.GetViper(), []string{
		"fbackend", "fcidr", "fiface",
	}, &configObj.Network)

	_ = UnmarshalFields(viper.GetViper(), []string{
		"api", "join", "node",
	}, &configObj.Dynamic)

	_ = UnmarshalFields(viper.GetViper(), []string{
		"fbackend", "fcidr", "finterface", "fenableIPv4", "fenableIPv6", "fmaskIPv6",
	}, &configObj.Flannel)
}

func Load(node string, environment *configuration.Environment) (*configuration.Configuration, error) {
	configObj := configuration.NewConfig()
	path := fmt.Sprintf("%s/%s/%s/%s.yaml", environment.Home, static.ROOTSMR, static.CONFIGDIR, node)

	file, err := os.Open(path)

	defer func() {
		file.Close()
	}()

	if err != nil {
		return configObj, err
	}

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(file)

	if err != nil {
		return configObj, err
	}

	err = viper.Unmarshal(configObj)

	if err != nil {
		return configObj, err
	}

	return configObj, err
}

func Save(configObj *configuration.Configuration) error {
	yamlObj, err := yaml.Marshal(*configObj)

	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s/%s/%s.yaml", configObj.Environment.Home, static.ROOTSMR, static.CONFIGDIR, configObj.Dynamic.Node)

	err = os.WriteFile(path, yamlObj, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadFlags() {
	// Cli flags
	flag.String("w", "", "Wait for container to be in defined state")
	flag.Bool("f", false, "Follow logs")
	flag.String("o", "d", "Output type: d(efault),s(hort)")
	flag.Bool("y", false, "Say yes to everything")
	flag.String("g", "default", "Group")
	flag.Bool("it", false, "Interactive exec")
	flag.String("c", "", "Command for exec")
	flag.String("context", "", "Context")
	flag.String("container", "main", "Which container to stream main or init?")

	// Node setup flags
	flag.String("image", "quay.io/simplecontainer/smr", "The smr image repo")
	flag.String("tag", "latest", "The smr image tag")
	flag.String("hostport", "1443", "Expose smr on hostport")
	flag.String("overlayport", "", "Expose overlay on port")
	flag.String("etcdport", "2379", "Etcd client port listen and advertise")

	flag.String("log", "info", "Log level: debug, info, warn, error, dpanic, panic, fatal")
	flag.String("domains", "localhost", "Comma separated list of the domains to add to the certs")
	flag.String("ips", "127.0.0.1", "Comma separated list of the IPs to add to the certs")

	// Flannel configuration
	flag.String("fbackend", "wireguard", "Flannel backend: vxlan, wireguard")
	flag.String("fcidr", "10.10.0.0/16", "Flannel overlay network CIDR")
	flag.String("fiface", "", "Network interface for flannel to use, if ommited default gateway will be used")

	flag.String("platform", static.PLATFORM_DOCKER, "Container engine name. Supported: [docker]")
	flag.String("node", "", "Name of the smr agent container")

	// Dynamic configuration (Not preserved in client config.yaml)
	flag.String("api", "", "Reachable Node https://URL:PORT URL")
	flag.String("join", "", "Reachable URL of one member of the cluster")
	flag.String("config", "client", "Name of configuration for specific node (used for node management only)")
	flag.String("entrypoint", "/opt/smr/smr", "Entrypoint for the smr")
	flag.String("args", "create smr --node smr", "args")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	os.Args = append([]string{os.Args[0]}, pflag.Args()...)
}

func UnmarshalFields(v *viper.Viper, keys []string, target interface{}) error {
	sub := viper.New()
	for _, key := range keys {
		sub.Set(key, v.Get(key))
	}
	return sub.Unmarshal(target)
}
