// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaischedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v19/googlevertexaischedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference interface {
	cdktn.ComplexObject
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
	CreateTime() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomEnvironmentSpec() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference
	CustomEnvironmentSpecInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec
	DataformRepositorySource() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference
	DataformRepositorySourceInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource
	DirectNotebookSource() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSourceOutputReference
	DirectNotebookSourceInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSource
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EncryptionSpec() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpecOutputReference
	EncryptionSpecInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpec
	ExecutionTimeout() *string
	SetExecutionTimeout(val *string)
	ExecutionTimeoutInput() *string
	ExecutionUser() *string
	SetExecutionUser(val *string)
	ExecutionUserInput() *string
	// Experimental.
	Fqn() *string
	GcsNotebookSource() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSourceOutputReference
	GcsNotebookSourceInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource
	GcsOutputUri() *string
	SetGcsOutputUri(val *string)
	GcsOutputUriInput() *string
	InternalValue() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob
	SetInternalValue(val *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob)
	JobState() *string
	KernelName() *string
	SetKernelName(val *string)
	KernelNameInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Name() *string
	NotebookRuntimeTemplateResourceName() *string
	SetNotebookRuntimeTemplateResourceName(val *string)
	NotebookRuntimeTemplateResourceNameInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	ScheduleResourceName() *string
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpdateTime() *string
	WorkbenchRuntime() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntimeOutputReference
	WorkbenchRuntimeInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntime
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutCustomEnvironmentSpec(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec)
	PutDataformRepositorySource(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource)
	PutDirectNotebookSource(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSource)
	PutEncryptionSpec(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpec)
	PutGcsNotebookSource(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource)
	PutWorkbenchRuntime(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntime)
	ResetCustomEnvironmentSpec()
	ResetDataformRepositorySource()
	ResetDirectNotebookSource()
	ResetDisplayName()
	ResetEncryptionSpec()
	ResetExecutionTimeout()
	ResetExecutionUser()
	ResetGcsNotebookSource()
	ResetGcsOutputUri()
	ResetKernelName()
	ResetLabels()
	ResetNotebookRuntimeTemplateResourceName()
	ResetParameters()
	ResetServiceAccount()
	ResetWorkbenchRuntime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference
type jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) CustomEnvironmentSpec() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference {
	var returns GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecOutputReference
	_jsii_.Get(
		j,
		"customEnvironmentSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) CustomEnvironmentSpecInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec {
	var returns *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec
	_jsii_.Get(
		j,
		"customEnvironmentSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DataformRepositorySource() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference {
	var returns GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference
	_jsii_.Get(
		j,
		"dataformRepositorySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DataformRepositorySourceInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource {
	var returns *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource
	_jsii_.Get(
		j,
		"dataformRepositorySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DirectNotebookSource() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSourceOutputReference {
	var returns GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSourceOutputReference
	_jsii_.Get(
		j,
		"directNotebookSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DirectNotebookSourceInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSource {
	var returns *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSource
	_jsii_.Get(
		j,
		"directNotebookSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) EncryptionSpec() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpecOutputReference {
	var returns GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpecOutputReference
	_jsii_.Get(
		j,
		"encryptionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) EncryptionSpecInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpec {
	var returns *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpec
	_jsii_.Get(
		j,
		"encryptionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsNotebookSource() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSourceOutputReference {
	var returns GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSourceOutputReference
	_jsii_.Get(
		j,
		"gcsNotebookSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsNotebookSourceInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource {
	var returns *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource
	_jsii_.Get(
		j,
		"gcsNotebookSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsOutputUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsOutputUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) InternalValue() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob {
	var returns *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) JobState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) KernelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) KernelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kernelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) NotebookRuntimeTemplateResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookRuntimeTemplateResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) NotebookRuntimeTemplateResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookRuntimeTemplateResourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ScheduleResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) WorkbenchRuntime() GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntimeOutputReference {
	var returns GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntimeOutputReference
	_jsii_.Get(
		j,
		"workbenchRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) WorkbenchRuntimeInput() *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntime {
	var returns *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntime
	_jsii_.Get(
		j,
		"workbenchRuntimeInput",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference_Override(g GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetExecutionTimeout(val *string) {
	if err := j.validateSetExecutionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetExecutionUser(val *string) {
	if err := j.validateSetExecutionUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionUser",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetGcsOutputUri(val *string) {
	if err := j.validateSetGcsOutputUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsOutputUri",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetInternalValue(val *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetKernelName(val *string) {
	if err := j.validateSetKernelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kernelName",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetNotebookRuntimeTemplateResourceName(val *string) {
	if err := j.validateSetNotebookRuntimeTemplateResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notebookRuntimeTemplateResourceName",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutCustomEnvironmentSpec(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec) {
	if err := g.validatePutCustomEnvironmentSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomEnvironmentSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutDataformRepositorySource(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource) {
	if err := g.validatePutDataformRepositorySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataformRepositorySource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutDirectNotebookSource(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSource) {
	if err := g.validatePutDirectNotebookSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDirectNotebookSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutEncryptionSpec(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpec) {
	if err := g.validatePutEncryptionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutGcsNotebookSource(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource) {
	if err := g.validatePutGcsNotebookSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGcsNotebookSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutWorkbenchRuntime(value *GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntime) {
	if err := g.validatePutWorkbenchRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkbenchRuntime",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetCustomEnvironmentSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomEnvironmentSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetDataformRepositorySource() {
	_jsii_.InvokeVoid(
		g,
		"resetDataformRepositorySource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetDirectNotebookSource() {
	_jsii_.InvokeVoid(
		g,
		"resetDirectNotebookSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetEncryptionSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetExecutionTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetExecutionTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetExecutionUser() {
	_jsii_.InvokeVoid(
		g,
		"resetExecutionUser",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetGcsNotebookSource() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsNotebookSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetGcsOutputUri() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsOutputUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetKernelName() {
	_jsii_.InvokeVoid(
		g,
		"resetKernelName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetNotebookRuntimeTemplateResourceName() {
	_jsii_.InvokeVoid(
		g,
		"resetNotebookRuntimeTemplateResourceName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetWorkbenchRuntime() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkbenchRuntime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

