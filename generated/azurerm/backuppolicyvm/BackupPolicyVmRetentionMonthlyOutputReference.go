package backuppolicyvm

import (
	_init_ "cdk.tf/go/stack/generated/azurerm/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"cdk.tf/go/stack/generated/azurerm/backuppolicyvm/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupPolicyVmRetentionMonthlyOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Days() *[]*float64
	SetDays(val *[]*float64)
	DaysInput() *[]*float64
	// Experimental.
	Fqn() *string
	IncludeLastDays() interface{}
	SetIncludeLastDays(val interface{})
	IncludeLastDaysInput() interface{}
	InternalValue() *BackupPolicyVmRetentionMonthly
	SetInternalValue(val *BackupPolicyVmRetentionMonthly)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Weekdays() *[]*string
	SetWeekdays(val *[]*string)
	WeekdaysInput() *[]*string
	Weeks() *[]*string
	SetWeeks(val *[]*string)
	WeeksInput() *[]*string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDays()
	ResetIncludeLastDays()
	ResetWeekdays()
	ResetWeeks()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupPolicyVmRetentionMonthlyOutputReference
type jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) Days() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"days",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) DaysInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"daysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) IncludeLastDays() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeLastDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) IncludeLastDaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeLastDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) InternalValue() *BackupPolicyVmRetentionMonthly {
	var returns *BackupPolicyVmRetentionMonthly
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) Weekdays() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekdays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) WeekdaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weekdaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) Weeks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) WeeksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeksInput",
		&returns,
	)
	return returns
}


func NewBackupPolicyVmRetentionMonthlyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupPolicyVmRetentionMonthlyOutputReference {
	_init_.Initialize()

	if err := validateNewBackupPolicyVmRetentionMonthlyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference{}

	_jsii_.Create(
		"azurerm.backupPolicyVm.BackupPolicyVmRetentionMonthlyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupPolicyVmRetentionMonthlyOutputReference_Override(b BackupPolicyVmRetentionMonthlyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.backupPolicyVm.BackupPolicyVmRetentionMonthlyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference)SetCount(val *float64) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference)SetDays(val *[]*float64) {
	if err := j.validateSetDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"days",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference)SetIncludeLastDays(val interface{}) {
	if err := j.validateSetIncludeLastDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeLastDays",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference)SetInternalValue(val *BackupPolicyVmRetentionMonthly) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference)SetWeekdays(val *[]*string) {
	if err := j.validateSetWeekdaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekdays",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference)SetWeeks(val *[]*string) {
	if err := j.validateSetWeeksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeks",
		val,
	)
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) ResetDays() {
	_jsii_.InvokeVoid(
		b,
		"resetDays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) ResetIncludeLastDays() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeLastDays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) ResetWeekdays() {
	_jsii_.InvokeVoid(
		b,
		"resetWeekdays",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) ResetWeeks() {
	_jsii_.InvokeVoid(
		b,
		"resetWeeks",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmRetentionMonthlyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

