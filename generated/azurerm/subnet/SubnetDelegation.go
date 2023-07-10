package subnet


type SubnetDelegation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/subnet#name Subnet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// service_delegation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/subnet#service_delegation Subnet#service_delegation}
	ServiceDelegation *SubnetDelegationServiceDelegation `field:"required" json:"serviceDelegation" yaml:"serviceDelegation"`
}
