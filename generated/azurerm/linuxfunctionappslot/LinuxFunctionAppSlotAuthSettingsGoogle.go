package linuxfunctionappslot


type LinuxFunctionAppSlotAuthSettingsGoogle struct {
	// The OpenID Connect Client ID for the Google web application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_function_app_slot#client_id LinuxFunctionAppSlot#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret associated with the Google web application.  Cannot be specified with `client_secret_setting_name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_function_app_slot#client_secret LinuxFunctionAppSlot#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The app setting name that contains the `client_secret` value used for Google Login. Cannot be specified with `client_secret`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_function_app_slot#client_secret_setting_name LinuxFunctionAppSlot#client_secret_setting_name}
	ClientSecretSettingName *string `field:"optional" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
	// Specifies a list of OAuth 2.0 scopes that will be requested as part of Google Sign-In authentication. If not specified, "openid", "profile", and "email" are used as default scopes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_function_app_slot#oauth_scopes LinuxFunctionAppSlot#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

