package linuxfunctionappslot


type LinuxFunctionAppSlotConnectionString struct {
	// The name which should be used for this Connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_function_app_slot#name LinuxFunctionAppSlot#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of database. Possible values include: `MySQL`, `SQLServer`, `SQLAzure`, `Custom`, `NotificationHub`, `ServiceBus`, `EventHub`, `APIHub`, `DocDb`, `RedisCache`, and `PostgreSQL`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_function_app_slot#type LinuxFunctionAppSlot#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The connection string value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/linux_function_app_slot#value LinuxFunctionAppSlot#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

