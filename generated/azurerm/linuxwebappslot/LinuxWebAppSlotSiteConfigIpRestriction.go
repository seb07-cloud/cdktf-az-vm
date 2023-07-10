package linuxwebappslot


type LinuxWebAppSlotSiteConfigIpRestriction struct {
	// The action to take. Possible values are `Allow` or `Deny`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_web_app_slot#action LinuxWebAppSlot#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_web_app_slot#headers LinuxWebAppSlot#headers}.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// The CIDR notation of the IP or IP Range to match.
	//
	// For example: `10.0.0.0/24` or `192.168.10.1/32` or `fe80::/64` or `13.107.6.152/31,13.107.128.0/22`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_web_app_slot#ip_address LinuxWebAppSlot#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// The name which should be used for this `ip_restriction`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_web_app_slot#name LinuxWebAppSlot#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The priority value of this `ip_restriction`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_web_app_slot#priority LinuxWebAppSlot#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The Service Tag used for this IP Restriction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_web_app_slot#service_tag LinuxWebAppSlot#service_tag}
	ServiceTag *string `field:"optional" json:"serviceTag" yaml:"serviceTag"`
	// The Virtual Network Subnet ID used for this IP Restriction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_web_app_slot#virtual_network_subnet_id LinuxWebAppSlot#virtual_network_subnet_id}
	VirtualNetworkSubnetId *string `field:"optional" json:"virtualNetworkSubnetId" yaml:"virtualNetworkSubnetId"`
}

