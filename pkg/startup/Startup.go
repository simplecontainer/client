package startup

import (
	"flag"
	"fmt"
	"github.com/flannel-io/flannel/pkg/ip"
	"github.com/simplecontainer/client/pkg/configuration"
	"github.com/simplecontainer/smr/pkg/static"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net"
	"os"
	"runtime"
)

func Load(configObj *configuration.Configuration, projectDir string) {
	configObj.Environment = GetEnvironmentInfo()
	ReadFlags(configObj)

	configObj.Flannel = &configuration.Flannel{
		Backend:            viper.GetString("fbackend"),
		CIDR:               make([]*net.IPNet, 0),
		InterfaceSpecified: nil,
		EnableIPv4:         viper.GetBool("fenableIPv4"),
		EnableIPv6:         viper.GetBool("fenableIPv6"),
		IPv6Masq:           viper.GetBool("fmaskIPv6"),
		//
		ConfigFile:       "/run/flannel/subnet.env",
		Network:          ip.IP4Net{},
		Networkv6:        ip.IP6Net{},
		InterfaceFlannel: nil,
	}

	var CIDR *net.IPNet
	var err error
	_, CIDR, err = net.ParseCIDR(viper.GetString("fcidr"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configObj.Flannel.CIDR = []*net.IPNet{CIDR}
	configObj.Flannel.InterfaceSpecified, _ = net.InterfaceByName(viper.GetString("finterface"))

	if runtime.GOOS != "windows" {
		if viper.GetBool("fenableIPv4") {
			if _, err = os.Stat("/proc/sys/net/bridge/bridge-nf-call-iptables"); os.IsNotExist(err) {
				fmt.Println("Failed to check br_netfilter: ", zap.Error(err))
				os.Exit(1)
			}
		}

		if viper.GetBool("fenableIPv6") {
			if _, err = os.Stat("/proc/sys/net/bridge/bridge-nf-call-ip6tables"); os.IsNotExist(err) {
				fmt.Println("Failed to check br_netfilter: ", zap.Error(err))
				os.Exit(1)
			}
		}
	}
}

func ReadFlags(configObj *configuration.Configuration) {
	/* Operation mode */
	flag.String("context", "", "Context name")
	flag.Bool("y", false, "Say yes to everything")

	HOMEDIR, _ := os.UserHomeDir()

	// Cli flags
	flag.Bool("wait", false, "Wait for container to exit")
	flag.Bool("f", false, "Follow logs")
	flag.String("o", "d", "Output type: d(efault),s(hort)")

	// Node flags
	flag.String("image", "quay.io/simplecontainer/smr", "The smr image repo")
	flag.String("tag", "latest", "The smr image tag")
	flag.String("entrypoint", "/opt/smr/smr", "Entrypoint for the smr")
	flag.String("args", "create smr --agent smr-agent", "args")
	flag.String("hostport", "1443", "Expose smr on hostport")
	flag.String("overlayport", "", "Expose overlay on port")
	flag.String("etcdport", "2379", "Etcd client port listen and advertise")

	flag.String("log", "info", "Log level: debug, info, warn, error, dpanic, panic, fatal")
	flag.String("domains", "localhost", "Comma separated list of the domains to add to the certs")
	flag.String("ips", "127.0.0.1", "Comma separated list of the IPs to add to the certs")
	flag.String("homedir", HOMEDIR, "Host homedir")

	// Flannel configuration
	flag.String("fbackend", "wireguard", "Flannel backend: vxlan, wireguard")
	flag.String("fcidr", "10.10.0.0/16", "Flannel overlay network CIDR")
	flag.String("fiface", "", "Network interface for flannel to use, if ommited default gateway will be used")

	flag.String("platform", static.PLATFORM_DOCKER, "Container engine name. Supported: [docker]")
	flag.String("name", "smr-agent", "Name of the smr agent container")

	// Cluster configuration
	flag.String("node", "", "Reachable Node https://URL:PORT URL")
	flag.String("join", "", "Reachable URL of one member of the cluster")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	configObj.Flags.Context = viper.GetString("context")
	configObj.Flags.Y = viper.GetBool("y")
}
