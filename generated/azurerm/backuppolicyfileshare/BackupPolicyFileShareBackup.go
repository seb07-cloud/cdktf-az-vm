package backuppolicyfileshare


type BackupPolicyFileShareBackup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/backup_policy_file_share#frequency BackupPolicyFileShare#frequency}.
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// hourly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/backup_policy_file_share#hourly BackupPolicyFileShare#hourly}
	Hourly *BackupPolicyFileShareBackupHourly `field:"optional" json:"hourly" yaml:"hourly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/backup_policy_file_share#time BackupPolicyFileShare#time}.
	Time *string `field:"optional" json:"time" yaml:"time"`
}

