package ps

import (
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms/engines/docker"
	"time"
)

type ContainerInformation struct {
	Group         string
	Name          string
	GeneratedName string
	Image         string
	Tag           string
	IPs           string
	Ports         string
	Dependencies  string
	DockerState   string
	SmrState      string
	LastUpdate    time.Duration
}

type Container struct {
	General *platforms.General
	Type    string
}

type ContainerDocker struct {
	Platform *docker.Docker
	General  *platforms.General
	Type     string
}
