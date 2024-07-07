package manager

import (
	"github.com/simplecontainer/client/pkg/configuration"
	"github.com/simplecontainer/client/pkg/context"
)

type Manager struct {
	Configuration *configuration.Configuration
	Context       *context.Context
}
