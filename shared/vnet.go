// a function to create a vnet and return a pointer to it.

package shared

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	virtualnetwork "cdk.tf/go/stack/generated/azurerm/virtualnetwork"
)

// new struct for vnet config
type VnetConfig struct {
	Name              string
	AddressSpace      string
	Location          string
	ResourceGroupName string
	Subnet            *SubnetConfig
}

// new struct for subnet config
type SubnetConfig struct {
	Name          string
	AddressPrefix string
}

// a function to create a virtualnetwork.VirtualNetworkConfig and return a pointer to it.
func (vnonf *VnetConfig) Init(subnetName string, vnetName string, vnetAddressSpace string, location string, resourceGroupName string, subnetAddressSpace string) {
	vnonf.Name = vnetName
	vnonf.AddressSpace = vnetAddressSpace
	vnonf.Location = location
	vnonf.ResourceGroupName = resourceGroupName
	vnonf.Subnet = &SubnetConfig{
		Name:          subnetName,
		AddressPrefix: subnetAddressSpace,
	}
}

// a function which takes a vnet config and creates a virtualnetwork.VirtualNetwork
func CreateAzVirtualNetwork(stack cdktf.TerraformStack, vnet VnetConfig) virtualnetwork.VirtualNetwork {
	return virtualnetwork.NewVirtualNetwork(stack, jsii.String("virtualnetwork"),
		&virtualnetwork.VirtualNetworkConfig{
			Name:              jsii.String(vnet.Name),
			AddressSpace:      StringSlice(vnet.AddressSpace),
			Location:          jsii.String(vnet.Location),
			ResourceGroupName: jsii.String(vnet.ResourceGroupName),
			Subnet: []*virtualnetwork.VirtualNetworkSubnet{
				{
					Name:          jsii.String(vnet.Subnet.Name),
					AddressPrefix: jsii.String(vnet.Subnet.AddressPrefix),
				},
			},
		},
	)
}
