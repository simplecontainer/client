package definitions

import (
	"github.com/simplecontainer/smr/pkg/definitions/commonv1"
	v1 "github.com/simplecontainer/smr/pkg/definitions/v1"
)

func FlannelDefinition(subnetCIDR string) *v1.NetworkDefinition {
	return &v1.NetworkDefinition{
		Meta: commonv1.Meta{
			Group: "internal",
			Name:  "cluster",
		},
		Spec: v1.NetworkSpec{
			Driver:          "bridge",
			IPV4AddressPool: subnetCIDR,
		},
	}
}
