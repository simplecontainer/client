package node

import (
	"fmt"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
	"github.com/spf13/viper"
	"strings"
)

func AgentDefinition() *v1.ContainerDefinition {
	return &v1.ContainerDefinition{
		Meta: v1.ContainerMeta{
			Name:   viper.GetString("agent"),
			Group:  "smr",
			Labels: nil,
		},
		Spec: v1.ContainerSpec{
			Container: v1.ContainerInternal{
				Image: viper.GetString("image"),
				Tag:   viper.GetString("tag"),
				Envs: []string{
					fmt.Sprintf("DOMAIN=%s", viper.GetString("domains")),
					fmt.Sprintf("EXTERNALIP=%s", viper.GetString("ips")),
					fmt.Sprintf("HOMEDIR=%s", viper.GetString("homedir")),
				},
				Entrypoint: strings.Split(viper.GetString("entrypoint"), " "),
				Args:       strings.Split(viper.GetString("args"), " "),
				Ports: []v1.ContainerPort{
					{
						Container: "1443",
						Host:      viper.GetString("hostport"),
					},
					{
						Container: "2379",
						Host:      "127.0.0.1:2379",
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
						HostPath:   fmt.Sprintf("~/.%s", viper.GetString("agent")),
						MountPoint: "/home/smr-agent/smr",
					},
					{
						Name:       "ssh",
						Type:       "bind",
						HostPath:   "~/.ssh",
						MountPoint: "/home/smr-agent/.ssh",
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
}