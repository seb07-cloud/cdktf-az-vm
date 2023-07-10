package expressroutecircuitpeering

import (
	_init_ "cdk.tf/go/stack/generated/azurerm/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"cdk.tf/go/stack/generated/azurerm/expressroutecircuitpeering/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference interface {
	cdktf.ComplexObject
	AdvertisedPublicPrefixes() *[]*string
	SetAdvertisedPublicPrefixes(val *[]*string)
	AdvertisedPublicPrefixesInput() *[]*string
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
	CustomerAsn() *float64
	SetCustomerAsn(val *float64)
	CustomerAsnInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *ExpressRouteCircuitPeeringMicrosoftPeeringConfig
	SetInternalValue(val *ExpressRouteCircuitPeeringMicrosoftPeeringConfig)
	RoutingRegistryName() *string
	SetRoutingRegistryName(val *string)
	RoutingRegistryNameInput() *string
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
	ResetCustomerAsn()
	ResetRoutingRegistryName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference
type jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) AdvertisedPublicPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advertisedPublicPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) AdvertisedPublicPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advertisedPublicPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) CustomerAsn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customerAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) CustomerAsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customerAsnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) InternalValue() *ExpressRouteCircuitPeeringMicrosoftPeeringConfig {
	var returns *ExpressRouteCircuitPeeringMicrosoftPeeringConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) RoutingRegistryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingRegistryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) RoutingRegistryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingRegistryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference {
	_init_.Initialize()

	if err := validateNewExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference{}

	_jsii_.Create(
		"azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference_Override(e ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.expressRouteCircuitPeering.ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference)SetAdvertisedPublicPrefixes(val *[]*string) {
	if err := j.validateSetAdvertisedPublicPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advertisedPublicPrefixes",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference)SetCustomerAsn(val *float64) {
	if err := j.validateSetCustomerAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerAsn",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference)SetInternalValue(val *ExpressRouteCircuitPeeringMicrosoftPeeringConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference)SetRoutingRegistryName(val *string) {
	if err := j.validateSetRoutingRegistryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingRegistryName",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) ResetCustomerAsn() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomerAsn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) ResetRoutingRegistryName() {
	_jsii_.InvokeVoid(
		e,
		"resetRoutingRegistryName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRouteCircuitPeeringMicrosoftPeeringConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

