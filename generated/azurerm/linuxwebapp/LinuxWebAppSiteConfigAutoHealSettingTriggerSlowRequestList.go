package linuxwebapp

import (
	_init_ "cdk.tf/go/stack/generated/azurerm/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"cdk.tf/go/stack/generated/azurerm/linuxwebapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList
type jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewLinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList {
	_init_.Initialize()

	if err := validateNewLinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList{}

	_jsii_.Create(
		"azurerm.linuxWebApp.LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewLinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList_Override(l LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.linuxWebApp.LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		l,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (l *jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList) Get(index *float64) LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference {
	if err := l.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestOutputReference

	_jsii_.Invoke(
		l,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxWebAppSiteConfigAutoHealSettingTriggerSlowRequestList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

