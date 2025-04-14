package upgrader

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/controler"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func Upgrader(mgr *manager.Manager) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("localhost:%s", mgr.Configuration.Static.EtcdPort)},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		return fmt.Errorf("failed to create etcd client: %w", err)
	}
	defer cli.Close()

	fmt.Println("Listening for control events...")

	watchCh := cli.Watch(context.Background(), "/smr/control/", clientv3.WithPrefix())

	for watchResp := range watchCh {
		for _, event := range watchResp.Events {
			if event.Type != mvccpb.PUT {
				continue
			}

			fmt.Println("New control event received")

			var c controler.Control

			if err := json.Unmarshal(event.Kv.Value, &c); err != nil {
				glog.Errorf("Failed to unmarshal control: %v", err)
				continue
			}

			switch {
			case c.GetDrain() != nil && c.GetUpgrade() == nil:
				if err := handleDrain(mgr); err != nil {
					glog.Errorf("Drain handling failed: %v", err)
				}
				break
			case c.GetUpgrade() != nil:
				if err := handleUpgrade(mgr, c); err != nil {
					glog.Errorf("Upgrade handling failed: %v", err)
				}
				break
			}
		}
	}

	return nil
}
