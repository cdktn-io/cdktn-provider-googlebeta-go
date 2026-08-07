// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaischedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlevertexaischedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference interface {
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
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EncryptionSpec() GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpecOutputReference
	EncryptionSpecInput() *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec
	EndTime() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJob
	SetInternalValue(val *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJob)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Name() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	PipelineSpec() *string
	SetPipelineSpec(val *string)
	PipelineSpecInput() *string
	PreflightValidations() interface{}
	SetPreflightValidations(val interface{})
	PreflightValidationsInput() interface{}
	PscInterfaceConfig() GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference
	PscInterfaceConfigInput() *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig
	ReservedIpRanges() *[]*string
	SetReservedIpRanges(val *[]*string)
	ReservedIpRangesInput() *[]*string
	RuntimeConfig() GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference
	RuntimeConfigInput() *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig
	ScheduleName() *string
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	StartTime() *string
	State() *string
	TemplateMetadata() GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList
	TemplateUri() *string
	SetTemplateUri(val *string)
	TemplateUriInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UpdateTime() *string
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
	PutEncryptionSpec(value *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec)
	PutPscInterfaceConfig(value *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig)
	PutRuntimeConfig(value *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig)
	ResetDisplayName()
	ResetEncryptionSpec()
	ResetLabels()
	ResetNetwork()
	ResetPipelineSpec()
	ResetPreflightValidations()
	ResetPscInterfaceConfig()
	ResetReservedIpRanges()
	ResetRuntimeConfig()
	ResetServiceAccount()
	ResetTemplateUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference
type jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) EncryptionSpec() GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpecOutputReference {
	var returns GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpecOutputReference
	_jsii_.Get(
		j,
		"encryptionSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) EncryptionSpecInput() *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec {
	var returns *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec
	_jsii_.Get(
		j,
		"encryptionSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) InternalValue() *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJob {
	var returns *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PipelineSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PipelineSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PreflightValidations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preflightValidations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PreflightValidationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preflightValidationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PscInterfaceConfig() GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference {
	var returns GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfigOutputReference
	_jsii_.Get(
		j,
		"pscInterfaceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PscInterfaceConfigInput() *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig {
	var returns *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig
	_jsii_.Get(
		j,
		"pscInterfaceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ReservedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ReservedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) RuntimeConfig() GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference {
	var returns GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfigOutputReference
	_jsii_.Get(
		j,
		"runtimeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) RuntimeConfigInput() *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig {
	var returns *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig
	_jsii_.Get(
		j,
		"runtimeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ScheduleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) TemplateMetadata() GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList {
	var returns GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobTemplateMetadataList
	_jsii_.Get(
		j,
		"templateMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) TemplateUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) TemplateUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference_Override(g GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiSchedule.GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetInternalValue(val *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetPipelineSpec(val *string) {
	if err := j.validateSetPipelineSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineSpec",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetPreflightValidations(val interface{}) {
	if err := j.validateSetPreflightValidationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preflightValidations",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetReservedIpRanges(val *[]*string) {
	if err := j.validateSetReservedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedIpRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetTemplateUri(val *string) {
	if err := j.validateSetTemplateUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateUri",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PutEncryptionSpec(value *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobEncryptionSpec) {
	if err := g.validatePutEncryptionSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PutPscInterfaceConfig(value *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobPscInterfaceConfig) {
	if err := g.validatePutPscInterfaceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPscInterfaceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) PutRuntimeConfig(value *GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobRuntimeConfig) {
	if err := g.validatePutRuntimeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRuntimeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetEncryptionSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetPipelineSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetPipelineSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetPreflightValidations() {
	_jsii_.InvokeVoid(
		g,
		"resetPreflightValidations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetPscInterfaceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPscInterfaceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetReservedIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetReservedIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetRuntimeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ResetTemplateUri() {
	_jsii_.InvokeVoid(
		g,
		"resetTemplateUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiScheduleCreatePipelineJobRequestPipelineJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

