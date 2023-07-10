package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	provider "github.com/seb07-cloud/cdktf-vm-win-linux/generated/azurerm/provider"
	resourceGroup "github.com/seb07-cloud/cdktf-vm-win-linux/generated/azurerm/resource_group"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// New Azurerm Provider instance
	azurermProvider := provider.NewAzurermProvider(stack, "azurerm", &provider.AzurermProviderConfig{
		Features: &[]*provider.AzurermProviderFeatures{},
	})

	// New Resource Group instance
	resGrp := resourceGroup.NewResourceGroup(stack, "resourceGroup", &provider.ResourceGroupConfig{
		Name:     &"cdktf-vm-win-linux",
		Location: &"westeurope",
	})
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdktf-vm-win-linux")

	app.Synth()
}
