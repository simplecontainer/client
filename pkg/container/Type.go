package container

import (
	"smr/pkg/definitions"
	"smr/pkg/network"
)

type Container struct {
	Static  Static
	Runtime Runtime
	Exports Exports
	Status  Status
}

type Static struct {
	Name                   string
	GeneratedName          string
	GeneratedNameNoProject string
	Group                  string
	Image                  string
	Tag                    string
	Replicas               int
	Networks               []string
	Env                    []string
	MappingFiles           []map[string]string
	MappingPorts           []network.PortMappings
	ExposedPorts           []string
	MountFiles             []string
	Definition             definitions.Container
}

type Runtime struct {
	Auth          string
	Id            string
	Networks      map[string]Network
	State         string
	FoundRunning  bool
	FirstObserved bool
	Ready         bool
	Configuration map[string]any
	Resources     []Resource
}

type Network struct {
	NetworkId string
	IP        string
}

type Status struct {
	DependsSolved  bool
	BackOffRestart bool
	Healthy        bool
	Ready          bool
	Running        bool
	Reconciling    bool
}

type Resource struct {
	Identifier string
	Key        string
	Data       map[string]string
	MountPoint string
}

type Exports struct {
	path string
}

type ExecResult struct {
	Stdout string
	Stderr string
	Exit   int
}
