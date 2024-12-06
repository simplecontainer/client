package bootstrap

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/logger"
	"github.com/simplecontainer/client/pkg/static"
	"os"
)

func CreateDirectoryTree(projectDir string) {
	for _, path := range static.CLIENT_STRUCTURE {
		dir := fmt.Sprintf("%s/%s", projectDir, path)

		err := os.MkdirAll(dir, 0750)

		if err != nil {
			logger.Log.Fatal(err.Error())

			err = os.RemoveAll(projectDir)
			if err != nil {
				logger.Log.Fatal(err.Error())
			}
		}
	}
}
