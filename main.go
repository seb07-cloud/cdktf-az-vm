package main

import (
	"cdk.tf/go/stack/shared"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

const (
	// Azure location to deploy the resources to
	resourceGroupName = "rg-cdktf-vm-win-linux"
	azureLocation     = "westeurope"
	vnetName          = "vnet-cdktf-vm-win-linux"
	subnetName        = "subnet-cdktf-vm-win-linux"
	vnetAddressSpace  = "10.0.0.0/16"
	snetAddressSpace  = "10.0.0.0/24"

	// Azure Virtual Machine based consts
	vmName          = "cdktf-vm-win-linux"
	vmSize          = "Standard_B2ms"
	vmOsDiskSizeGb  = 30
	vmOsType        = "Linux"
	vmAdminUsername = "cdktfuser"
	vmAdminPassword = "cdktfpassword"

	// Azure Storage based consts
	storageOsDisk = "cdktfosdisk"
	publisher     = "Canonical"
	offer         = "UbuntuServer"
	sku           = "18.04-LTS"
	version       = "latest"
)

// Tag Definitions
var tags = &shared.Tags{
	Owner:       "seb@testing.local",
	Service:     "cdktf-testing",
	Environment: "dev",
}

func NewResourceGroupStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// Create tags
	tg := shared.CreateTags(&stack, *tags)

	// Create an Azure provider
	shared.NewAzProvider(stack, "azurerm")

	// Create a resource group
	rgrp := &shared.ResourceGroupConfig{
		Name:     resourceGroupName,
		Location: azureLocation,
		Tags:     &tg,
	}
	shared.CreateAzResourceGroup(stack, *rgrp)

	return stack
}

func NewVnetStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// Create an Azure provider
	shared.NewAzProvider(stack, "azurerm")

	// Create a virtual network + subnet config
	vnc := &shared.VnetConfig{
		Name:              vnetName,
		AddressSpace:      vnetAddressSpace,
		Location:          azureLocation,
		ResourceGroupName: resourceGroupName,
		Subnet: &shared.SubnetConfig{
			Name:          subnetName,
			AddressPrefix: snetAddressSpace,
		},
	}

	// Create a virtual network + subnet
	shared.CreateAzVirtualNetwork(stack, *vnc)

	return stack
}

func NewVmStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// Create an Azure provider
	shared.NewAzProvider(stack, "azurerm")

	// Create a vm with a public ip
	vmic := &shared.NetworkInterfaceConfig{
		Name:              "nic-cdktf-vm-win-linux",
		Location:          azureLocation,
		ResourceGroupName: resourceGroupName,
		Vnet:              vnc,
		// NEED  to pass the subnet id
		Subnet: vnc.Subnet,
	}

	// Create a network interface
	nicId := shared.CreateAzNetworkInterface(stack, *vmic)

	// Create Storage Image Reference
	vmir := &shared.VirtualMachineStorageImageReferenceConfig{
		Publisher: publisher,
		Offer:     offer,
		Sku:       sku,
		Version:   version,
	}

	// Create Storage Os Disk Config
	vmosd := &shared.VirtualMachineStorageOsDiskConfig{
		Name:            storageOsDisk,
		Caching:         "ReadWrite",
		CreateOption:    "FromImage",
		DiskSizeGb:      vmOsDiskSizeGb,
		ManagedDiskType: "Standard_LRS",
	}

	// Create a virtual machine config
	vmconfig := &shared.VirtualMachineConfig{
		Name:                        vmName,
		Location:                    azureLocation,
		ResourceGroupName:           resourceGroupName,
		VmSize:                      vmSize,
		OsDiskSizeGb:                vmOsDiskSizeGb,
		OsType:                      vmOsType,
		AdminUsername:               vmAdminUsername,
		AdminPassword:               vmAdminPassword,
		Vnet:                        vnc,
		Subnet:                      vnc.Subnet,
		NicId:                       nicId,
		StorageOsDiskConfig:         vmosd,
		StorageImageReferenceConfig: vmir,
	}

	// Create a virtual machine
	shared.CreateAzVirtualMachine(stack, *vmconfig)

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewResourceGroupStack(app, "cdktf-rg-win-linux")
	NewVnetStack(app, "cdktf-vnet-win-linux")
	NewVmStack(app, "cdktf-vm-win-linux")

	app.Synth()
}
