package eventgrideventsubscription


type EventgridEventSubscriptionAdvancedFilterNumberIn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}
