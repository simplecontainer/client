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
	"github.com/simplecontainer/smr/pkg/network"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"net/http"
	"os"
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

	// Remove on start
	os.Remove("/run/flannel/subnet.env")

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

	recursion := make(map[string]string)

	for {
		select {
		case watchResp, ok := <-watcher:
			if ok {
				for _, event := range watchResp.Events {
					switch event.Type {
					case mvccpb.PUT:
						if strings.Contains(string(event.Kv.Key), "subnet") {
							if event.Kv.Lease != 0 {
								var subnet = Subnet{}
								err = json.Unmarshal(event.Kv.Value, &subnet)

								if err != nil {
									return err
								}

								fmt.Println("got subnet ", subnet)

								switch netMode {
								case ipv4:
									fmt.Println("checking if mine  ", config.Flannel.InterfaceFlannel.ExtAddr.String(), " == ", subnet.PublicIP)

									if config.Flannel.InterfaceFlannel.ExtAddr.String() == subnet.PublicIP {
										fmt.Println("adding it as my own subnet", string(event.Kv.Key))

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

								if recursion[string(event.Kv.Key)] == string(event.Kv.Value) {
									continue
								}

								recursion[string(event.Kv.Key)] = string(event.Kv.Value)
								response := network.Send(smrCtx.Client, fmt.Sprintf("%s/api/v1/key/propose/%s", smrCtx.ApiURL, event.Kv.Key), http.MethodPost, event.Kv.Value)

								if response.Success {
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
								} else {
									fmt.Println("flannel failed to inform members about subnet decision - abort startup")
									os.Exit(1)
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
