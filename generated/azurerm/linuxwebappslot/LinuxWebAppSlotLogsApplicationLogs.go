package linuxwebappslot


type LinuxWebAppSlotLogsApplicationLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_web_app_slot#file_system_level LinuxWebAppSlot#file_system_level}.
	FileSystemLevel *string `field:"required" json:"fileSystemLevel" yaml:"fileSystemLevel"`
	// azure_blob_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_web_app_slot#azure_blob_storage LinuxWebAppSlot#azure_blob_storage}
	AzureBlobStorage *LinuxWebAppSlotLogsApplicationLogsAzureBlobStorage `field:"optional" json:"azureBlobStorage" yaml:"azureBlobStorage"`
}

