package hdinsighthbasecluster

import (
	_init_ "cdk.tf/go/stack/generated/azurerm/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"cdk.tf/go/stack/generated/azurerm/hdinsighthbasecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HdinsightHbaseClusterStorageAccountGen2OutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FilesystemId() *string
	SetFilesystemId(val *string)
	FilesystemIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *HdinsightHbaseClusterStorageAccountGen2
	SetInternalValue(val *HdinsightHbaseClusterStorageAccountGen2)
	IsDefault() interface{}
	SetIsDefault(val interface{})
	IsDefaultInput() interface{}
	ManagedIdentityResourceId() *string
	SetManagedIdentityResourceId(val *string)
	ManagedIdentityResourceIdInput() *string
	StorageResourceId() *string
	SetStorageResourceId(val *string)
	StorageResourceIdInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHbaseClusterStorageAccountGen2OutputReference
type jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) FilesystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filesystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) FilesystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filesystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) InternalValue() *HdinsightHbaseClusterStorageAccountGen2 {
	var returns *HdinsightHbaseClusterStorageAccountGen2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) IsDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) IsDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) ManagedIdentityResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) ManagedIdentityResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) StorageResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) StorageResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHbaseClusterStorageAccountGen2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHbaseClusterStorageAccountGen2OutputReference {
	_init_.Initialize()

	if err := validateNewHdinsightHbaseClusterStorageAccountGen2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference{}

	_jsii_.Create(
		"azurerm.hdinsightHbaseCluster.HdinsightHbaseClusterStorageAccountGen2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHbaseClusterStorageAccountGen2OutputReference_Override(h HdinsightHbaseClusterStorageAccountGen2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.hdinsightHbaseCluster.HdinsightHbaseClusterStorageAccountGen2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference)SetFilesystemId(val *string) {
	if err := j.validateSetFilesystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filesystemId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference)SetInternalValue(val *HdinsightHbaseClusterStorageAccountGen2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference)SetIsDefault(val interface{}) {
	if err := j.validateSetIsDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDefault",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference)SetManagedIdentityResourceId(val *string) {
	if err := j.validateSetManagedIdentityResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedIdentityResourceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference)SetStorageResourceId(val *string) {
	if err := j.validateSetStorageResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageResourceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (h *jsiiProxy_HdinsightHbaseClusterStorageAccountGen2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

