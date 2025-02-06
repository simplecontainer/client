package node

import (
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
	"github.com/simplecontainer/smr/pkg/kinds/containers/platforms"
)

type Node struct {
	Name       string
	Home       string
	Definition *v1.ContainersDefinition
	Container  platforms.IPlatform
	Platform   string
}
