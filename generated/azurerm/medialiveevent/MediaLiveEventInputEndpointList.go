package medialiveevent

import (
	_init_ "cdk.tf/go/stack/generated/azurerm/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"cdk.tf/go/stack/generated/azurerm/medialiveevent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaLiveEventInputEndpointList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) MediaLiveEventInputEndpointOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaLiveEventInputEndpointList
type jsiiProxy_MediaLiveEventInputEndpointList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MediaLiveEventInputEndpointList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputEndpointList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputEndpointList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputEndpointList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaLiveEventInputEndpointList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMediaLiveEventInputEndpointList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MediaLiveEventInputEndpointList {
	_init_.Initialize()

	if err := validateNewMediaLiveEventInputEndpointListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaLiveEventInputEndpointList{}

	_jsii_.Create(
		"azurerm.mediaLiveEvent.MediaLiveEventInputEndpointList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMediaLiveEventInputEndpointList_Override(m MediaLiveEventInputEndpointList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.mediaLiveEvent.MediaLiveEventInputEndpointList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MediaLiveEventInputEndpointList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventInputEndpointList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaLiveEventInputEndpointList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MediaLiveEventInputEndpointList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventInputEndpointList) Get(index *float64) MediaLiveEventInputEndpointOutputReference {
	if err := m.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns MediaLiveEventInputEndpointOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventInputEndpointList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaLiveEventInputEndpointList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

