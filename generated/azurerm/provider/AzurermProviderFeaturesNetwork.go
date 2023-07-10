package provider


type AzurermProviderFeaturesNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs#relaxed_locking AzurermProvider#relaxed_locking}.
	RelaxedLocking interface{} `field:"required" json:"relaxedLocking" yaml:"relaxedLocking"`
}
