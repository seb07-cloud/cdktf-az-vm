package cognitiveaccount


type CognitiveAccountNetworkAclsVirtualNetworkRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/cognitive_account#subnet_id CognitiveAccount#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/cognitive_account#ignore_missing_vnet_service_endpoint CognitiveAccount#ignore_missing_vnet_service_endpoint}.
	IgnoreMissingVnetServiceEndpoint interface{} `field:"optional" json:"ignoreMissingVnetServiceEndpoint" yaml:"ignoreMissingVnetServiceEndpoint"`
}

