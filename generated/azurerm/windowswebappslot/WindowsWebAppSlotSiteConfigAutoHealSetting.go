package windowswebappslot


type WindowsWebAppSlotSiteConfigAutoHealSetting struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/windows_web_app_slot#action WindowsWebAppSlot#action}
	Action *WindowsWebAppSlotSiteConfigAutoHealSettingAction `field:"required" json:"action" yaml:"action"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/windows_web_app_slot#trigger WindowsWebAppSlot#trigger}
	Trigger *WindowsWebAppSlotSiteConfigAutoHealSettingTrigger `field:"required" json:"trigger" yaml:"trigger"`
}

