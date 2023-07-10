package shared

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	azurermprovider "cdk.tf/go/stack/generated/azurerm/provider"
)

func NewAzProvider(scope constructs.Construct, id string) cdktf.TerraformProvider {
	return azurermprovider.NewAzurermProvider(scope, &id,
		&azurermprovider.AzurermProviderConfig{
			Features: &azurermprovider.AzurermProviderFeatures{},
		})
}
