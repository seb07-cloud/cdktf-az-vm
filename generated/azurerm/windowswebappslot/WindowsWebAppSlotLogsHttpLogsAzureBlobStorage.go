package windowswebappslot


type WindowsWebAppSlotLogsHttpLogsAzureBlobStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#sas_url WindowsWebAppSlot#sas_url}.
	SasUrl *string `field:"required" json:"sasUrl" yaml:"sasUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#retention_in_days WindowsWebAppSlot#retention_in_days}.
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
}

