package upgrader

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/simplecontainer/client/pkg/cluster"
	"github.com/simplecontainer/client/pkg/commands/cluster/upgrade"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/client/pkg/node"
	"github.com/simplecontainer/smr/pkg/controler"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/net/context"
	"time"
)

func Upgrader(mgr *manager.Manager) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("localhost:%s", mgr.Configuration.Static.EtcdPort)},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		return err
	}

	fmt.Println("listening for control events")

	watcher := cli.Watch(context.Background(), "/smr/control/", clientv3.WithPrefix())

	for {
		select {
		case watchResp, ok := <-watcher:
			if ok {
				for _, event := range watchResp.Events {
					switch event.Type {
					case mvccpb.PUT:
						fmt.Println("new control event")
						var c *controler.Control
						err = json.Unmarshal(watchResp.Events[0].Kv.Value, &c)

						if err != nil {
							glog.Error(err.Error())
							break
						}

						if c.GetUpgrade() != nil {
							// Currently react only to upgrade control (drain and start are ignored)
							mgr.Configuration.Args = "start"

							var n1 *node.Node
							var n2 *node.Node

							n1, err = node.New(mgr.Configuration.Node, mgr.Configuration)

							if err != nil {
								glog.Error(err.Error())
								break
							}

							mgr.Configuration.Image = c.Upgrade.Image
							mgr.Configuration.Tag = c.Upgrade.Tag

							n2, err = node.New(mgr.Configuration.Node, mgr.Configuration)

							if err != nil {
								glog.Error(err.Error())
								break
							}

							err = upgrade.Upgrader(n1, n2)

							if err != nil {
								fmt.Println(err)
								break
							}

							glog.Info("node started again - attempt to join cluster will proceed after node is healthy")

							err = mgr.Context.Connect(true)

							if err != nil {
								fmt.Println(err)
								break
							}

							cluster.ReJoin(mgr)
						}
					}
					break
				}
			}
		}
	}
}
