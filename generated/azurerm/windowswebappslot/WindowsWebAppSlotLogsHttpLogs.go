package windowswebappslot


type WindowsWebAppSlotLogsHttpLogs struct {
	// azure_blob_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#azure_blob_storage WindowsWebAppSlot#azure_blob_storage}
	AzureBlobStorage *WindowsWebAppSlotLogsHttpLogsAzureBlobStorage `field:"optional" json:"azureBlobStorage" yaml:"azureBlobStorage"`
	// file_system block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#file_system WindowsWebAppSlot#file_system}
	FileSystem *WindowsWebAppSlotLogsHttpLogsFileSystem `field:"optional" json:"fileSystem" yaml:"fileSystem"`
}

