package manager

import (
	"smr/pkg/runtime"
)
import "smr/pkg/config"

type Manager struct {
	Config  *config.Config
	Runtime runtime.Runtime
}
