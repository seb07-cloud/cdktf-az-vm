package linuxwebapp


type LinuxWebAppLogsHttpLogsAzureBlobStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app#sas_url LinuxWebApp#sas_url}.
	SasUrl *string `field:"required" json:"sasUrl" yaml:"sasUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app#retention_in_days LinuxWebApp#retention_in_days}.
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
}

