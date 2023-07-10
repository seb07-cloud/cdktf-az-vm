package batchpool


type BatchPoolFixedScale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/batch_pool#resize_timeout BatchPool#resize_timeout}.
	ResizeTimeout *string `field:"optional" json:"resizeTimeout" yaml:"resizeTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/batch_pool#target_dedicated_nodes BatchPool#target_dedicated_nodes}.
	TargetDedicatedNodes *float64 `field:"optional" json:"targetDedicatedNodes" yaml:"targetDedicatedNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/batch_pool#target_low_priority_nodes BatchPool#target_low_priority_nodes}.
	TargetLowPriorityNodes *float64 `field:"optional" json:"targetLowPriorityNodes" yaml:"targetLowPriorityNodes"`
}
