package shared

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	networkinterface "cdk.tf/go/stack/generated/azurerm/networkinterface"
	virtualmachine "cdk.tf/go/stack/generated/azurerm/virtualmachine"
)

// a function to create a virtual machine and return a pointer to it.

// new struct for virtual machine config
type VirtualMachineConfig struct {
	Name                        string
	Location                    string
	ResourceGroupName           string
	VmSize                      string
	OsDiskSizeGb                int
	OsType                      string
	AdminUsername               string
	AdminPassword               string
	Vnet                        *VnetConfig
	Subnet                      *SubnetConfig
	NicId                       string
	StorageOsDiskConfig         *VirtualMachineStorageOsDiskConfig
	StorageImageReferenceConfig *VirtualMachineStorageImageReferenceConfig
}

// a struct for network interface config
type NetworkInterfaceConfig struct {
	Name              string
	Location          string
	ResourceGroupName string
	Vnet              *VnetConfig
	Subnet            *SubnetConfig
}

// virtual machine storage os disk struct
type VirtualMachineStorageOsDiskConfig struct {
	Name            string
	Caching         string
	CreateOption    string
	DiskSizeGb      int
	ManagedDiskType string
}

// Storage Image Reference struct
type VirtualMachineStorageImageReferenceConfig struct {
	Publisher string
	Offer     string
	Sku       string
	Version   string
}

// a function to initialize a VirtualMachineConfig struct
func (vmc *VirtualMachineConfig) Init(name string, location string, resourceGroupName string, vmSize string, osDiskSizeGb int, osType string, adminUsername string, adminPassword string, vnet *VnetConfig, subnet *SubnetConfig, nicId string) {
	vmc.Name = name
	vmc.Location = location
	vmc.ResourceGroupName = resourceGroupName
	vmc.VmSize = vmSize
	vmc.OsDiskSizeGb = osDiskSizeGb
	vmc.OsType = osType
	vmc.AdminUsername = adminUsername
	vmc.AdminPassword = adminPassword
	vmc.Vnet = vnet
	vmc.Subnet = subnet
	vmc.NicId = nicId
}

// a function to initialize a NetworkInterfaceConfig struct
func (nic *NetworkInterfaceConfig) Init(name string, location string, resourceGroupName string, vnet *VnetConfig, subnet *SubnetConfig) {
	nic.Name = name
	nic.Location = location
	nic.ResourceGroupName = resourceGroupName
	nic.Vnet = vnet
	nic.Subnet = subnet
}

// a function to initialize a VirtualMachineStorageOsDisk struct
func (vmos *VirtualMachineStorageOsDiskConfig) Init(name string, caching string, createOption string, diskSizeGb int, managedDiskType string) {
	vmos.Name = name
	vmos.Caching = caching
	vmos.CreateOption = createOption
	vmos.DiskSizeGb = diskSizeGb
	vmos.ManagedDiskType = managedDiskType
}

// a function to initialize a VirtualMachineStorageImageReference struct
func (vmir *VirtualMachineStorageImageReferenceConfig) Init(publisher string, offer string, sku string, version string) {
	vmir.Publisher = publisher
	vmir.Offer = offer
	vmir.Sku = sku
	vmir.Version = version
}

// a function to create a virtualmachine.VirtualMachineConfig and return a pointer to it.
func NewAzVirtualMachineConfig(name string, location string, resourceGroupName string, vmSize string, osDiskSizeGb int, osType string, adminUsername string, adminPassword string, vnet *VnetConfig, subnet *SubnetConfig, nicId string) *VirtualMachineConfig {

	return &VirtualMachineConfig{
		Name:              name,
		Location:          location,
		ResourceGroupName: resourceGroupName,
		VmSize:            vmSize,
		OsDiskSizeGb:      osDiskSizeGb,
		OsType:            osType,
		AdminUsername:     adminUsername,
		AdminPassword:     adminPassword,
		Vnet:              vnet,
		Subnet:            subnet,
		NicId:             nicId,
	}
}

// a function to create a network interface and return the id
func NewAzNetworkInterfaceConfig(name string, location string, resourceGroupName string, vnet *VnetConfig, subnet *SubnetConfig) *NetworkInterfaceConfig {

	return &NetworkInterfaceConfig{
		Name:              name,
		Location:          location,
		ResourceGroupName: resourceGroupName,
		Vnet:              vnet,
		Subnet:            subnet,
	}
}

// a function which takes a network interface config and creates a virtualmachine.NetworkInterface and returns the id
func CreateAzNetworkInterface(stack cdktf.TerraformStack, nic NetworkInterfaceConfig) string {
	nicID := networkinterface.NewNetworkInterface(stack, jsii.String("networkinterface"),
		&networkinterface.NetworkInterfaceConfig{
			Name:              jsii.String(nic.Name),
			Location:          jsii.String(nic.Location),
			ResourceGroupName: jsii.String(nic.ResourceGroupName),
			IpConfiguration: []*networkinterface.NetworkInterfaceIpConfiguration{
				{
					Name:                       jsii.String(nic.Name),
					PrivateIpAddressAllocation: jsii.String("Dynamic"),
					Primary:                    jsii.Bool(true),
				},
			},
		},
	)
	n := fmt.Sprintf("%v", nicID.Id())
	return n
}

// a function which takes a virtual machine config and creates a virtualmachine.VirtualMachine
func CreateAzVirtualMachine(stack cdktf.TerraformStack, vm VirtualMachineConfig) virtualmachine.VirtualMachine {

	// convert networkinterface.NetworkInterface to *[]*string
	nicIdSlice := StringSlice(vm.NicId)

	return virtualmachine.NewVirtualMachine(stack, jsii.String("virtualmachine"),
		&virtualmachine.VirtualMachineConfig{
			Name:              jsii.String(vm.Name),
			Location:          jsii.String(vm.Location),
			ResourceGroupName: jsii.String(vm.ResourceGroupName),
			VmSize:            jsii.String(vm.VmSize),
			OsProfile: &virtualmachine.VirtualMachineOsProfile{
				ComputerName:  jsii.String(vm.Name),
				AdminUsername: jsii.String(vm.AdminUsername),
				AdminPassword: jsii.String(vm.AdminPassword),
			},
			NetworkInterfaceIds: nicIdSlice,
			StorageOsDisk: &virtualmachine.VirtualMachineStorageOsDisk{
				Name:            jsii.String(vm.StorageOsDiskConfig.Name),
				Caching:         jsii.String(vm.StorageOsDiskConfig.Caching),
				CreateOption:    jsii.String(vm.StorageOsDiskConfig.CreateOption),
				DiskSizeGb:      jsii.Number(vm.StorageOsDiskConfig.DiskSizeGb),
				ManagedDiskType: jsii.String(vm.StorageOsDiskConfig.ManagedDiskType),
			},
			StorageImageReference: &virtualmachine.VirtualMachineStorageImageReference{
				Publisher: jsii.String(vm.StorageImageReferenceConfig.Publisher),
				Offer:     jsii.String(vm.StorageImageReferenceConfig.Offer),
				Sku:       jsii.String(vm.StorageImageReferenceConfig.Sku),
				Version:   jsii.String(vm.StorageImageReferenceConfig.Version),
			},
		},
	)
}
