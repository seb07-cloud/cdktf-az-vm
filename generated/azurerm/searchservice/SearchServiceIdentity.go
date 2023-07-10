package searchservice


type SearchServiceIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/search_service#type SearchService#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

