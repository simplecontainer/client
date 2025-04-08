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
		"api", "join", "id", "node", "entrypoint", "args", "image", "tag", "w", "f", "o", "y", "g", "it", "c",
	}, &configObj)

	if configObj.Image == "" {
		configObj.Image = configObj.Static.Image
	}

	if configObj.Tag == "" {
		configObj.Tag = configObj.Static.Tag
	}
}

func LoadFromFlags(configObj *configuration.Configuration) {
	err := viper.Unmarshal(configObj)

	if err != nil {
		panic(err)
	}
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

	path := fmt.Sprintf("%s/%s/%s/%s.yaml", configObj.Environment.Home, static.ROOTSMR, static.CONFIGDIR, configObj.Node)

	err = os.WriteFile(path, yamlObj, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadFlags() {
	// Node setup flags
	flag.String("static.image", "quay.io/simplecontainer/smr", "The smr image repo")
	flag.String("static.tag", "latest", "The smr image tag")
	flag.String("static.hostport", "1443", "Expose smr on hostport")
	flag.String("static.overlayport", "", "Expose overlay on port")
	flag.String("static.etcdport", "2379", "Etcd client port listen and advertise")

	flag.String("static.log", "info", "Log level: debug, info, warn, error, dpanic, panic, fatal")
	flag.String("static.domains", "localhost", "Comma separated list of the domains to add to the certs")
	flag.String("static.ips", "127.0.0.1", "Comma separated list of the IPs to add to the certs")

	// Flannel configuration
	flag.String("flannel.backend", "wireguard", "Flannel backend: vxlan, wireguard")
	flag.String("flannel.cidr", "10.10.0.0/16", "Flannel overlay network CIDR")
	flag.String("flannel.iface", "", "Network interface for flannel to use, if ommited default gateway will be used")

	flag.String("static.platform", static.PLATFORM_DOCKER, "Container engine name. Supported: [docker]")

	// Dynamic configuration (Not preserved in client config.yaml)
	flag.String("context", "", "Context")
	flag.String("container", "main", "Which container to stream main or init?")
	flag.Uint64("id", 0, "Id of the node")
	flag.String("node", "", "Name of the smr agent container")
	flag.String("api", "", "Reachable Node https://URL:PORT URL")
	flag.String("join", "", "Reachable URL of one member of the cluster")
	flag.String("config", "client", "Name of configuration for specific node (used for node management only)")
	flag.String("image", "", "Image to run")
	flag.String("tag", "", "Image to run")
	flag.String("entrypoint", "/opt/smr/smr", "Entrypoint for the smr")
	flag.String("args", "create smr --node smr", "args")

	// Dynamic - Cli flags
	flag.String("w", "", "Wait for container to be in defined state")
	flag.Bool("f", false, "Follow logs")
	flag.String("o", "d", "Output type: d(efault),s(hort)")
	flag.Bool("y", false, "Say yes to everything")
	flag.String("g", "default", "Group")
	flag.Bool("it", false, "Interactive exec")
	flag.String("c", "", "Command for exec")

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
