package windowswebappslot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsWebAppSlotConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#app_service_id WindowsWebAppSlot#app_service_id}.
	AppServiceId *string `field:"required" json:"appServiceId" yaml:"appServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#name WindowsWebAppSlot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// site_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#site_config WindowsWebAppSlot#site_config}
	SiteConfig *WindowsWebAppSlotSiteConfig `field:"required" json:"siteConfig" yaml:"siteConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#app_settings WindowsWebAppSlot#app_settings}.
	AppSettings *map[string]*string `field:"optional" json:"appSettings" yaml:"appSettings"`
	// auth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#auth_settings WindowsWebAppSlot#auth_settings}
	AuthSettings *WindowsWebAppSlotAuthSettings `field:"optional" json:"authSettings" yaml:"authSettings"`
	// backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#backup WindowsWebAppSlot#backup}
	Backup *WindowsWebAppSlotBackup `field:"optional" json:"backup" yaml:"backup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#client_affinity_enabled WindowsWebAppSlot#client_affinity_enabled}.
	ClientAffinityEnabled interface{} `field:"optional" json:"clientAffinityEnabled" yaml:"clientAffinityEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#client_certificate_enabled WindowsWebAppSlot#client_certificate_enabled}.
	ClientCertificateEnabled interface{} `field:"optional" json:"clientCertificateEnabled" yaml:"clientCertificateEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#client_certificate_mode WindowsWebAppSlot#client_certificate_mode}.
	ClientCertificateMode *string `field:"optional" json:"clientCertificateMode" yaml:"clientCertificateMode"`
	// connection_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#connection_string WindowsWebAppSlot#connection_string}
	ConnectionString interface{} `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#enabled WindowsWebAppSlot#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#https_only WindowsWebAppSlot#https_only}.
	HttpsOnly interface{} `field:"optional" json:"httpsOnly" yaml:"httpsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#id WindowsWebAppSlot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#identity WindowsWebAppSlot#identity}
	Identity *WindowsWebAppSlotIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#key_vault_reference_identity_id WindowsWebAppSlot#key_vault_reference_identity_id}.
	KeyVaultReferenceIdentityId *string `field:"optional" json:"keyVaultReferenceIdentityId" yaml:"keyVaultReferenceIdentityId"`
	// logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#logs WindowsWebAppSlot#logs}
	Logs *WindowsWebAppSlotLogs `field:"optional" json:"logs" yaml:"logs"`
	// storage_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#storage_account WindowsWebAppSlot#storage_account}
	StorageAccount interface{} `field:"optional" json:"storageAccount" yaml:"storageAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#tags WindowsWebAppSlot#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#timeouts WindowsWebAppSlot#timeouts}
	Timeouts *WindowsWebAppSlotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
