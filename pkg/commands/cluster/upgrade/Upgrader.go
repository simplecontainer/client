package upgrade

import (
	"fmt"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/node"
	"time"
)

func Upgrader(mgr *manager.Manager, n1 *node.Node, n2 *node.Node) error {
	err := n1.Stop()

	if err != nil {
		return err
	}

	err = n1.Wait("exited")

	if err != nil {
		fmt.Println("Wait for exited state failed. Falling back to sleep 30s. Reason: ", err)
		time.Sleep(30 * time.Second)
	}

	err = n1.Rename(fmt.Sprintf("%s-%s-backup", n1.Name, n1.Container.GetId()))

	if err != nil {
		fmt.Println("Rename failed starting old container again. Reason: ", err)

		err = n1.Start()

		if err != nil {
			fmt.Println("Starting of old container failed. Manual intervention needed! Reason: ", err)
			return err
		}
	}

	err = n2.Run()

	if err != nil {
		return err
	}

	err = n2.Wait("running")

	if err != nil {
		fmt.Println("Wait for running state failed. Falling back to sleep 30s. Reason: ", err)
		time.Sleep(30 * time.Second)
	}

	mgr.Context.ConnectionTest()
	return nil
}
