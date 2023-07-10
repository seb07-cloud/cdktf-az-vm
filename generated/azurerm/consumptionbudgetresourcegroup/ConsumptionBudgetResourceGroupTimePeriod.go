package consumptionbudgetresourcegroup


type ConsumptionBudgetResourceGroupTimePeriod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/consumption_budget_resource_group#start_date ConsumptionBudgetResourceGroup#start_date}.
	StartDate *string `field:"required" json:"startDate" yaml:"startDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/consumption_budget_resource_group#end_date ConsumptionBudgetResourceGroup#end_date}.
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
}

