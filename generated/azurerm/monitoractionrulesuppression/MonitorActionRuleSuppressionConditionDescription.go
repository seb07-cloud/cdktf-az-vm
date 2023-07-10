package monitoractionrulesuppression


type MonitorActionRuleSuppressionConditionDescription struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/monitor_action_rule_suppression#operator MonitorActionRuleSuppression#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/monitor_action_rule_suppression#values MonitorActionRuleSuppression#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

