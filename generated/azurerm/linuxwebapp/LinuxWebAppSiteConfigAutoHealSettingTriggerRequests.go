package linuxwebapp


type LinuxWebAppSiteConfigAutoHealSettingTriggerRequests struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_web_app#count LinuxWebApp#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_web_app#interval LinuxWebApp#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
}

