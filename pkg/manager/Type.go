package manager

import (
	"github.com/qdnqn/smr-client/pkg/context"
	"github.com/qdnqn/smr/pkg/config"
	"github.com/qdnqn/smr/pkg/runtime"
)

type Manager struct {
	Config  *config.Config
	Runtime *runtime.Runtime
	Context *context.Context
}
