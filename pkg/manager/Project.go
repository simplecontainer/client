package manager

import (
	"fmt"
	"smr/pkg/bootstrap"
	"smr/pkg/logger"
	"smr/pkg/static"
)

func (mgr *Manager) CreateProject(projectName string) {
	if projectName == "" {
		panic("Project name cannot be empty")
	}

	projectDir := fmt.Sprintf("%s/%s/%s", mgr.Runtime.HOMEDIR, static.SMR, projectName)

	bootstrap.CreateDirectoryTree(projectDir)
	config := bootstrap.GenerateConfigProject(projectDir)

	if !bootstrap.WriteConfiguration(config, projectDir, projectName) {
		logger.Log.Fatal("failed to create new project")
	}
}

func (mgr *Manager) DeleteProject(projectName string) {
	if projectName == "" {
		projectName = static.SMR
	}

	projectDir := fmt.Sprintf("%s/%s/%s", mgr.Runtime.HOMEDIR, static.SMR, projectName)

	bootstrap.ClearDirectoryTree(projectDir)
}
