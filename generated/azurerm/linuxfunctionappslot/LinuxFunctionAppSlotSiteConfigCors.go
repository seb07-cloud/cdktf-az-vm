package linuxfunctionappslot


type LinuxFunctionAppSlotSiteConfigCors struct {
	// Specifies a list of origins that should be allowed to make cross-origin calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_function_app_slot#allowed_origins LinuxFunctionAppSlot#allowed_origins}
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Are credentials allowed in CORS requests? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_function_app_slot#support_credentials LinuxFunctionAppSlot#support_credentials}
	SupportCredentials interface{} `field:"optional" json:"supportCredentials" yaml:"supportCredentials"`
}

