// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleagenticapplicationsanalystagentpersona/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona google_agentic_applications_analyst_agent_persona}.
type GoogleAgenticApplicationsAnalystAgentPersona interface {
	cdktn.TerraformResource
	AnalystAgentPersonaId() *string
	SetAnalystAgentPersonaId(val *string)
	AnalystAgentPersonaIdInput() *string
	ArtifactExamples() GoogleAgenticApplicationsAnalystAgentPersonaArtifactExamplesList
	ArtifactExamplesInput() interface{}
	ArtifactsConfig() GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigOutputReference
	ArtifactsConfigInput() *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfig
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	CreateTime() *string
	CustomerContext() *[]*string
	SetCustomerContext(val *[]*string)
	CustomerContextInput() *[]*string
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayDescription() *string
	SetDisplayDescription(val *string)
	DisplayDescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	ExternalDataSources() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesList
	ExternalDataSourcesInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeminiEnterpriseEngine() *string
	SetGeminiEnterpriseEngine(val *string)
	GeminiEnterpriseEngineInput() *string
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
	McpDataSources() GoogleAgenticApplicationsAnalystAgentPersonaMcpDataSourcesList
	McpDataSourcesInput() interface{}
	ModelDescription() *string
	SetModelDescription(val *string)
	ModelDescriptionInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
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
	// Experimental.
	RawOverrides() interface{}
	Resources() GoogleAgenticApplicationsAnalystAgentPersonaResourcesList
	ResourcesInput() interface{}
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	Skills() GoogleAgenticApplicationsAnalystAgentPersonaSkillsList
	SkillsInput() interface{}
	Tables() GoogleAgenticApplicationsAnalystAgentPersonaTablesList
	TablesInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleAgenticApplicationsAnalystAgentPersonaTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
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
	PutArtifactExamples(value interface{})
	PutArtifactsConfig(value *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfig)
	PutExternalDataSources(value interface{})
	PutMcpDataSources(value interface{})
	PutResources(value interface{})
	PutSkills(value interface{})
	PutTables(value interface{})
	PutTimeouts(value *GoogleAgenticApplicationsAnalystAgentPersonaTimeouts)
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
	ResetArtifactExamples()
	ResetArtifactsConfig()
	ResetCustomerContext()
	ResetDeletionPolicy()
	ResetDisplayDescription()
	ResetExternalDataSources()
	ResetGeminiEnterpriseEngine()
	ResetId()
	ResetMcpDataSources()
	ResetModelDescription()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetResources()
	ResetRole()
	ResetSkills()
	ResetTables()
	ResetTimeouts()
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

// The jsii proxy struct for GoogleAgenticApplicationsAnalystAgentPersona
type jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) AnalystAgentPersonaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analystAgentPersonaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) AnalystAgentPersonaIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analystAgentPersonaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ArtifactExamples() GoogleAgenticApplicationsAnalystAgentPersonaArtifactExamplesList {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaArtifactExamplesList
	_jsii_.Get(
		j,
		"artifactExamples",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ArtifactExamplesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"artifactExamplesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ArtifactsConfig() GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigOutputReference {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigOutputReference
	_jsii_.Get(
		j,
		"artifactsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ArtifactsConfigInput() *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfig {
	var returns *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfig
	_jsii_.Get(
		j,
		"artifactsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) CustomerContext() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customerContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) CustomerContextInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customerContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) DisplayDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) DisplayDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ExternalDataSources() GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesList {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaExternalDataSourcesList
	_jsii_.Get(
		j,
		"externalDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ExternalDataSourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalDataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) GeminiEnterpriseEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geminiEnterpriseEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) GeminiEnterpriseEngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geminiEnterpriseEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) McpDataSources() GoogleAgenticApplicationsAnalystAgentPersonaMcpDataSourcesList {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaMcpDataSourcesList
	_jsii_.Get(
		j,
		"mcpDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) McpDataSourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mcpDataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ModelDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Resources() GoogleAgenticApplicationsAnalystAgentPersonaResourcesList {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaResourcesList
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Skills() GoogleAgenticApplicationsAnalystAgentPersonaSkillsList {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaSkillsList
	_jsii_.Get(
		j,
		"skills",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) SkillsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skillsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Tables() GoogleAgenticApplicationsAnalystAgentPersonaTablesList {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaTablesList
	_jsii_.Get(
		j,
		"tables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) TablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) Timeouts() GoogleAgenticApplicationsAnalystAgentPersonaTimeoutsOutputReference {
	var returns GoogleAgenticApplicationsAnalystAgentPersonaTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona google_agentic_applications_analyst_agent_persona} Resource.
func NewGoogleAgenticApplicationsAnalystAgentPersona(scope constructs.Construct, id *string, config *GoogleAgenticApplicationsAnalystAgentPersonaConfig) GoogleAgenticApplicationsAnalystAgentPersona {
	_init_.Initialize()

	if err := validateNewGoogleAgenticApplicationsAnalystAgentPersonaParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAgenticApplicationsAnalystAgentPersona.GoogleAgenticApplicationsAnalystAgentPersona",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona google_agentic_applications_analyst_agent_persona} Resource.
func NewGoogleAgenticApplicationsAnalystAgentPersona_Override(g GoogleAgenticApplicationsAnalystAgentPersona, scope constructs.Construct, id *string, config *GoogleAgenticApplicationsAnalystAgentPersonaConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAgenticApplicationsAnalystAgentPersona.GoogleAgenticApplicationsAnalystAgentPersona",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetAnalystAgentPersonaId(val *string) {
	if err := j.validateSetAnalystAgentPersonaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analystAgentPersonaId",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetCustomerContext(val *[]*string) {
	if err := j.validateSetCustomerContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerContext",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetDisplayDescription(val *string) {
	if err := j.validateSetDisplayDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayDescription",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetGeminiEnterpriseEngine(val *string) {
	if err := j.validateSetGeminiEnterpriseEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geminiEnterpriseEngine",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetModelDescription(val *string) {
	if err := j.validateSetModelDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDescription",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

// Generates CDKTN code for importing a GoogleAgenticApplicationsAnalystAgentPersona resource upon running "cdktn plan <stack-name>".
func GoogleAgenticApplicationsAnalystAgentPersona_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleAgenticApplicationsAnalystAgentPersona_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleAgenticApplicationsAnalystAgentPersona.GoogleAgenticApplicationsAnalystAgentPersona",
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
func GoogleAgenticApplicationsAnalystAgentPersona_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAgenticApplicationsAnalystAgentPersona_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleAgenticApplicationsAnalystAgentPersona.GoogleAgenticApplicationsAnalystAgentPersona",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAgenticApplicationsAnalystAgentPersona_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAgenticApplicationsAnalystAgentPersona_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleAgenticApplicationsAnalystAgentPersona.GoogleAgenticApplicationsAnalystAgentPersona",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAgenticApplicationsAnalystAgentPersona_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAgenticApplicationsAnalystAgentPersona_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleAgenticApplicationsAnalystAgentPersona.GoogleAgenticApplicationsAnalystAgentPersona",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleAgenticApplicationsAnalystAgentPersona_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleAgenticApplicationsAnalystAgentPersona.GoogleAgenticApplicationsAnalystAgentPersona",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) PutArtifactExamples(value interface{}) {
	if err := g.validatePutArtifactExamplesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putArtifactExamples",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) PutArtifactsConfig(value *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfig) {
	if err := g.validatePutArtifactsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putArtifactsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) PutExternalDataSources(value interface{}) {
	if err := g.validatePutExternalDataSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExternalDataSources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) PutMcpDataSources(value interface{}) {
	if err := g.validatePutMcpDataSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMcpDataSources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) PutResources(value interface{}) {
	if err := g.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) PutSkills(value interface{}) {
	if err := g.validatePutSkillsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSkills",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) PutTables(value interface{}) {
	if err := g.validatePutTablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTables",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) PutTimeouts(value *GoogleAgenticApplicationsAnalystAgentPersonaTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetArtifactExamples() {
	_jsii_.InvokeVoid(
		g,
		"resetArtifactExamples",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetArtifactsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetArtifactsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetCustomerContext() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomerContext",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetDisplayDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetExternalDataSources() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalDataSources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetGeminiEnterpriseEngine() {
	_jsii_.InvokeVoid(
		g,
		"resetGeminiEnterpriseEngine",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetMcpDataSources() {
	_jsii_.InvokeVoid(
		g,
		"resetMcpDataSources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetModelDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetModelDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetResources() {
	_jsii_.InvokeVoid(
		g,
		"resetResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetRole() {
	_jsii_.InvokeVoid(
		g,
		"resetRole",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetSkills() {
	_jsii_.InvokeVoid(
		g,
		"resetSkills",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetTables() {
	_jsii_.InvokeVoid(
		g,
		"resetTables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgenticApplicationsAnalystAgentPersona) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

