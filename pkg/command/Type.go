package command

import "github.com/simplecontainer/client/pkg/manager"

type Command struct {
	Name      string
	Flag      string
	Condition func(*manager.Manager) bool
	Functions []func(*manager.Manager, []string)
	DependsOn []func(*manager.Manager, []string)
}
