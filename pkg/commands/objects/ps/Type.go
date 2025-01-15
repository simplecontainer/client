package ps

import (
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
	NodeName      string
	NodeURL       string
	Recreated     bool
	LastUpdate    time.Duration
}
