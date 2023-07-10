package windowswebappslot


type WindowsWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/windows_web_app_slot#count WindowsWebAppSlot#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/windows_web_app_slot#interval WindowsWebAppSlot#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/windows_web_app_slot#time_taken WindowsWebAppSlot#time_taken}.
	TimeTaken *string `field:"required" json:"timeTaken" yaml:"timeTaken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/windows_web_app_slot#path WindowsWebAppSlot#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

