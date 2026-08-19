// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlecestool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool google_ces_tool}.
type GoogleCesTool interface {
	cdktn.TerraformResource
	AgentTool() GoogleCesToolAgentToolOutputReference
	AgentToolInput() *GoogleCesToolAgentTool
	App() *string
	SetApp(val *string)
	AppInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientFunction() GoogleCesToolClientFunctionOutputReference
	ClientFunctionInput() *GoogleCesToolClientFunction
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectorTool() GoogleCesToolConnectorToolList
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DataStoreTool() GoogleCesToolDataStoreToolOutputReference
	DataStoreToolInput() *GoogleCesToolDataStoreTool
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	Etag() *string
	ExecutionType() *string
	SetExecutionType(val *string)
	ExecutionTypeInput() *string
	FileSearchTool() GoogleCesToolFileSearchToolOutputReference
	FileSearchToolInput() *GoogleCesToolFileSearchTool
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeneratedSummary() *string
	GoogleSearchTool() GoogleCesToolGoogleSearchToolOutputReference
	GoogleSearchToolInput() *GoogleCesToolGoogleSearchTool
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	McpTool() GoogleCesToolMcpToolList
	Name() *string
	// The tree node.
	Node() constructs.Node
	OpenApiTool() GoogleCesToolOpenApiToolList
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PythonFunction() GoogleCesToolPythonFunctionOutputReference
	PythonFunctionInput() *GoogleCesToolPythonFunction
	// Experimental.
	RawOverrides() interface{}
	RemoteAgentTool() GoogleCesToolRemoteAgentToolList
	SystemTool() GoogleCesToolSystemToolList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
	Timeouts() GoogleCesToolTimeoutsOutputReference
	TimeoutsInput() interface{}
	ToolFakeConfig() GoogleCesToolToolFakeConfigOutputReference
	ToolFakeConfigInput() *GoogleCesToolToolFakeConfig
	ToolId() *string
	SetToolId(val *string)
	ToolIdInput() *string
	UpdateTime() *string
	WidgetTool() GoogleCesToolWidgetToolOutputReference
	WidgetToolInput() *GoogleCesToolWidgetTool
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAgentTool(value *GoogleCesToolAgentTool)
	PutClientFunction(value *GoogleCesToolClientFunction)
	PutDataStoreTool(value *GoogleCesToolDataStoreTool)
	PutFileSearchTool(value *GoogleCesToolFileSearchTool)
	PutGoogleSearchTool(value *GoogleCesToolGoogleSearchTool)
	PutPythonFunction(value *GoogleCesToolPythonFunction)
	PutTimeouts(value *GoogleCesToolTimeouts)
	PutToolFakeConfig(value *GoogleCesToolToolFakeConfig)
	PutWidgetTool(value *GoogleCesToolWidgetTool)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetAgentTool()
	ResetClientFunction()
	ResetDataStoreTool()
	ResetDeletionPolicy()
	ResetExecutionType()
	ResetFileSearchTool()
	ResetGoogleSearchTool()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPythonFunction()
	ResetTimeout()
	ResetTimeouts()
	ResetToolFakeConfig()
	ResetWidgetTool()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for GoogleCesTool
type jsiiProxy_GoogleCesTool struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleCesTool) AgentTool() GoogleCesToolAgentToolOutputReference {
	var returns GoogleCesToolAgentToolOutputReference
	_jsii_.Get(
		j,
		"agentTool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) AgentToolInput() *GoogleCesToolAgentTool {
	var returns *GoogleCesToolAgentTool
	_jsii_.Get(
		j,
		"agentToolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) App() *string {
	var returns *string
	_jsii_.Get(
		j,
		"app",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) AppInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) ClientFunction() GoogleCesToolClientFunctionOutputReference {
	var returns GoogleCesToolClientFunctionOutputReference
	_jsii_.Get(
		j,
		"clientFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) ClientFunctionInput() *GoogleCesToolClientFunction {
	var returns *GoogleCesToolClientFunction
	_jsii_.Get(
		j,
		"clientFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) ConnectorTool() GoogleCesToolConnectorToolList {
	var returns GoogleCesToolConnectorToolList
	_jsii_.Get(
		j,
		"connectorTool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) DataStoreTool() GoogleCesToolDataStoreToolOutputReference {
	var returns GoogleCesToolDataStoreToolOutputReference
	_jsii_.Get(
		j,
		"dataStoreTool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) DataStoreToolInput() *GoogleCesToolDataStoreTool {
	var returns *GoogleCesToolDataStoreTool
	_jsii_.Get(
		j,
		"dataStoreToolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) ExecutionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) ExecutionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) FileSearchTool() GoogleCesToolFileSearchToolOutputReference {
	var returns GoogleCesToolFileSearchToolOutputReference
	_jsii_.Get(
		j,
		"fileSearchTool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) FileSearchToolInput() *GoogleCesToolFileSearchTool {
	var returns *GoogleCesToolFileSearchTool
	_jsii_.Get(
		j,
		"fileSearchToolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) GeneratedSummary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generatedSummary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) GoogleSearchTool() GoogleCesToolGoogleSearchToolOutputReference {
	var returns GoogleCesToolGoogleSearchToolOutputReference
	_jsii_.Get(
		j,
		"googleSearchTool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) GoogleSearchToolInput() *GoogleCesToolGoogleSearchTool {
	var returns *GoogleCesToolGoogleSearchTool
	_jsii_.Get(
		j,
		"googleSearchToolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) McpTool() GoogleCesToolMcpToolList {
	var returns GoogleCesToolMcpToolList
	_jsii_.Get(
		j,
		"mcpTool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) OpenApiTool() GoogleCesToolOpenApiToolList {
	var returns GoogleCesToolOpenApiToolList
	_jsii_.Get(
		j,
		"openApiTool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) PythonFunction() GoogleCesToolPythonFunctionOutputReference {
	var returns GoogleCesToolPythonFunctionOutputReference
	_jsii_.Get(
		j,
		"pythonFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) PythonFunctionInput() *GoogleCesToolPythonFunction {
	var returns *GoogleCesToolPythonFunction
	_jsii_.Get(
		j,
		"pythonFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) RemoteAgentTool() GoogleCesToolRemoteAgentToolList {
	var returns GoogleCesToolRemoteAgentToolList
	_jsii_.Get(
		j,
		"remoteAgentTool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) SystemTool() GoogleCesToolSystemToolList {
	var returns GoogleCesToolSystemToolList
	_jsii_.Get(
		j,
		"systemTool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) Timeouts() GoogleCesToolTimeoutsOutputReference {
	var returns GoogleCesToolTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) ToolFakeConfig() GoogleCesToolToolFakeConfigOutputReference {
	var returns GoogleCesToolToolFakeConfigOutputReference
	_jsii_.Get(
		j,
		"toolFakeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) ToolFakeConfigInput() *GoogleCesToolToolFakeConfig {
	var returns *GoogleCesToolToolFakeConfig
	_jsii_.Get(
		j,
		"toolFakeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) ToolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) ToolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"toolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) WidgetTool() GoogleCesToolWidgetToolOutputReference {
	var returns GoogleCesToolWidgetToolOutputReference
	_jsii_.Get(
		j,
		"widgetTool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesTool) WidgetToolInput() *GoogleCesToolWidgetTool {
	var returns *GoogleCesToolWidgetTool
	_jsii_.Get(
		j,
		"widgetToolInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool google_ces_tool} Resource.
func NewGoogleCesTool(scope constructs.Construct, id *string, config *GoogleCesToolConfig) GoogleCesTool {
	_init_.Initialize()

	if err := validateNewGoogleCesToolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesTool{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesTool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_ces_tool google_ces_tool} Resource.
func NewGoogleCesTool_Override(g GoogleCesTool, scope constructs.Construct, id *string, config *GoogleCesToolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesTool",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetApp(val *string) {
	if err := j.validateSetAppParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"app",
		val,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetExecutionType(val *string) {
	if err := j.validateSetExecutionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionType",
		val,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_GoogleCesTool)SetToolId(val *string) {
	if err := j.validateSetToolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"toolId",
		val,
	)
}

// Generates CDKTN code for importing a GoogleCesTool resource upon running "cdktn plan <stack-name>".
func GoogleCesTool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleCesTool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesTool",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func GoogleCesTool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCesTool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesTool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCesTool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCesTool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesTool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCesTool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCesTool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesTool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleCesTool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleCesTool.GoogleCesTool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleCesTool) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleCesTool) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleCesTool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleCesTool) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := g.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleCesTool) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleCesTool) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleCesTool) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleCesTool) PutAgentTool(value *GoogleCesToolAgentTool) {
	if err := g.validatePutAgentToolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAgentTool",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesTool) PutClientFunction(value *GoogleCesToolClientFunction) {
	if err := g.validatePutClientFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientFunction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesTool) PutDataStoreTool(value *GoogleCesToolDataStoreTool) {
	if err := g.validatePutDataStoreToolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataStoreTool",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesTool) PutFileSearchTool(value *GoogleCesToolFileSearchTool) {
	if err := g.validatePutFileSearchToolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFileSearchTool",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesTool) PutGoogleSearchTool(value *GoogleCesToolGoogleSearchTool) {
	if err := g.validatePutGoogleSearchToolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoogleSearchTool",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesTool) PutPythonFunction(value *GoogleCesToolPythonFunction) {
	if err := g.validatePutPythonFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPythonFunction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesTool) PutTimeouts(value *GoogleCesToolTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesTool) PutToolFakeConfig(value *GoogleCesToolToolFakeConfig) {
	if err := g.validatePutToolFakeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolFakeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesTool) PutWidgetTool(value *GoogleCesToolWidgetTool) {
	if err := g.validatePutWidgetToolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWidgetTool",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesTool) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetAgentTool() {
	_jsii_.InvokeVoid(
		g,
		"resetAgentTool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetClientFunction() {
	_jsii_.InvokeVoid(
		g,
		"resetClientFunction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetDataStoreTool() {
	_jsii_.InvokeVoid(
		g,
		"resetDataStoreTool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetExecutionType() {
	_jsii_.InvokeVoid(
		g,
		"resetExecutionType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetFileSearchTool() {
	_jsii_.InvokeVoid(
		g,
		"resetFileSearchTool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetGoogleSearchTool() {
	_jsii_.InvokeVoid(
		g,
		"resetGoogleSearchTool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetPythonFunction() {
	_jsii_.InvokeVoid(
		g,
		"resetPythonFunction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetToolFakeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetToolFakeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) ResetWidgetTool() {
	_jsii_.InvokeVoid(
		g,
		"resetWidgetTool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesTool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesTool) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		g,
		"with",
		args,
		&returns,
	)

	return returns
}

