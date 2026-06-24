// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaischedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googlevertexaischedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_vertex_ai_schedule google_vertex_ai_schedule}.
type GoogleVertexAiSchedule interface {
	cdktn.TerraformResource
	AllowQueueing() interface{}
	SetAllowQueueing(val interface{})
	AllowQueueingInput() interface{}
	CatchUp() cdktn.IResolvable
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
	CreateNotebookExecutionJobRequest() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestOutputReference
	CreateNotebookExecutionJobRequestInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequest
	CreatePipelineJobRequest() GoogleVertexAiScheduleCreatePipelineJobRequestOutputReference
	CreatePipelineJobRequestInput() *GoogleVertexAiScheduleCreatePipelineJobRequest
	CreateTime() *string
	Cron() *string
	SetCron(val *string)
	CronInput() *string
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EndTime() *string
	SetEndTime(val *string)
	EndTimeInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LastPauseTime() *string
	LastResumeTime() *string
	LastScheduledRunResponse() GoogleVertexAiScheduleLastScheduledRunResponseList
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaxConcurrentActiveRunCount() *string
	SetMaxConcurrentActiveRunCount(val *string)
	MaxConcurrentActiveRunCountInput() *string
	MaxConcurrentRunCount() *string
	SetMaxConcurrentRunCount(val *string)
	MaxConcurrentRunCountInput() *string
	MaxRunCount() *string
	SetMaxRunCount(val *string)
	MaxRunCountInput() *string
	Name() *string
	NextRunTime() *string
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
	StartedRunCount() *string
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleVertexAiScheduleTimeoutsOutputReference
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
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
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
	PutCreateNotebookExecutionJobRequest(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequest)
	PutCreatePipelineJobRequest(value *GoogleVertexAiScheduleCreatePipelineJobRequest)
	PutTimeouts(value *GoogleVertexAiScheduleTimeouts)
	ResetAllowQueueing()
	ResetCreateNotebookExecutionJobRequest()
	ResetCreatePipelineJobRequest()
	ResetCron()
	ResetDeletionPolicy()
	ResetEndTime()
	ResetId()
	ResetMaxConcurrentActiveRunCount()
	ResetMaxRunCount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetStartTime()
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

// The jsii proxy struct for GoogleVertexAiSchedule
type jsiiProxy_GoogleVertexAiSchedule struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleVertexAiSchedule) AllowQueueing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowQueueing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) AllowQueueingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowQueueingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) CatchUp() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"catchUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) CreateNotebookExecutionJobRequest() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestOutputReference {
	var returns GoogleVertexAiScheduleCreateNotebookExecutionJobRequestOutputReference
	_jsii_.Get(
		j,
		"createNotebookExecutionJobRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) CreateNotebookExecutionJobRequestInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequest {
	var returns *GoogleVertexAiScheduleCreateNotebookExecutionJobRequest
	_jsii_.Get(
		j,
		"createNotebookExecutionJobRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) CreatePipelineJobRequest() GoogleVertexAiScheduleCreatePipelineJobRequestOutputReference {
	var returns GoogleVertexAiScheduleCreatePipelineJobRequestOutputReference
	_jsii_.Get(
		j,
		"createPipelineJobRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) CreatePipelineJobRequestInput() *GoogleVertexAiScheduleCreatePipelineJobRequest {
	var returns *GoogleVertexAiScheduleCreatePipelineJobRequest
	_jsii_.Get(
		j,
		"createPipelineJobRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) Cron() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cron",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) CronInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) EndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) LastPauseTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastPauseTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) LastResumeTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastResumeTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) LastScheduledRunResponse() GoogleVertexAiScheduleLastScheduledRunResponseList {
	var returns GoogleVertexAiScheduleLastScheduledRunResponseList
	_jsii_.Get(
		j,
		"lastScheduledRunResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) MaxConcurrentActiveRunCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrentActiveRunCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) MaxConcurrentActiveRunCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrentActiveRunCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) MaxConcurrentRunCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrentRunCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) MaxConcurrentRunCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxConcurrentRunCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) MaxRunCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRunCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) MaxRunCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRunCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) NextRunTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextRunTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) StartedRunCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startedRunCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) Timeouts() GoogleVertexAiScheduleTimeoutsOutputReference {
	var returns GoogleVertexAiScheduleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiSchedule) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_vertex_ai_schedule google_vertex_ai_schedule} Resource.
func NewGoogleVertexAiSchedule(scope constructs.Construct, id *string, config *GoogleVertexAiScheduleConfig) GoogleVertexAiSchedule {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiScheduleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiSchedule{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiSchedule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_vertex_ai_schedule google_vertex_ai_schedule} Resource.
func NewGoogleVertexAiSchedule_Override(g GoogleVertexAiSchedule, scope constructs.Construct, id *string, config *GoogleVertexAiScheduleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiSchedule",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetAllowQueueing(val interface{}) {
	if err := j.validateSetAllowQueueingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowQueueing",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetCron(val *string) {
	if err := j.validateSetCronParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cron",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetEndTime(val *string) {
	if err := j.validateSetEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetMaxConcurrentActiveRunCount(val *string) {
	if err := j.validateSetMaxConcurrentActiveRunCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentActiveRunCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetMaxConcurrentRunCount(val *string) {
	if err := j.validateSetMaxConcurrentRunCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentRunCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetMaxRunCount(val *string) {
	if err := j.validateSetMaxRunCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRunCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiSchedule)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

// Generates CDKTN code for importing a GoogleVertexAiSchedule resource upon running "cdktn plan <stack-name>".
func GoogleVertexAiSchedule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleVertexAiSchedule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiSchedule",
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
func GoogleVertexAiSchedule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiSchedule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiSchedule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleVertexAiSchedule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiSchedule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiSchedule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleVertexAiSchedule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVertexAiSchedule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiSchedule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleVertexAiSchedule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiSchedule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleVertexAiSchedule) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiSchedule) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiSchedule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiSchedule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiSchedule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiSchedule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiSchedule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiSchedule) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiSchedule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiSchedule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiSchedule) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) PutCreateNotebookExecutionJobRequest(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequest) {
	if err := g.validatePutCreateNotebookExecutionJobRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCreateNotebookExecutionJobRequest",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) PutCreatePipelineJobRequest(value *GoogleVertexAiScheduleCreatePipelineJobRequest) {
	if err := g.validatePutCreatePipelineJobRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCreatePipelineJobRequest",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) PutTimeouts(value *GoogleVertexAiScheduleTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ResetAllowQueueing() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowQueueing",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ResetCreateNotebookExecutionJobRequest() {
	_jsii_.InvokeVoid(
		g,
		"resetCreateNotebookExecutionJobRequest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ResetCreatePipelineJobRequest() {
	_jsii_.InvokeVoid(
		g,
		"resetCreatePipelineJobRequest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ResetCron() {
	_jsii_.InvokeVoid(
		g,
		"resetCron",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ResetEndTime() {
	_jsii_.InvokeVoid(
		g,
		"resetEndTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ResetMaxConcurrentActiveRunCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConcurrentActiveRunCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ResetMaxRunCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxRunCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ResetStartTime() {
	_jsii_.InvokeVoid(
		g,
		"resetStartTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiSchedule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiSchedule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiSchedule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiSchedule) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

