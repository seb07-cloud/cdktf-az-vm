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
	vnetAddressSpace  = "10.0.0.0/16"
	snetAddressSpace  = "10.0.0.0/24"

	// Azure Virtual Machine based consts
	vmName          = "cdktf-vm-win-linux"
	vmSize          = "Standard_B1s"
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

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// Create an Azure provider
	shared.NewAzProvider(stack, "azurerm")

	// Create a resource group
	rgrp := shared.NewAzResourceGroupConfig(resourceGroupName, azureLocation)
	shared.CreateAzResourceGroup(stack, *rgrp)

	// Create a virtual network
	vnet := shared.NewAzVnetConfig(vnetName, vnetAddressSpace, azureLocation, rgrp.Name, snetAddressSpace)
	shared.CreateAzVirtualNetwork(stack, *vnet)

	// Create a vm with a public ip
	nicconfig := shared.NewAzNetworkInterfaceConfig(vmName, azureLocation, rgrp.Name, vnet, vnet.Subnet)
	nicId := shared.CreateAzNetworkInterface(stack, *nicconfig)

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
		ResourceGroupName:           rgrp.Name,
		VmSize:                      vmSize,
		OsDiskSizeGb:                vmOsDiskSizeGb,
		OsType:                      vmOsType,
		AdminUsername:               vmAdminUsername,
		AdminPassword:               vmAdminPassword,
		Vnet:                        vnet,
		Subnet:                      vnet.Subnet,
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

	NewMyStack(app, "cdktf-vm-win-linux")

	app.Synth()
}
