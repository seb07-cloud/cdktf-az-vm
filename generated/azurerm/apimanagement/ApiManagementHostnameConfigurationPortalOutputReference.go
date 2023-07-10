package apimanagement

import (
	_init_ "cdk.tf/go/stack/generated/azurerm/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"cdk.tf/go/stack/generated/azurerm/apimanagement/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementHostnameConfigurationPortalOutputReference interface {
	cdktf.ComplexObject
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	CertificatePassword() *string
	SetCertificatePassword(val *string)
	CertificatePasswordInput() *string
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
	Expiry() *string
	// Experimental.
	Fqn() *string
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyVaultId() *string
	SetKeyVaultId(val *string)
	KeyVaultIdInput() *string
	NegotiateClientCertificate() interface{}
	SetNegotiateClientCertificate(val interface{})
	NegotiateClientCertificateInput() interface{}
	SslKeyvaultIdentityClientId() *string
	SetSslKeyvaultIdentityClientId(val *string)
	SslKeyvaultIdentityClientIdInput() *string
	Subject() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thumbprint() *string
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
	ResetCertificate()
	ResetCertificatePassword()
	ResetKeyVaultId()
	ResetNegotiateClientCertificate()
	ResetSslKeyvaultIdentityClientId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementHostnameConfigurationPortalOutputReference
type jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) CertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) CertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) Expiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) KeyVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) KeyVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) NegotiateClientCertificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negotiateClientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) NegotiateClientCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negotiateClientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SslKeyvaultIdentityClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyvaultIdentityClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) SslKeyvaultIdentityClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyvaultIdentityClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}


func NewApiManagementHostnameConfigurationPortalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApiManagementHostnameConfigurationPortalOutputReference {
	_init_.Initialize()

	if err := validateNewApiManagementHostnameConfigurationPortalOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference{}

	_jsii_.Create(
		"azurerm.apiManagement.ApiManagementHostnameConfigurationPortalOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApiManagementHostnameConfigurationPortalOutputReference_Override(a ApiManagementHostnameConfigurationPortalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.apiManagement.ApiManagementHostnameConfigurationPortalOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference)SetCertificatePassword(val *string) {
	if err := j.validateSetCertificatePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificatePassword",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference)SetHostName(val *string) {
	if err := j.validateSetHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference)SetKeyVaultId(val *string) {
	if err := j.validateSetKeyVaultIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference)SetNegotiateClientCertificate(val interface{}) {
	if err := j.validateSetNegotiateClientCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negotiateClientCertificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference)SetSslKeyvaultIdentityClientId(val *string) {
	if err := j.validateSetSslKeyvaultIdentityClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslKeyvaultIdentityClientId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ResetCertificatePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificatePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ResetKeyVaultId() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyVaultId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ResetNegotiateClientCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetNegotiateClientCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ResetSslKeyvaultIdentityClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetSslKeyvaultIdentityClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationPortalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
