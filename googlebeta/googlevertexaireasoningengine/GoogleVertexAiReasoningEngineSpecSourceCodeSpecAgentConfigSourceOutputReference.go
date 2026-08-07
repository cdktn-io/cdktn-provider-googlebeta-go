// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlevertexaireasoningengine/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference interface {
	cdktn.ComplexObject
	AdkConfig() GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceAdkConfigOutputReference
	AdkConfigInput() *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceAdkConfig
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
	// Experimental.
	Fqn() *string
	InlineSource() GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceInlineSourceOutputReference
	InlineSourceInput() *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceInlineSource
	InternalValue() *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSource
	SetInternalValue(val *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSource)
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
	PutAdkConfig(value *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceAdkConfig)
	PutInlineSource(value *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceInlineSource)
	ResetAdkConfig()
	ResetInlineSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference
type jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) AdkConfig() GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceAdkConfigOutputReference {
	var returns GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceAdkConfigOutputReference
	_jsii_.Get(
		j,
		"adkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) AdkConfigInput() *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceAdkConfig {
	var returns *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceAdkConfig
	_jsii_.Get(
		j,
		"adkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) InlineSource() GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceInlineSourceOutputReference {
	var returns GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceInlineSourceOutputReference
	_jsii_.Get(
		j,
		"inlineSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) InlineSourceInput() *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceInlineSource {
	var returns *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceInlineSource
	_jsii_.Get(
		j,
		"inlineSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) InternalValue() *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSource {
	var returns *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiReasoningEngine.GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference_Override(g GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiReasoningEngine.GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference)SetInternalValue(val *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) PutAdkConfig(value *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceAdkConfig) {
	if err := g.validatePutAdkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) PutInlineSource(value *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceInlineSource) {
	if err := g.validatePutInlineSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInlineSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) ResetAdkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAdkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) ResetInlineSource() {
	_jsii_.InvokeVoid(
		g,
		"resetInlineSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

