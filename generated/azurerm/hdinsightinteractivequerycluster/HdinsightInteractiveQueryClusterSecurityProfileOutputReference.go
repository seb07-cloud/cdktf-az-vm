package hdinsightinteractivequerycluster

import (
	_init_ "cdk.tf/go/stack/generated/azurerm/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"cdk.tf/go/stack/generated/azurerm/hdinsightinteractivequerycluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HdinsightInteractiveQueryClusterSecurityProfileOutputReference interface {
	cdktf.ComplexObject
	AaddsResourceId() *string
	SetAaddsResourceId(val *string)
	AaddsResourceIdInput() *string
	ClusterUsersGroupDns() *[]*string
	SetClusterUsersGroupDns(val *[]*string)
	ClusterUsersGroupDnsInput() *[]*string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	DomainUsername() *string
	SetDomainUsername(val *string)
	DomainUsernameInput() *string
	DomainUserPassword() *string
	SetDomainUserPassword(val *string)
	DomainUserPasswordInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *HdinsightInteractiveQueryClusterSecurityProfile
	SetInternalValue(val *HdinsightInteractiveQueryClusterSecurityProfile)
	LdapsUrls() *[]*string
	SetLdapsUrls(val *[]*string)
	LdapsUrlsInput() *[]*string
	MsiResourceId() *string
	SetMsiResourceId(val *string)
	MsiResourceIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetClusterUsersGroupDns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightInteractiveQueryClusterSecurityProfileOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) AaddsResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aaddsResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) AaddsResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aaddsResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) ClusterUsersGroupDns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterUsersGroupDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) ClusterUsersGroupDnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterUsersGroupDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) DomainUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) DomainUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) DomainUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) DomainUserPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) InternalValue() *HdinsightInteractiveQueryClusterSecurityProfile {
	var returns *HdinsightInteractiveQueryClusterSecurityProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) LdapsUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ldapsUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) LdapsUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ldapsUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) MsiResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msiResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) MsiResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msiResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterSecurityProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterSecurityProfileOutputReference {
	_init_.Initialize()

	if err := validateNewHdinsightInteractiveQueryClusterSecurityProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference{}

	_jsii_.Create(
		"azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterSecurityProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterSecurityProfileOutputReference_Override(h HdinsightInteractiveQueryClusterSecurityProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterSecurityProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference)SetAaddsResourceId(val *string) {
	if err := j.validateSetAaddsResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aaddsResourceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference)SetClusterUsersGroupDns(val *[]*string) {
	if err := j.validateSetClusterUsersGroupDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterUsersGroupDns",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference)SetDomainUsername(val *string) {
	if err := j.validateSetDomainUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainUsername",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference)SetDomainUserPassword(val *string) {
	if err := j.validateSetDomainUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainUserPassword",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference)SetInternalValue(val *HdinsightInteractiveQueryClusterSecurityProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference)SetLdapsUrls(val *[]*string) {
	if err := j.validateSetLdapsUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ldapsUrls",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference)SetMsiResourceId(val *string) {
	if err := j.validateSetMsiResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"msiResourceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) ResetClusterUsersGroupDns() {
	_jsii_.InvokeVoid(
		h,
		"resetClusterUsersGroupDns",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

