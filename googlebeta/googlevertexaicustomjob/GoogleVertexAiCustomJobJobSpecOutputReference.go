// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaicustomjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlevertexaicustomjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiCustomJobJobSpecOutputReference interface {
	cdktn.ComplexObject
	BaseOutputDirectory() GoogleVertexAiCustomJobJobSpecBaseOutputDirectoryOutputReference
	BaseOutputDirectoryInput() *GoogleVertexAiCustomJobJobSpecBaseOutputDirectory
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
	EnableDashboardAccess() interface{}
	SetEnableDashboardAccess(val interface{})
	EnableDashboardAccessInput() interface{}
	EnableWebAccess() interface{}
	SetEnableWebAccess(val interface{})
	EnableWebAccessInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleVertexAiCustomJobJobSpec
	SetInternalValue(val *GoogleVertexAiCustomJobJobSpec)
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	PscInterfaceConfig() GoogleVertexAiCustomJobJobSpecPscInterfaceConfigOutputReference
	PscInterfaceConfigInput() *GoogleVertexAiCustomJobJobSpecPscInterfaceConfig
	ReservedIpRanges() *[]*string
	SetReservedIpRanges(val *[]*string)
	ReservedIpRangesInput() *[]*string
	Scheduling() GoogleVertexAiCustomJobJobSpecSchedulingOutputReference
	SchedulingInput() *GoogleVertexAiCustomJobJobSpecScheduling
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	Tensorboard() *string
	SetTensorboard(val *string)
	TensorboardInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkerPoolSpecs() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsList
	WorkerPoolSpecsInput() interface{}
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
	PutBaseOutputDirectory(value *GoogleVertexAiCustomJobJobSpecBaseOutputDirectory)
	PutPscInterfaceConfig(value *GoogleVertexAiCustomJobJobSpecPscInterfaceConfig)
	PutScheduling(value *GoogleVertexAiCustomJobJobSpecScheduling)
	PutWorkerPoolSpecs(value interface{})
	ResetBaseOutputDirectory()
	ResetEnableDashboardAccess()
	ResetEnableWebAccess()
	ResetNetwork()
	ResetPscInterfaceConfig()
	ResetReservedIpRanges()
	ResetScheduling()
	ResetServiceAccount()
	ResetTensorboard()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiCustomJobJobSpecOutputReference
type jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) BaseOutputDirectory() GoogleVertexAiCustomJobJobSpecBaseOutputDirectoryOutputReference {
	var returns GoogleVertexAiCustomJobJobSpecBaseOutputDirectoryOutputReference
	_jsii_.Get(
		j,
		"baseOutputDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) BaseOutputDirectoryInput() *GoogleVertexAiCustomJobJobSpecBaseOutputDirectory {
	var returns *GoogleVertexAiCustomJobJobSpecBaseOutputDirectory
	_jsii_.Get(
		j,
		"baseOutputDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) EnableDashboardAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDashboardAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) EnableDashboardAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDashboardAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) EnableWebAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWebAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) EnableWebAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWebAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) InternalValue() *GoogleVertexAiCustomJobJobSpec {
	var returns *GoogleVertexAiCustomJobJobSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) PscInterfaceConfig() GoogleVertexAiCustomJobJobSpecPscInterfaceConfigOutputReference {
	var returns GoogleVertexAiCustomJobJobSpecPscInterfaceConfigOutputReference
	_jsii_.Get(
		j,
		"pscInterfaceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) PscInterfaceConfigInput() *GoogleVertexAiCustomJobJobSpecPscInterfaceConfig {
	var returns *GoogleVertexAiCustomJobJobSpecPscInterfaceConfig
	_jsii_.Get(
		j,
		"pscInterfaceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ReservedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ReservedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reservedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) Scheduling() GoogleVertexAiCustomJobJobSpecSchedulingOutputReference {
	var returns GoogleVertexAiCustomJobJobSpecSchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) SchedulingInput() *GoogleVertexAiCustomJobJobSpecScheduling {
	var returns *GoogleVertexAiCustomJobJobSpecScheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) Tensorboard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tensorboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) TensorboardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tensorboardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) WorkerPoolSpecs() GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsList {
	var returns GoogleVertexAiCustomJobJobSpecWorkerPoolSpecsList
	_jsii_.Get(
		j,
		"workerPoolSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) WorkerPoolSpecsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workerPoolSpecsInput",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiCustomJobJobSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiCustomJobJobSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiCustomJobJobSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiCustomJob.GoogleVertexAiCustomJobJobSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiCustomJobJobSpecOutputReference_Override(g GoogleVertexAiCustomJobJobSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiCustomJob.GoogleVertexAiCustomJobJobSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference)SetEnableDashboardAccess(val interface{}) {
	if err := j.validateSetEnableDashboardAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDashboardAccess",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference)SetEnableWebAccess(val interface{}) {
	if err := j.validateSetEnableWebAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableWebAccess",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference)SetInternalValue(val *GoogleVertexAiCustomJobJobSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference)SetReservedIpRanges(val *[]*string) {
	if err := j.validateSetReservedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedIpRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference)SetTensorboard(val *string) {
	if err := j.validateSetTensorboardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tensorboard",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) PutBaseOutputDirectory(value *GoogleVertexAiCustomJobJobSpecBaseOutputDirectory) {
	if err := g.validatePutBaseOutputDirectoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBaseOutputDirectory",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) PutPscInterfaceConfig(value *GoogleVertexAiCustomJobJobSpecPscInterfaceConfig) {
	if err := g.validatePutPscInterfaceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPscInterfaceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) PutScheduling(value *GoogleVertexAiCustomJobJobSpecScheduling) {
	if err := g.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScheduling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) PutWorkerPoolSpecs(value interface{}) {
	if err := g.validatePutWorkerPoolSpecsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWorkerPoolSpecs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ResetBaseOutputDirectory() {
	_jsii_.InvokeVoid(
		g,
		"resetBaseOutputDirectory",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ResetEnableDashboardAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableDashboardAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ResetEnableWebAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableWebAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ResetPscInterfaceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPscInterfaceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ResetReservedIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetReservedIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ResetScheduling() {
	_jsii_.InvokeVoid(
		g,
		"resetScheduling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ResetTensorboard() {
	_jsii_.InvokeVoid(
		g,
		"resetTensorboard",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiCustomJobJobSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

