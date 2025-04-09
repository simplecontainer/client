package bootstrap

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/simplecontainer/client/pkg/static"
	"os"
)

func CreateDirectoryTree(projectDir string) {
	for _, path := range static.ClientStructure {
		dir := fmt.Sprintf("%s/%s", projectDir, path)

		err := os.MkdirAll(dir, 0750)

		if err != nil {
			glog.Fatal(err.Error())

			err = os.RemoveAll(projectDir)
			if err != nil {
				glog.Fatal(err.Error())
			}
		}
	}
}
