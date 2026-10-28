// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexairagfile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlevertexairagfile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference interface {
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
	CustomParsingPrompt() *string
	SetCustomParsingPrompt(val *string)
	CustomParsingPromptInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleVertexAiRagFileRagFileParsingConfigLlmParser
	SetInternalValue(val *GoogleVertexAiRagFileRagFileParsingConfigLlmParser)
	MaxParsingRequestsPerMin() *float64
	SetMaxParsingRequestsPerMin(val *float64)
	MaxParsingRequestsPerMinInput() *float64
	ModelName() *string
	SetModelName(val *string)
	ModelNameInput() *string
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
	ResetCustomParsingPrompt()
	ResetMaxParsingRequestsPerMin()
	ResetModelName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference
type jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) CustomParsingPrompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customParsingPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) CustomParsingPromptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customParsingPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) InternalValue() *GoogleVertexAiRagFileRagFileParsingConfigLlmParser {
	var returns *GoogleVertexAiRagFileRagFileParsingConfigLlmParser
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) MaxParsingRequestsPerMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParsingRequestsPerMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) MaxParsingRequestsPerMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParsingRequestsPerMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiRagFile.GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference_Override(g GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiRagFile.GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetCustomParsingPrompt(val *string) {
	if err := j.validateSetCustomParsingPromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customParsingPrompt",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetInternalValue(val *GoogleVertexAiRagFileRagFileParsingConfigLlmParser) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetMaxParsingRequestsPerMin(val *float64) {
	if err := j.validateSetMaxParsingRequestsPerMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParsingRequestsPerMin",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetModelName(val *string) {
	if err := j.validateSetModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelName",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ResetCustomParsingPrompt() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomParsingPrompt",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ResetMaxParsingRequestsPerMin() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxParsingRequestsPerMin",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ResetModelName() {
	_jsii_.InvokeVoid(
		g,
		"resetModelName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiRagFileRagFileParsingConfigLlmParserOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

