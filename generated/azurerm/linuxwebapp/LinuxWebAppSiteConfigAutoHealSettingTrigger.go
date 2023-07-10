package linuxwebapp


type LinuxWebAppSiteConfigAutoHealSettingTrigger struct {
	// requests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app#requests LinuxWebApp#requests}
	Requests *LinuxWebAppSiteConfigAutoHealSettingTriggerRequests `field:"optional" json:"requests" yaml:"requests"`
	// slow_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app#slow_request LinuxWebApp#slow_request}
	SlowRequest interface{} `field:"optional" json:"slowRequest" yaml:"slowRequest"`
	// status_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_web_app#status_code LinuxWebApp#status_code}
	StatusCode interface{} `field:"optional" json:"statusCode" yaml:"statusCode"`
}

