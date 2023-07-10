package linuxfunctionappslot


type LinuxFunctionAppSlotAuthSettingsFacebook struct {
	// The App ID of the Facebook app used for login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_function_app_slot#app_id LinuxFunctionAppSlot#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The App Secret of the Facebook app used for Facebook Login. Cannot be specified with `app_secret_setting_name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_function_app_slot#app_secret LinuxFunctionAppSlot#app_secret}
	AppSecret *string `field:"optional" json:"appSecret" yaml:"appSecret"`
	// The app setting name that contains the `app_secret` value used for Facebook Login. Cannot be specified with `app_secret`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_function_app_slot#app_secret_setting_name LinuxFunctionAppSlot#app_secret_setting_name}
	AppSecretSettingName *string `field:"optional" json:"appSecretSettingName" yaml:"appSecretSettingName"`
	// Specifies a list of OAuth 2.0 scopes to be requested as part of Facebook Login authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_function_app_slot#oauth_scopes LinuxFunctionAppSlot#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

