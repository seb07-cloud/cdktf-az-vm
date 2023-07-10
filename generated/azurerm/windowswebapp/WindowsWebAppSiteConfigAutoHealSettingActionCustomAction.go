package windowswebapp


type WindowsWebAppSiteConfigAutoHealSettingActionCustomAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/windows_web_app#executable WindowsWebApp#executable}.
	Executable *string `field:"required" json:"executable" yaml:"executable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/windows_web_app#parameters WindowsWebApp#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

