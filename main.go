package main

import (
	"cdk.tf/go/stack/shared"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

const (
	// Azure location to deploy the resources to
	azureLocation    = "westeurope"
	vnetAddressSpace = "10.0.0.0/16"
	snetAddressSpace = "10.0.0.0/24"

	// Azure Virtual Machine based consts
	vmName          = "cdktf-vm-win-linux"
	vmSize          = "Standard_B1s"
	vmOsDiskSizeGb  = 30
	vmOsType        = "Linux"
	vmAdminUsername = "cdktfuser"
	vmAdminPassword = "cdktfpassword"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// Create an Azure provider
	shared.NewAzProvider(stack, "azurerm")

	// Create a resource group
	rgrp := shared.NewAzResourceGroupConfig("cdktf-vm-win-linux", azureLocation)
	shared.CreateAzResourceGroup(stack, *rgrp)

	// Create a virtual network
	vnet := shared.NewAzVnetConfig("cdktf-vm-win-linux", vnetAddressSpace, azureLocation, rgrp.Name, snetAddressSpace)
	shared.CreateAzVirtualNetwork(stack, *vnet)

	// Create a vm with a public ip
	vm := shared.NewAzVirtualMachineConfig(vmName, azureLocation, rgrp.Name, vmSize, 30, vmOsType, vmAdminUsername, vmAdminPassword, vnet, vnet.Subnet)
	shared.CreateAzVirtualMachine(stack, *vm)

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdktf-vm-win-linux")

	app.Synth()
}
