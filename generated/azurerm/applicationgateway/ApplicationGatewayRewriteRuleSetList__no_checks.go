//go:build no_runtime_type_checking

package applicationgateway

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApplicationGatewayRewriteRuleSetList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApplicationGatewayRewriteRuleSetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApplicationGatewayRewriteRuleSetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

