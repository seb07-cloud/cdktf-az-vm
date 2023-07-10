package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	// Import the local Azure provider

	azurermprovider "./generated/azurerm/provider"
	resourcegroup "./generated/azurerm/resourcegroup"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// Create a new instance of the Azure Provider
	azurermprovider.New(stack, "azure", &azurermprovider.AzurermProviderConfig{
		Features: &azurermprovider.AzurermProviderConfig{},
	})

	resourcegroup.New(stack, "resourcegroup", &resourcegroup.ResourcegroupConfig{
		Name:     jsii.String("cdktf-vm-win-linux"),
		Location: jsii.String("westeurope"),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdktf-vm-win-linux")

	app.Synth()
}
