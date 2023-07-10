package resourcegrouppolicyassignment


type ResourceGroupPolicyAssignmentIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/resource_group_policy_assignment#type ResourceGroupPolicyAssignment#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/resource_group_policy_assignment#identity_ids ResourceGroupPolicyAssignment#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

