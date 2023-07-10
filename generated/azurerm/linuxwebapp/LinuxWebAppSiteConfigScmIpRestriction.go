package linuxwebapp


type LinuxWebAppSiteConfigScmIpRestriction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app#action LinuxWebApp#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app#headers LinuxWebApp#headers}.
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app#ip_address LinuxWebApp#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app#name LinuxWebApp#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app#priority LinuxWebApp#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app#service_tag LinuxWebApp#service_tag}.
	ServiceTag *string `field:"optional" json:"serviceTag" yaml:"serviceTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app#virtual_network_subnet_id LinuxWebApp#virtual_network_subnet_id}.
	VirtualNetworkSubnetId *string `field:"optional" json:"virtualNetworkSubnetId" yaml:"virtualNetworkSubnetId"`
}

