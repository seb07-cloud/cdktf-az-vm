package linuxwebappslot


type LinuxWebAppSlotSiteConfigAutoHealSettingTriggerStatusCode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app_slot#count LinuxWebAppSlot#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app_slot#interval LinuxWebAppSlot#interval}.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app_slot#status_code_range LinuxWebAppSlot#status_code_range}.
	StatusCodeRange *string `field:"required" json:"statusCodeRange" yaml:"statusCodeRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app_slot#path LinuxWebAppSlot#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app_slot#sub_status LinuxWebAppSlot#sub_status}.
	SubStatus *float64 `field:"optional" json:"subStatus" yaml:"subStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app_slot#win32_status LinuxWebAppSlot#win32_status}.
	Win32Status *string `field:"optional" json:"win32Status" yaml:"win32Status"`
}

