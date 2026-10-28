// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipipelinejob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlevertexaipipelinejob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiPipelineJobRuntimeConfigOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultRuntime() GoogleVertexAiPipelineJobRuntimeConfigDefaultRuntimeOutputReference
	DefaultRuntimeInput() *GoogleVertexAiPipelineJobRuntimeConfigDefaultRuntime
	FailurePolicy() *string
	SetFailurePolicy(val *string)
	FailurePolicyInput() *string
	// Experimental.
	Fqn() *string
	GcsOutputDirectory() *string
	SetGcsOutputDirectory(val *string)
	GcsOutputDirectoryInput() *string
	InputArtifacts() GoogleVertexAiPipelineJobRuntimeConfigInputArtifactsList
	InputArtifactsInput() interface{}
	InternalValue() *GoogleVertexAiPipelineJobRuntimeConfig
	SetInternalValue(val *GoogleVertexAiPipelineJobRuntimeConfig)
	ParameterValues() *map[string]*string
	SetParameterValues(val *map[string]*string)
	ParameterValuesInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutDefaultRuntime(value *GoogleVertexAiPipelineJobRuntimeConfigDefaultRuntime)
	PutInputArtifacts(value interface{})
	ResetDefaultRuntime()
	ResetFailurePolicy()
	ResetInputArtifacts()
	ResetParameterValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiPipelineJobRuntimeConfigOutputReference
type jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) DefaultRuntime() GoogleVertexAiPipelineJobRuntimeConfigDefaultRuntimeOutputReference {
	var returns GoogleVertexAiPipelineJobRuntimeConfigDefaultRuntimeOutputReference
	_jsii_.Get(
		j,
		"defaultRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) DefaultRuntimeInput() *GoogleVertexAiPipelineJobRuntimeConfigDefaultRuntime {
	var returns *GoogleVertexAiPipelineJobRuntimeConfigDefaultRuntime
	_jsii_.Get(
		j,
		"defaultRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) FailurePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) FailurePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failurePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) GcsOutputDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) GcsOutputDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) InputArtifacts() GoogleVertexAiPipelineJobRuntimeConfigInputArtifactsList {
	var returns GoogleVertexAiPipelineJobRuntimeConfigInputArtifactsList
	_jsii_.Get(
		j,
		"inputArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) InputArtifactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) InternalValue() *GoogleVertexAiPipelineJobRuntimeConfig {
	var returns *GoogleVertexAiPipelineJobRuntimeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) ParameterValues() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameterValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) ParameterValuesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameterValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiPipelineJobRuntimeConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiPipelineJobRuntimeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiPipelineJobRuntimeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiPipelineJob.GoogleVertexAiPipelineJobRuntimeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiPipelineJobRuntimeConfigOutputReference_Override(g GoogleVertexAiPipelineJobRuntimeConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiPipelineJob.GoogleVertexAiPipelineJobRuntimeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference)SetFailurePolicy(val *string) {
	if err := j.validateSetFailurePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failurePolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference)SetGcsOutputDirectory(val *string) {
	if err := j.validateSetGcsOutputDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsOutputDirectory",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference)SetInternalValue(val *GoogleVertexAiPipelineJobRuntimeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference)SetParameterValues(val *map[string]*string) {
	if err := j.validateSetParameterValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterValues",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) PutDefaultRuntime(value *GoogleVertexAiPipelineJobRuntimeConfigDefaultRuntime) {
	if err := g.validatePutDefaultRuntimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultRuntime",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) PutInputArtifacts(value interface{}) {
	if err := g.validatePutInputArtifactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInputArtifacts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) ResetDefaultRuntime() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultRuntime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) ResetFailurePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetFailurePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) ResetInputArtifacts() {
	_jsii_.InvokeVoid(
		g,
		"resetInputArtifacts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) ResetParameterValues() {
	_jsii_.InvokeVoid(
		g,
		"resetParameterValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiPipelineJobRuntimeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

