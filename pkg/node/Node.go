package node

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/simplecontainer/client/pkg/configuration"
	"github.com/simplecontainer/client/pkg/helpers"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
	"github.com/simplecontainer/smr/pkg/kinds/container/platforms/engines/docker"
	"github.com/simplecontainer/smr/pkg/static"
	"golang.org/x/net/context"
	"os"
	"strings"
	"time"
)

func New(name string, config *configuration.Configuration) (*Node, error) {
	node := &Node{
		Name:       name,
		Home:       config.Environment.Home,
		Definition: Definition(name, config),
		Platform:   config.Startup.Platform,
	}

	var err error

	switch config.Startup.Platform {
	case static.PLATFORM_DOCKER:
		if err = docker.IsDaemonRunning(); err != nil {
			helpers.ExitWithErr(err)
		}

		node.Container, err = docker.New(config.Startup.Name, node.Definition)
		break
	default:
		return nil, errors.New("platform not supported")
	}

	if err != nil {
		return nil, err
	} else {
		return node, nil
	}
}

func Definition(name string, config *configuration.Configuration) *v1.ContainerDefinition {
	container := &v1.ContainerDefinition{
		Meta: v1.ContainerMeta{
			Name:   name,
			Group:  "internal",
			Labels: nil,
		},
		Spec: v1.ContainerSpec{
			Container: v1.ContainerInternal{
				Image: config.Startup.Image,
				Tag:   config.Startup.Tag,
				Envs: []string{
					fmt.Sprintf("LOG_LEVEL=%s", config.Startup.LogLevel),
				},
				Entrypoint: strings.Split(config.Startup.Entrypoint, " "),
				Args:       strings.Split(config.Startup.Args, " "),
				Ports: []v1.ContainerPort{
					{
						Container: "1443",
						Host:      config.Startup.HostPort,
					},
					{
						Container: "2379",
						Host:      fmt.Sprintf("127.0.0.1:%s", config.Startup.EtcdPort),
					},
					{
						Container: "9212",
						Host:      config.Startup.OverlayPort,
					},
				},
				Volumes: []v1.ContainerVolume{
					{
						Name:       "docker-socket",
						Type:       "bind",
						HostPath:   "/var/run/docker.sock",
						MountPoint: "/var/run/docker.sock",
					},
					{
						Name:       "smr",
						Type:       "bind",
						HostPath:   fmt.Sprintf("%s/.%s", config.Environment.Home, name),
						MountPoint: "/home/node/smr",
					},
					{
						Name:       "ssh",
						Type:       "bind",
						HostPath:   fmt.Sprintf("%s/.ssh", config.Environment.Home),
						MountPoint: "/home/node/.ssh",
					},
					{
						Name:       "tmp",
						Type:       "bind",
						HostPath:   "/tmp",
						MountPoint: "/tmp",
					},
				},
				Replicas: 1,
				Dns:      []string{"127.0.0.1"},
			},
		},
	}

	return container
}

func (node *Node) Start() error {
	_, err := node.Container.Run()

	return err
}

func (node *Node) Wait(desired string) error {
	ch := make(chan bool, 1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	go node.Check(desired, ch)

	select {
	case <-ctx.Done():
		return errors.New("timed out waiting for the desired state")
	case <-ch:
		return nil
	}
}

func (node *Node) Check(desired string, ch chan bool) {
	for {
		state, _ := node.Container.GetContainerState()

		if desired == state {
			ch <- true
			return
		}
	}
}

func (node *Node) Stop() error {
	return node.Container.Stop(static.SIGTERM)
}

func (node *Node) Rename(name string) error {
	return node.Container.Rename(name)
}

func (node *Node) Directory(name string, homedir string) error {
	smrDir := fmt.Sprintf("%s/.%s", homedir, name)
	return os.MkdirAll(smrDir, 0750)
}
