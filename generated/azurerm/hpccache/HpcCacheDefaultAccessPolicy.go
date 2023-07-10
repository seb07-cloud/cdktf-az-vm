package hpccache


type HpcCacheDefaultAccessPolicy struct {
	// access_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/hpc_cache#access_rule HpcCache#access_rule}
	AccessRule interface{} `field:"required" json:"accessRule" yaml:"accessRule"`
}

