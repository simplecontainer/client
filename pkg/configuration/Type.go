package configuration

type Configuration struct {
	Environment *Environment

	Static  Static
	Flannel *Flannel

	Id   uint64 `yaml:"-"`
	Node string `yaml:"-"`
	API  string `yaml:"-"`
	Peer string `yaml:"-"`

	Image      string `yaml:"-"`
	Tag        string `yaml:"-"`
	Entrypoint string `yaml:"-"`
	Args       string `yaml:"-"`

	Y bool   `yaml:"-"`
	F bool   `yaml:"-"`
	O string `yaml:"-"`
	W string `yaml:"-"`
	G string `yaml:"-"`
}

type Static struct {
	Platform string
	Node     string
	Context  string

	LogLevel string
	Domains  string
	IPs      string

	Image string
	Tag   string

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
