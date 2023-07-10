package windowswebapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsWebAppConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#location WindowsWebApp#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#name WindowsWebApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#resource_group_name WindowsWebApp#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#service_plan_id WindowsWebApp#service_plan_id}.
	ServicePlanId *string `field:"required" json:"servicePlanId" yaml:"servicePlanId"`
	// site_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#site_config WindowsWebApp#site_config}
	SiteConfig *WindowsWebAppSiteConfig `field:"required" json:"siteConfig" yaml:"siteConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#app_settings WindowsWebApp#app_settings}.
	AppSettings *map[string]*string `field:"optional" json:"appSettings" yaml:"appSettings"`
	// auth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#auth_settings WindowsWebApp#auth_settings}
	AuthSettings *WindowsWebAppAuthSettings `field:"optional" json:"authSettings" yaml:"authSettings"`
	// backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#backup WindowsWebApp#backup}
	Backup *WindowsWebAppBackup `field:"optional" json:"backup" yaml:"backup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#client_affinity_enabled WindowsWebApp#client_affinity_enabled}.
	ClientAffinityEnabled interface{} `field:"optional" json:"clientAffinityEnabled" yaml:"clientAffinityEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#client_certificate_enabled WindowsWebApp#client_certificate_enabled}.
	ClientCertificateEnabled interface{} `field:"optional" json:"clientCertificateEnabled" yaml:"clientCertificateEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#client_certificate_mode WindowsWebApp#client_certificate_mode}.
	ClientCertificateMode *string `field:"optional" json:"clientCertificateMode" yaml:"clientCertificateMode"`
	// connection_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#connection_string WindowsWebApp#connection_string}
	ConnectionString interface{} `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#enabled WindowsWebApp#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#https_only WindowsWebApp#https_only}.
	HttpsOnly interface{} `field:"optional" json:"httpsOnly" yaml:"httpsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#id WindowsWebApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#identity WindowsWebApp#identity}
	Identity *WindowsWebAppIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#key_vault_reference_identity_id WindowsWebApp#key_vault_reference_identity_id}.
	KeyVaultReferenceIdentityId *string `field:"optional" json:"keyVaultReferenceIdentityId" yaml:"keyVaultReferenceIdentityId"`
	// logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#logs WindowsWebApp#logs}
	Logs *WindowsWebAppLogs `field:"optional" json:"logs" yaml:"logs"`
	// storage_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#storage_account WindowsWebApp#storage_account}
	StorageAccount interface{} `field:"optional" json:"storageAccount" yaml:"storageAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#tags WindowsWebApp#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app#timeouts WindowsWebApp#timeouts}
	Timeouts *WindowsWebAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
