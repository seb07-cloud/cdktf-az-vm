package containerregistrytask

import (
	_init_ "cdk.tf/go/stack/generated/azurerm/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"cdk.tf/go/stack/generated/azurerm/containerregistrytask/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/container_registry_task azurerm_container_registry_task}.
type ContainerRegistryTask interface {
	cdktf.TerraformResource
	AgentPoolName() *string
	SetAgentPoolName(val *string)
	AgentPoolNameInput() *string
	AgentSetting() ContainerRegistryTaskAgentSettingOutputReference
	AgentSettingInput() *ContainerRegistryTaskAgentSetting
	BaseImageTrigger() ContainerRegistryTaskBaseImageTriggerOutputReference
	BaseImageTriggerInput() *ContainerRegistryTaskBaseImageTrigger
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerRegistryId() *string
	SetContainerRegistryId(val *string)
	ContainerRegistryIdInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DockerStep() ContainerRegistryTaskDockerStepOutputReference
	DockerStepInput() *ContainerRegistryTaskDockerStep
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EncodedStep() ContainerRegistryTaskEncodedStepOutputReference
	EncodedStepInput() *ContainerRegistryTaskEncodedStep
	FileStep() ContainerRegistryTaskFileStepOutputReference
	FileStepInput() *ContainerRegistryTaskFileStep
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
	Identity() ContainerRegistryTaskIdentityOutputReference
	IdentityInput() *ContainerRegistryTaskIdentity
	IdInput() *string
	IsSystemTask() interface{}
	SetIsSystemTask(val interface{})
	IsSystemTaskInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogTemplate() *string
	SetLogTemplate(val *string)
	LogTemplateInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Platform() ContainerRegistryTaskPlatformOutputReference
	PlatformInput() *ContainerRegistryTaskPlatform
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
	RegistryCredential() ContainerRegistryTaskRegistryCredentialOutputReference
	RegistryCredentialInput() *ContainerRegistryTaskRegistryCredential
	SourceTrigger() ContainerRegistryTaskSourceTriggerList
	SourceTriggerInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeoutInSeconds() *float64
	SetTimeoutInSeconds(val *float64)
	TimeoutInSecondsInput() *float64
	Timeouts() ContainerRegistryTaskTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimerTrigger() ContainerRegistryTaskTimerTriggerList
	TimerTriggerInput() interface{}
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
	PutAgentSetting(value *ContainerRegistryTaskAgentSetting)
	PutBaseImageTrigger(value *ContainerRegistryTaskBaseImageTrigger)
	PutDockerStep(value *ContainerRegistryTaskDockerStep)
	PutEncodedStep(value *ContainerRegistryTaskEncodedStep)
	PutFileStep(value *ContainerRegistryTaskFileStep)
	PutIdentity(value *ContainerRegistryTaskIdentity)
	PutPlatform(value *ContainerRegistryTaskPlatform)
	PutRegistryCredential(value *ContainerRegistryTaskRegistryCredential)
	PutSourceTrigger(value interface{})
	PutTimeouts(value *ContainerRegistryTaskTimeouts)
	PutTimerTrigger(value interface{})
	ResetAgentPoolName()
	ResetAgentSetting()
	ResetBaseImageTrigger()
	ResetDockerStep()
	ResetEnabled()
	ResetEncodedStep()
	ResetFileStep()
	ResetId()
	ResetIdentity()
	ResetIsSystemTask()
	ResetLogTemplate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatform()
	ResetRegistryCredential()
	ResetSourceTrigger()
	ResetTags()
	ResetTimeoutInSeconds()
	ResetTimeouts()
	ResetTimerTrigger()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ContainerRegistryTask
type jsiiProxy_ContainerRegistryTask struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ContainerRegistryTask) AgentPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) AgentPoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentPoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) AgentSetting() ContainerRegistryTaskAgentSettingOutputReference {
	var returns ContainerRegistryTaskAgentSettingOutputReference
	_jsii_.Get(
		j,
		"agentSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) AgentSettingInput() *ContainerRegistryTaskAgentSetting {
	var returns *ContainerRegistryTaskAgentSetting
	_jsii_.Get(
		j,
		"agentSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) BaseImageTrigger() ContainerRegistryTaskBaseImageTriggerOutputReference {
	var returns ContainerRegistryTaskBaseImageTriggerOutputReference
	_jsii_.Get(
		j,
		"baseImageTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) BaseImageTriggerInput() *ContainerRegistryTaskBaseImageTrigger {
	var returns *ContainerRegistryTaskBaseImageTrigger
	_jsii_.Get(
		j,
		"baseImageTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) ContainerRegistryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) ContainerRegistryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) DockerStep() ContainerRegistryTaskDockerStepOutputReference {
	var returns ContainerRegistryTaskDockerStepOutputReference
	_jsii_.Get(
		j,
		"dockerStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) DockerStepInput() *ContainerRegistryTaskDockerStep {
	var returns *ContainerRegistryTaskDockerStep
	_jsii_.Get(
		j,
		"dockerStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) EncodedStep() ContainerRegistryTaskEncodedStepOutputReference {
	var returns ContainerRegistryTaskEncodedStepOutputReference
	_jsii_.Get(
		j,
		"encodedStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) EncodedStepInput() *ContainerRegistryTaskEncodedStep {
	var returns *ContainerRegistryTaskEncodedStep
	_jsii_.Get(
		j,
		"encodedStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) FileStep() ContainerRegistryTaskFileStepOutputReference {
	var returns ContainerRegistryTaskFileStepOutputReference
	_jsii_.Get(
		j,
		"fileStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) FileStepInput() *ContainerRegistryTaskFileStep {
	var returns *ContainerRegistryTaskFileStep
	_jsii_.Get(
		j,
		"fileStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Identity() ContainerRegistryTaskIdentityOutputReference {
	var returns ContainerRegistryTaskIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) IdentityInput() *ContainerRegistryTaskIdentity {
	var returns *ContainerRegistryTaskIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) IsSystemTask() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSystemTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) IsSystemTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSystemTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) LogTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) LogTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Platform() ContainerRegistryTaskPlatformOutputReference {
	var returns ContainerRegistryTaskPlatformOutputReference
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) PlatformInput() *ContainerRegistryTaskPlatform {
	var returns *ContainerRegistryTaskPlatform
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) RegistryCredential() ContainerRegistryTaskRegistryCredentialOutputReference {
	var returns ContainerRegistryTaskRegistryCredentialOutputReference
	_jsii_.Get(
		j,
		"registryCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) RegistryCredentialInput() *ContainerRegistryTaskRegistryCredential {
	var returns *ContainerRegistryTaskRegistryCredential
	_jsii_.Get(
		j,
		"registryCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) SourceTrigger() ContainerRegistryTaskSourceTriggerList {
	var returns ContainerRegistryTaskSourceTriggerList
	_jsii_.Get(
		j,
		"sourceTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) SourceTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) Timeouts() ContainerRegistryTaskTimeoutsOutputReference {
	var returns ContainerRegistryTaskTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TimerTrigger() ContainerRegistryTaskTimerTriggerList {
	var returns ContainerRegistryTaskTimerTriggerList
	_jsii_.Get(
		j,
		"timerTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistryTask) TimerTriggerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timerTriggerInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/container_registry_task azurerm_container_registry_task} Resource.
func NewContainerRegistryTask(scope constructs.Construct, id *string, config *ContainerRegistryTaskConfig) ContainerRegistryTask {
	_init_.Initialize()

	if err := validateNewContainerRegistryTaskParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerRegistryTask{}

	_jsii_.Create(
		"azurerm.containerRegistryTask.ContainerRegistryTask",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/container_registry_task azurerm_container_registry_task} Resource.
func NewContainerRegistryTask_Override(c ContainerRegistryTask, scope constructs.Construct, id *string, config *ContainerRegistryTaskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.containerRegistryTask.ContainerRegistryTask",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetAgentPoolName(val *string) {
	if err := j.validateSetAgentPoolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentPoolName",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetContainerRegistryId(val *string) {
	if err := j.validateSetContainerRegistryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRegistryId",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetIsSystemTask(val interface{}) {
	if err := j.validateSetIsSystemTaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSystemTask",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetLogTemplate(val *string) {
	if err := j.validateSetLogTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logTemplate",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistryTask)SetTimeoutInSeconds(val *float64) {
	if err := j.validateSetTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInSeconds",
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
func ContainerRegistryTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerRegistryTask_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.containerRegistryTask.ContainerRegistryTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerRegistryTask_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerRegistryTask_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.containerRegistryTask.ContainerRegistryTask",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerRegistryTask_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerRegistryTask_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.containerRegistryTask.ContainerRegistryTask",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ContainerRegistryTask_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"azurerm.containerRegistryTask.ContainerRegistryTask",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutAgentSetting(value *ContainerRegistryTaskAgentSetting) {
	if err := c.validatePutAgentSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAgentSetting",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutBaseImageTrigger(value *ContainerRegistryTaskBaseImageTrigger) {
	if err := c.validatePutBaseImageTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBaseImageTrigger",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutDockerStep(value *ContainerRegistryTaskDockerStep) {
	if err := c.validatePutDockerStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDockerStep",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutEncodedStep(value *ContainerRegistryTaskEncodedStep) {
	if err := c.validatePutEncodedStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEncodedStep",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutFileStep(value *ContainerRegistryTaskFileStep) {
	if err := c.validatePutFileStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFileStep",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutIdentity(value *ContainerRegistryTaskIdentity) {
	if err := c.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIdentity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutPlatform(value *ContainerRegistryTaskPlatform) {
	if err := c.validatePutPlatformParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPlatform",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutRegistryCredential(value *ContainerRegistryTaskRegistryCredential) {
	if err := c.validatePutRegistryCredentialParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRegistryCredential",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutSourceTrigger(value interface{}) {
	if err := c.validatePutSourceTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSourceTrigger",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutTimeouts(value *ContainerRegistryTaskTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) PutTimerTrigger(value interface{}) {
	if err := c.validatePutTimerTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimerTrigger",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetAgentPoolName() {
	_jsii_.InvokeVoid(
		c,
		"resetAgentPoolName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetAgentSetting() {
	_jsii_.InvokeVoid(
		c,
		"resetAgentSetting",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetBaseImageTrigger() {
	_jsii_.InvokeVoid(
		c,
		"resetBaseImageTrigger",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetDockerStep() {
	_jsii_.InvokeVoid(
		c,
		"resetDockerStep",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetEncodedStep() {
	_jsii_.InvokeVoid(
		c,
		"resetEncodedStep",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetFileStep() {
	_jsii_.InvokeVoid(
		c,
		"resetFileStep",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetIdentity() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetIsSystemTask() {
	_jsii_.InvokeVoid(
		c,
		"resetIsSystemTask",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetLogTemplate() {
	_jsii_.InvokeVoid(
		c,
		"resetLogTemplate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetPlatform() {
	_jsii_.InvokeVoid(
		c,
		"resetPlatform",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetRegistryCredential() {
	_jsii_.InvokeVoid(
		c,
		"resetRegistryCredential",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetSourceTrigger() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceTrigger",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) ResetTimerTrigger() {
	_jsii_.InvokeVoid(
		c,
		"resetTimerTrigger",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistryTask) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistryTask) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

