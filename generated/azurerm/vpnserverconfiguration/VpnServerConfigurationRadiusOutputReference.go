package vpnserverconfiguration

import (
	_init_ "cdk.tf/go/stack/generated/azurerm/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"cdk.tf/go/stack/generated/azurerm/vpnserverconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnServerConfigurationRadiusOutputReference interface {
	cdktf.ComplexObject
	ClientRootCertificate() VpnServerConfigurationRadiusClientRootCertificateList
	ClientRootCertificateInput() interface{}
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
	// Experimental.
	Fqn() *string
	InternalValue() *VpnServerConfigurationRadius
	SetInternalValue(val *VpnServerConfigurationRadius)
	Server() VpnServerConfigurationRadiusServerList
	ServerInput() interface{}
	ServerRootCertificate() VpnServerConfigurationRadiusServerRootCertificateList
	ServerRootCertificateInput() interface{}
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
	PutClientRootCertificate(value interface{})
	PutServer(value interface{})
	PutServerRootCertificate(value interface{})
	ResetClientRootCertificate()
	ResetServer()
	ResetServerRootCertificate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnServerConfigurationRadiusOutputReference
type jsiiProxy_VpnServerConfigurationRadiusOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference) ClientRootCertificate() VpnServerConfigurationRadiusClientRootCertificateList {
	var returns VpnServerConfigurationRadiusClientRootCertificateList
	_jsii_.Get(
		j,
		"clientRootCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference) ClientRootCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientRootCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference) InternalValue() *VpnServerConfigurationRadius {
	var returns *VpnServerConfigurationRadius
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference) Server() VpnServerConfigurationRadiusServerList {
	var returns VpnServerConfigurationRadiusServerList
	_jsii_.Get(
		j,
		"server",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference) ServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference) ServerRootCertificate() VpnServerConfigurationRadiusServerRootCertificateList {
	var returns VpnServerConfigurationRadiusServerRootCertificateList
	_jsii_.Get(
		j,
		"serverRootCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference) ServerRootCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverRootCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnServerConfigurationRadiusOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpnServerConfigurationRadiusOutputReference {
	_init_.Initialize()

	if err := validateNewVpnServerConfigurationRadiusOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnServerConfigurationRadiusOutputReference{}

	_jsii_.Create(
		"azurerm.vpnServerConfiguration.VpnServerConfigurationRadiusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpnServerConfigurationRadiusOutputReference_Override(v VpnServerConfigurationRadiusOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.vpnServerConfiguration.VpnServerConfigurationRadiusOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference)SetInternalValue(val *VpnServerConfigurationRadius) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationRadiusOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) PutClientRootCertificate(value interface{}) {
	if err := v.validatePutClientRootCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putClientRootCertificate",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) PutServer(value interface{}) {
	if err := v.validatePutServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putServer",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) PutServerRootCertificate(value interface{}) {
	if err := v.validatePutServerRootCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putServerRootCertificate",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) ResetClientRootCertificate() {
	_jsii_.InvokeVoid(
		v,
		"resetClientRootCertificate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) ResetServer() {
	_jsii_.InvokeVoid(
		v,
		"resetServer",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) ResetServerRootCertificate() {
	_jsii_.InvokeVoid(
		v,
		"resetServerRootCertificate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationRadiusOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

