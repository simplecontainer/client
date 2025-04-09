package node

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/simplecontainer/client/pkg/configuration"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/smr/pkg/definitions/commonv1"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
	"github.com/simplecontainer/smr/pkg/kinds/containers/platforms/engines/docker"
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
		Platform:   config.Static.Platform,
	}

	var err error

	switch config.Static.Platform {
	case static.PLATFORM_DOCKER:
		if err = docker.IsDaemonRunning(); err != nil {
			helpers.ExitWithErr(err)
		}

		node.Container, err = docker.New(config.Node, node.Definition)
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

func Definition(name string, config *configuration.Configuration) *v1.ContainersDefinition {
	container := &v1.ContainersDefinition{
		Meta: commonv1.Meta{
			Name:   name,
			Group:  "internal",
			Labels: nil,
		},
		Spec: v1.ContainersInternal{
			Image: config.Image,
			Tag:   config.Tag,
			Envs: []string{
				fmt.Sprintf("LOG_LEVEL=%s", config.Static.LogLevel),
			},
			Entrypoint: strings.Split(config.Entrypoint, " "),
			Args:       strings.Split(config.Args, " "),
			Ports: []v1.ContainersPort{
				{
					Container: "1443",
					Host:      config.Static.HostPort,
				},
				{
					Container: "2379",
					Host:      fmt.Sprintf("127.0.0.1:%s", config.Static.EtcdPort),
				},
				{
					Container: "9212",
					Host:      config.Static.OverlayPort,
				},
			},
			Volumes: []v1.ContainersVolume{
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
	}

	return container
}

func (node *Node) Start() error {
	err := node.Container.Start()

	return err
}

func (node *Node) Run() error {
	return node.Container.Run()
}

func (node *Node) Wait(desired string) error {
	ch := make(chan bool, 1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	go node.Check(desired, ch)

	for {
		select {
		case <-ctx.Done():
			return errors.New("timed out waiting for the desired state")
		case <-ch:
			return nil
		}
	}
}

func (node *Node) Check(desired string, ch chan bool) {
	for {
		state, _ := node.Container.GetState()

		if desired == state.State {
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

func Directory(name string, homedir string) error {
	smrDir := fmt.Sprintf("%s/.%s", homedir, name)
	return os.MkdirAll(smrDir, 0750)
}
