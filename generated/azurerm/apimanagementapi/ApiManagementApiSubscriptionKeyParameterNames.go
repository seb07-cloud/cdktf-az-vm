package apimanagementapi


type ApiManagementApiSubscriptionKeyParameterNames struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/api_management_api#header ApiManagementApi#header}.
	Header *string `field:"required" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/api_management_api#query ApiManagementApi#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
}

