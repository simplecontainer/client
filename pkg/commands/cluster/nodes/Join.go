package nodes

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/cluster"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/upgrader"
	"os"
)

func Join(mgr *manager.Manager) {
	go func() {
		err := upgrader.Upgrader(mgr)

		if err != nil {
			fmt.Println("failed to start control watcher")
			os.Exit(3)
		}
	}()

	cluster.Join(mgr)
}
