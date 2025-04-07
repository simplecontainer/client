package upgrader

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/cluster/upgrade"
	"github.com/simplecontainer/client/pkg/logger"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/node"
	"github.com/simplecontainer/smr/pkg/upgrader"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func Upgrader(mgr *manager.Manager) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("localhost:%s", mgr.Configuration.Setup.EtcdPort)},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		return err
	}

	fmt.Println("WATCHING UPGRADESSSS")

	watcher := cli.Watch(context.Background(), "/smr/upgrade", clientv3.WithPrefix())

	for {
		select {
		case watchResp, ok := <-watcher:
			if ok {
				fmt.Println("XXXXXXXXXX")
				var u *upgrader.Upgrade

				err = json.Unmarshal(watchResp.Events[0].Kv.Value, &u)

				if err != nil {
					logger.Log.Error(err.Error())
					break
				}

				var n *node.Node
				n, err = node.New(mgr.Configuration.Dynamic.Node, mgr.Configuration)

				if err != nil {
					panic(err)
				}

				fmt.Println("GOT UPGRADE REQUEST")

				err = upgrade.Upgrade(n, u.Image, u.Tag)

				if err != nil {
					fmt.Println(err)
					break
				}

				fmt.Println("start cluster again")
			}
			break
		}
	}
}
