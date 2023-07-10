package main

import (
	"cdk.tf/go/stack/shared"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

const (
	// Azure location to deploy the resources to
	azureLocation = "westeurope"
)

var (
	// Address space for the virtual network
	vnetAddressSpace = "10.0.0.0/16"
	// Address space for the subnet
	snetAddressSpace = "10.0.0.0/24"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// Create an Azure provider
	shared.NewAzProvider(stack, "azurerm")

	// Create a resource group
	rgrp := shared.NewAzResourceGroupConfig("cdktf-vm-win-linux", azureLocation)
	shared.CreateAzResourceGroup(stack, rgrp)

	// Create a virtual network
	vnet := shared.NewAzVnetConfig("cdktf-vm-win-linux", vnetAddressSpace, azureLocation, rgrp.Name, snetAddressSpace)
	shared.CreateAzVirtualNetwork(stack, *vnet)
	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdktf-vm-win-linux")

	app.Synth()
}
