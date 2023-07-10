package resourcepolicyremediation


type ResourcePolicyRemediationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/resource_policy_remediation#create ResourcePolicyRemediation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/resource_policy_remediation#delete ResourcePolicyRemediation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/resource_policy_remediation#read ResourcePolicyRemediation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/resource_policy_remediation#update ResourcePolicyRemediation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

