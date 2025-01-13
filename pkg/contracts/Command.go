package contracts

import "github.com/simplecontainer/client/pkg/manager"

type Command interface {
	GetName() string
	GetCondition(*manager.Manager) bool
	GetFunctions() []func(*manager.Manager, []string)
	GetDependsOn() []func(*manager.Manager, []string)
}
