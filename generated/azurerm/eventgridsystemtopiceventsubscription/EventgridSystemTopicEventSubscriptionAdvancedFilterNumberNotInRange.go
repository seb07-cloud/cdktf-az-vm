package eventgridsystemtopiceventsubscription


type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values interface{} `field:"required" json:"values" yaml:"values"`
}
