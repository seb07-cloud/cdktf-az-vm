package vpnserverconfiguration


type VpnServerConfigurationClientRevokedCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/vpn_server_configuration#name VpnServerConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/vpn_server_configuration#thumbprint VpnServerConfiguration#thumbprint}.
	Thumbprint *string `field:"required" json:"thumbprint" yaml:"thumbprint"`
}
