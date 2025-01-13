package command

import "github.com/simplecontainer/client/pkg/manager"

func (command Command) GetName() string {
	return command.Name
}

func (command Command) GetCondition(mgr *manager.Manager) bool {
	return command.Condition(mgr)
}

func (command Command) GetFunctions() []func(*manager.Manager, []string) {
	return command.Functions
}

func (command Command) GetDependsOn() []func(*manager.Manager, []string) {
	return command.DependsOn
}
