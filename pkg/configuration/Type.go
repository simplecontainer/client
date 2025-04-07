package configuration

type Configuration struct {
	Environment *Environment
	Flannel     *Flannel

	Flags   Flags `yaml:"-"`
	Setup   Setup
	Network Network
	Dynamic Dynamic `yaml:"-"`
}

type Flags struct {
	Y bool   `yaml:"-"`
	F bool   `yaml:"-"`
	O string `yaml:"-"`
	W string `yaml:"-"`
	G string `yaml:"-"`
}

type Setup struct {
	Platform string
	Node     string
	Context  string

	LogLevel string
	Domains  string
	IPs      string

	Image      string
	Tag        string
	Entrypoint string `yaml:"-"`
	Args       string `yaml:"-"`

	HostPort    string
	OverlayPort string
	EtcdPort    string
}

type Flannel struct {
	Backend            string
	CIDR               string
	InterfaceSpecified string
	EnableIPv4         bool
	EnableIPv6         bool
	IPv6Masq           bool
}

type Environment struct {
	Home            string
	LogsDirectory   string
	ClientDirectory string
}

type Network struct {
	Fbackend   string
	Fcidr      string
	Finterface string
}

type Dynamic struct {
	Node string
	API  string
	Join string
}
