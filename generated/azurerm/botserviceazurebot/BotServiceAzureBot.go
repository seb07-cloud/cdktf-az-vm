package botserviceazurebot

import (
	_init_ "cdk.tf/go/stack/generated/azurerm/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"cdk.tf/go/stack/generated/azurerm/botserviceazurebot/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/bot_service_azure_bot azurerm_bot_service_azure_bot}.
type BotServiceAzureBot interface {
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
	DeveloperAppInsightsApiKey() *string
	SetDeveloperAppInsightsApiKey(val *string)
	DeveloperAppInsightsApiKeyInput() *string
	DeveloperAppInsightsApplicationId() *string
	SetDeveloperAppInsightsApplicationId(val *string)
	DeveloperAppInsightsApplicationIdInput() *string
	DeveloperAppInsightsKey() *string
	SetDeveloperAppInsightsKey(val *string)
	DeveloperAppInsightsKeyInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
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
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LuisAppIds() *[]*string
	SetLuisAppIds(val *[]*string)
	LuisAppIdsInput() *[]*string
	LuisKey() *string
	SetLuisKey(val *string)
	LuisKeyInput() *string
	MicrosoftAppId() *string
	SetMicrosoftAppId(val *string)
	MicrosoftAppIdInput() *string
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BotServiceAzureBotTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutTimeouts(value *BotServiceAzureBotTimeouts)
	ResetDeveloperAppInsightsApiKey()
	ResetDeveloperAppInsightsApplicationId()
	ResetDeveloperAppInsightsKey()
	ResetDisplayName()
	ResetEndpoint()
	ResetId()
	ResetLuisAppIds()
	ResetLuisKey()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for BotServiceAzureBot
type jsiiProxy_BotServiceAzureBot struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BotServiceAzureBot) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DeveloperAppInsightsApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DeveloperAppInsightsApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DeveloperAppInsightsApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DeveloperAppInsightsApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsApplicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DeveloperAppInsightsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DeveloperAppInsightsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) LuisAppIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"luisAppIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) LuisAppIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"luisAppIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) LuisKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"luisKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) LuisKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"luisKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) MicrosoftAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) MicrosoftAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) Timeouts() BotServiceAzureBotTimeoutsOutputReference {
	var returns BotServiceAzureBotTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotServiceAzureBot) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/bot_service_azure_bot azurerm_bot_service_azure_bot} Resource.
func NewBotServiceAzureBot(scope constructs.Construct, id *string, config *BotServiceAzureBotConfig) BotServiceAzureBot {
	_init_.Initialize()

	if err := validateNewBotServiceAzureBotParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BotServiceAzureBot{}

	_jsii_.Create(
		"azurerm.botServiceAzureBot.BotServiceAzureBot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/bot_service_azure_bot azurerm_bot_service_azure_bot} Resource.
func NewBotServiceAzureBot_Override(b BotServiceAzureBot, scope constructs.Construct, id *string, config *BotServiceAzureBotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.botServiceAzureBot.BotServiceAzureBot",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetDeveloperAppInsightsApiKey(val *string) {
	if err := j.validateSetDeveloperAppInsightsApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerAppInsightsApiKey",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetDeveloperAppInsightsApplicationId(val *string) {
	if err := j.validateSetDeveloperAppInsightsApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerAppInsightsApplicationId",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetDeveloperAppInsightsKey(val *string) {
	if err := j.validateSetDeveloperAppInsightsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerAppInsightsKey",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetLuisAppIds(val *[]*string) {
	if err := j.validateSetLuisAppIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"luisAppIds",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetLuisKey(val *string) {
	if err := j.validateSetLuisKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"luisKey",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetMicrosoftAppId(val *string) {
	if err := j.validateSetMicrosoftAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"microsoftAppId",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetSku(val *string) {
	if err := j.validateSetSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_BotServiceAzureBot)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
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
func BotServiceAzureBot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBotServiceAzureBot_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.botServiceAzureBot.BotServiceAzureBot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BotServiceAzureBot_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBotServiceAzureBot_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.botServiceAzureBot.BotServiceAzureBot",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BotServiceAzureBot_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBotServiceAzureBot_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.botServiceAzureBot.BotServiceAzureBot",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BotServiceAzureBot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"azurerm.botServiceAzureBot.BotServiceAzureBot",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BotServiceAzureBot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BotServiceAzureBot) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BotServiceAzureBot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BotServiceAzureBot) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BotServiceAzureBot) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BotServiceAzureBot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BotServiceAzureBot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BotServiceAzureBot) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BotServiceAzureBot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BotServiceAzureBot) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BotServiceAzureBot) PutTimeouts(value *BotServiceAzureBotTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetDeveloperAppInsightsApiKey() {
	_jsii_.InvokeVoid(
		b,
		"resetDeveloperAppInsightsApiKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetDeveloperAppInsightsApplicationId() {
	_jsii_.InvokeVoid(
		b,
		"resetDeveloperAppInsightsApplicationId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetDeveloperAppInsightsKey() {
	_jsii_.InvokeVoid(
		b,
		"resetDeveloperAppInsightsKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetDisplayName() {
	_jsii_.InvokeVoid(
		b,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetEndpoint() {
	_jsii_.InvokeVoid(
		b,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetLuisAppIds() {
	_jsii_.InvokeVoid(
		b,
		"resetLuisAppIds",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetLuisKey() {
	_jsii_.InvokeVoid(
		b,
		"resetLuisKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotServiceAzureBot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotServiceAzureBot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

