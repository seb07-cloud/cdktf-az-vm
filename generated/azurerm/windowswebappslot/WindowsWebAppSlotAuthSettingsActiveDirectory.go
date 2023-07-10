package windowswebappslot


type WindowsWebAppSlotAuthSettingsActiveDirectory struct {
	// The ID of the Client to use to authenticate with Azure Active Directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#client_id WindowsWebAppSlot#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Specifies a list of Allowed audience values to consider when validating JWTs issued by Azure Active Directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#allowed_audiences WindowsWebAppSlot#allowed_audiences}
	AllowedAudiences *[]*string `field:"optional" json:"allowedAudiences" yaml:"allowedAudiences"`
	// The Client Secret for the Client ID. Cannot be used with `client_secret_setting_name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#client_secret WindowsWebAppSlot#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The App Setting name that contains the client secret of the Client. Cannot be used with `client_secret`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/windows_web_app_slot#client_secret_setting_name WindowsWebAppSlot#client_secret_setting_name}
	ClientSecretSettingName *string `field:"optional" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
}
