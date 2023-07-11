package shared

// a function to create a virtual machine and return a pointer to it.

// new struct for virtual machine config

type VirtualMachineConfig struct {
	Name              string
	Location          string
	ResourceGroupName string
	VmSize            string
	OsDiskSizeGb      int
	OsType            string
	AdminUsername     string
	AdminPassword     string
	Vnet              *VnetConfig
	Subnet            *SubnetConfig
}

// a function to create a virtualmachine.VirtualMachineConfig and return a pointer to it.
func NewAzVirtualMachineConfig(name string, location string, resourceGroupName string, vmSize string, osDiskSizeGb int, osType string, adminUsername string, adminPassword string, vnet *VnetConfig, subnet *SubnetConfig) *VirtualMachineConfig {

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
	}
}

// a function which takes a virtual machine config and creates a virtualmachine.VirtualMachine
func CreateAzVirtualMachine(stack cdktf.TerraformStack, vm VirtualMachineConfig) virtualmachine.VirtualMachine {
	return virtualmachine.NewVirtualMachine(stack, jsii.String("virtualmachine"),
		&virtualmachine.VirtualMachineConfig{
			Name:              jsii.String(vm.Name),
			Location:          jsii.String(vm.Location),
			ResourceGroupName: jsii.String(vm.ResourceGroupName),
			VmSize:            jsii.String(vm.VmSize),
			OsDisk: []*virtualmachine.VirtualMachineOsDisk{
				{
					Name: jsii.String("osdisk"),
					Caching: func() *string {
						caching := "ReadWrite"
						return &caching
					}(),
					CreateOption: func() *string {
						createOption := "FromImage"
						return &createOption
					}(),
					DiskSizeGb: jsii.Number(vm.OsDiskSizeGb),
					ManagedDiskType: func() *string {
						managedDiskType := "Standard_LRS"
						return &managedDiskType
					}(),
					OsType: jsii.String(vm.OsType),
				},
			},
			StorageImageReference: []*virtualmachine.VirtualMachineStorageImageReference{
				{
					Id: jsii.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/terraform-vm-win-linux/providers/Microsoft.Compute/images/terraform-vm-win-linux-image"),
				},
			},
			StorageOsDisk: []*virtualmachine.VirtualMachineStorageOsDisk{
				{
					Name: jsii.String("osdisk"),
					Caching: func() *string {
						caching := "ReadWrite"
						return &caching
					}(),
					CreateOption: func() *string {
						createOption := "FromImage"
						return &createOption
					}(),
					DiskSizeGb: jsii.Number(vm.OsDiskSizeGb),
					ManagedDiskType: func() *string {
						managedDiskType := "Standard_LRS"
						return &managedDiskType
					}
				},
			},
			OsProfile: []*virtualmachine.VirtualMachineOsProfile{
				{
					AdminUsername: jsii.String(vm.AdminUsername),
					AdminPassword: jsii.String(vm.AdminPassword),
					ComputerName:  jsii.String(vm.Name),
				},
			},
			NetworkInterface: []*virtualmachine.VirtualMachineNetworkInterface{
				{
					Name: jsii.String("nic"),
					IpConfiguration: []*virtualmachine.VirtualMachineNetworkInterfaceIpConfiguration{
						{
							Name: jsii.String("ipconfig"),
							Primary: func() *bool {
								primary := true
								return &primary
							},
							SubnetId: jsii.String(vm.Subnet.Id),
						},
					},
				},
			},
		},
	)
}
