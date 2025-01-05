package configuration

import (
	"github.com/flannel-io/flannel/pkg/backend"
	"github.com/flannel-io/flannel/pkg/ip"
	"net"
)

type Configuration struct {
	Target      string `default:"development" json:"target"`
	Root        string `json:"root"`
	Log         string `json:"log"`
	Environment *Environment
	Flags       Flags
	Flannel     *Flannel
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
	HOMEDIR    string
	FLANNELDIR string
	ROOTDIR    string
	LOGDIR     string
	CLIENTIP   string
}
