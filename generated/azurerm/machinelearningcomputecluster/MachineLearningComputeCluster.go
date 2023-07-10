package machinelearningcomputecluster

import (
	_init_ "cdk.tf/go/stack/generated/azurerm/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"cdk.tf/go/stack/generated/azurerm/machinelearningcomputecluster/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/machine_learning_compute_cluster azurerm_machine_learning_compute_cluster}.
type MachineLearningComputeCluster interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	Identity() MachineLearningComputeClusterIdentityOutputReference
	IdentityInput() *MachineLearningComputeClusterIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalAuthEnabled() interface{}
	SetLocalAuthEnabled(val interface{})
	LocalAuthEnabledInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MachineLearningWorkspaceId() *string
	SetMachineLearningWorkspaceId(val *string)
	MachineLearningWorkspaceIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ScaleSettings() MachineLearningComputeClusterScaleSettingsOutputReference
	ScaleSettingsInput() *MachineLearningComputeClusterScaleSettings
	Ssh() MachineLearningComputeClusterSshOutputReference
	SshInput() *MachineLearningComputeClusterSsh
	SshPublicAccessEnabled() interface{}
	SetSshPublicAccessEnabled(val interface{})
	SshPublicAccessEnabledInput() interface{}
	SubnetResourceId() *string
	SetSubnetResourceId(val *string)
	SubnetResourceIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MachineLearningComputeClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	VmPriority() *string
	SetVmPriority(val *string)
	VmPriorityInput() *string
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutIdentity(value *MachineLearningComputeClusterIdentity)
	PutScaleSettings(value *MachineLearningComputeClusterScaleSettings)
	PutSsh(value *MachineLearningComputeClusterSsh)
	PutTimeouts(value *MachineLearningComputeClusterTimeouts)
	ResetDescription()
	ResetId()
	ResetIdentity()
	ResetLocalAuthEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSsh()
	ResetSshPublicAccessEnabled()
	ResetSubnetResourceId()
	ResetTags()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for MachineLearningComputeCluster
type jsiiProxy_MachineLearningComputeCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MachineLearningComputeCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Identity() MachineLearningComputeClusterIdentityOutputReference {
	var returns MachineLearningComputeClusterIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) IdentityInput() *MachineLearningComputeClusterIdentity {
	var returns *MachineLearningComputeClusterIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) LocalAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) LocalAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) MachineLearningWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineLearningWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) MachineLearningWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineLearningWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) ScaleSettings() MachineLearningComputeClusterScaleSettingsOutputReference {
	var returns MachineLearningComputeClusterScaleSettingsOutputReference
	_jsii_.Get(
		j,
		"scaleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) ScaleSettingsInput() *MachineLearningComputeClusterScaleSettings {
	var returns *MachineLearningComputeClusterScaleSettings
	_jsii_.Get(
		j,
		"scaleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Ssh() MachineLearningComputeClusterSshOutputReference {
	var returns MachineLearningComputeClusterSshOutputReference
	_jsii_.Get(
		j,
		"ssh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) SshInput() *MachineLearningComputeClusterSsh {
	var returns *MachineLearningComputeClusterSsh
	_jsii_.Get(
		j,
		"sshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) SshPublicAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sshPublicAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) SshPublicAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sshPublicAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) SubnetResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) SubnetResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) Timeouts() MachineLearningComputeClusterTimeoutsOutputReference {
	var returns MachineLearningComputeClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) VmPriority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) VmPriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningComputeCluster) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/machine_learning_compute_cluster azurerm_machine_learning_compute_cluster} Resource.
func NewMachineLearningComputeCluster(scope constructs.Construct, id *string, config *MachineLearningComputeClusterConfig) MachineLearningComputeCluster {
	_init_.Initialize()

	if err := validateNewMachineLearningComputeClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MachineLearningComputeCluster{}

	_jsii_.Create(
		"azurerm.machineLearningComputeCluster.MachineLearningComputeCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/machine_learning_compute_cluster azurerm_machine_learning_compute_cluster} Resource.
func NewMachineLearningComputeCluster_Override(m MachineLearningComputeCluster, scope constructs.Construct, id *string, config *MachineLearningComputeClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.machineLearningComputeCluster.MachineLearningComputeCluster",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetLocalAuthEnabled(val interface{}) {
	if err := j.validateSetLocalAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetMachineLearningWorkspaceId(val *string) {
	if err := j.validateSetMachineLearningWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineLearningWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetSshPublicAccessEnabled(val interface{}) {
	if err := j.validateSetSshPublicAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetSubnetResourceId(val *string) {
	if err := j.validateSetSubnetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetResourceId",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetVmPriority(val *string) {
	if err := j.validateSetVmPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmPriority",
		val,
	)
}

func (j *jsiiProxy_MachineLearningComputeCluster)SetVmSize(val *string) {
	if err := j.validateSetVmSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func MachineLearningComputeCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMachineLearningComputeCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.machineLearningComputeCluster.MachineLearningComputeCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MachineLearningComputeCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMachineLearningComputeCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.machineLearningComputeCluster.MachineLearningComputeCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MachineLearningComputeCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMachineLearningComputeCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.machineLearningComputeCluster.MachineLearningComputeCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MachineLearningComputeCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"azurerm.machineLearningComputeCluster.MachineLearningComputeCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MachineLearningComputeCluster) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningComputeCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningComputeCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningComputeCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningComputeCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningComputeCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningComputeCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningComputeCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningComputeCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningComputeCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningComputeCluster) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) PutIdentity(value *MachineLearningComputeClusterIdentity) {
	if err := m.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIdentity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) PutScaleSettings(value *MachineLearningComputeClusterScaleSettings) {
	if err := m.validatePutScaleSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putScaleSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) PutSsh(value *MachineLearningComputeClusterSsh) {
	if err := m.validatePutSshParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSsh",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) PutTimeouts(value *MachineLearningComputeClusterTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) ResetIdentity() {
	_jsii_.InvokeVoid(
		m,
		"resetIdentity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) ResetLocalAuthEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetLocalAuthEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) ResetSsh() {
	_jsii_.InvokeVoid(
		m,
		"resetSsh",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) ResetSshPublicAccessEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetSshPublicAccessEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) ResetSubnetResourceId() {
	_jsii_.InvokeVoid(
		m,
		"resetSubnetResourceId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningComputeCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningComputeCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningComputeCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningComputeCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

