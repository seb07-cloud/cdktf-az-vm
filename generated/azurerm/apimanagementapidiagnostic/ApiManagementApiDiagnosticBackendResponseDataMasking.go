package apimanagementapidiagnostic


type ApiManagementApiDiagnosticBackendResponseDataMasking struct {
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/api_management_api_diagnostic#headers ApiManagementApiDiagnostic#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// query_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/api_management_api_diagnostic#query_params ApiManagementApiDiagnostic#query_params}
	QueryParams interface{} `field:"optional" json:"queryParams" yaml:"queryParams"`
}

