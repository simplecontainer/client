package flannel

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/simplecontainer/client/pkg/commands/objects/apply"
	"github.com/simplecontainer/client/pkg/configuration"
	smrContext "github.com/simplecontainer/client/pkg/context"
	"github.com/simplecontainer/client/pkg/definitions"
	"github.com/simplecontainer/client/pkg/logger"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	emptyIPv6Network = "::/0"

	ipv4 = iota
	ipv6
)

func Run(ctx context.Context, smrCtx *smrContext.Context, config *configuration.Configuration, agent string) error {
	logger.LogFlannel.Info("starting flannel with backend", zap.String("backend", config.Flannel.Backend))

	netMode, err := findNetMode(config.Flannel.CIDR)
	if err != nil {
		return errors.Wrap(err, "failed to check netMode for flannel")
	}

	go func() {
		err = flannel(ctx, config, config.Flannel.InterfaceSpecified, config.Flannel.IPv6Masq, netMode)
		if err != nil {
			fmt.Println("flannel exited: %v", zap.Error(err))
		}
	}()

	var cli *clientv3.Client
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		return err
	}

	watcher := cli.Watch(ctx, "/coreos.com/network/subnets", clientv3.WithPrefix())
	fmt.Println("client will wait for flannel to return subnet range")

	for {
		select {
		case watchResp, ok := <-watcher:
			if ok {
				for _, event := range watchResp.Events {
					switch event.Type {
					case mvccpb.PUT:
						if strings.Contains(string(event.Kv.Key), "subnet") {
							// Handle only case when it has lease because request came from local flannel
							if event.Kv.Lease != 0 {
								var subnet = Subnet{}
								err = json.Unmarshal(event.Kv.Value, &subnet)

								if err != nil {
									return err
								}

								go func() {
									var kach <-chan *clientv3.LeaseKeepAliveResponse
									kach, err = cli.KeepAlive(ctx, clientv3.LeaseID(event.Kv.Lease))

									for {
										select {
										case data, ok := <-kach:
											if ok {
												fmt.Println(fmt.Sprintf("keep alived: %s", data.String()))
												break
											} else {
												fmt.Println(fmt.Sprintf("closed keep alive channel for lease: %s", event.Kv.Lease))
												return
											}
										}
									}
								}()

								switch netMode {
								case ipv4:
									if config.Flannel.InterfaceFlannel.ExtAddr.String() == subnet.PublicIP {
										split := strings.Split(string(event.Kv.Key), "/")
										CIDR := strings.Replace(split[len(split)-1], "-", "/", 1)

										NetworkDefinition, _ := definitions.FlannelDefinition(CIDR).ToJson()
										apply.Apply(smrCtx, NetworkDefinition)
									}
									break
								case ipv6:
									if config.Flannel.InterfaceFlannel.ExtV6Addr.String() == subnet.PublicIPv6 {
										split := strings.Split(string(event.Kv.Key), "/")
										CIDR := strings.Replace(split[len(split)-1], "-", "/", 1)

										NetworkDefinition, _ := definitions.FlannelDefinition(CIDR).ToJson()
										apply.Apply(smrCtx, NetworkDefinition)
									}
									break
								case ipv4 | ipv6:
									break
								}
							}
						}
					}
				}
			}
		case <-ctx.Done():
			return errors.New("closed watcher channel - should not block")
		}
	}
}
