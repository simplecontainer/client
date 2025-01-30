package node

import (
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms"
)

type Node struct {
	Name       string
	Home       string
	Definition *v1.ContainerDefinition
	Container  platforms.IPlatform
	Platform   string
}
