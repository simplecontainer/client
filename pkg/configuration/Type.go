package configuration

import (
	"github.com/flannel-io/flannel/pkg/backend"
	"github.com/flannel-io/flannel/pkg/ip"
	"net"
)

type Configuration struct {
	Environment *Environment
	Flannel     *Flannel

	Startup Startup
}

type Flannel struct {
	Backend            string
	CIDR               []*net.IPNet
	InterfaceSpecified *net.Interface
	EnableIPv4         bool
	EnableIPv6         bool
	IPv6Masq           bool
	ConfigFile         string
	InterfaceFlannel   *backend.ExternalInterface
	Network            ip.IP4Net
	Networkv6          ip.IP6Net
}

type Flags struct {
	Context string
	Y       bool
}

type Environment struct {
	Home            string
	LogsDirectory   string
	ClientDirectory string
}

type Startup struct {
	Platform string
	Name     string
	Context  string

	LogLevel string
	Domains  string
	IPs      string

	Y bool
	F bool
	O string
	W string
	G string

	Image      string
	Tag        string
	Entrypoint string
	Args       string

	HostPort    string
	OverlayPort string
	EtcdPort    string

	Node string
	Join string

	Fbackend   string
	Fcidr      string
	Finterface string
}
