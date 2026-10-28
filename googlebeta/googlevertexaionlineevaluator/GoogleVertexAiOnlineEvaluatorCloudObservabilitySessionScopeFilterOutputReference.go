// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaionlineevaluator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googlevertexaionlineevaluator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference interface {
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
	Duration() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterDurationOutputReference
	DurationInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterDuration
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModelCallErrors() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCallErrorsOutputReference
	ModelCallErrorsInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCallErrors
	ModelCalls() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCallsOutputReference
	ModelCallsInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCalls
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ToolCallErrors() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCallErrorsOutputReference
	ToolCallErrorsInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCallErrors
	ToolCalls() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCallsOutputReference
	ToolCallsInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCalls
	TotalTokenUsage() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterTotalTokenUsageOutputReference
	TotalTokenUsageInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterTotalTokenUsage
	UserTurns() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterUserTurnsOutputReference
	UserTurnsInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterUserTurns
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
	PutDuration(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterDuration)
	PutModelCallErrors(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCallErrors)
	PutModelCalls(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCalls)
	PutToolCallErrors(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCallErrors)
	PutToolCalls(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCalls)
	PutTotalTokenUsage(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterTotalTokenUsage)
	PutUserTurns(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterUserTurns)
	ResetDuration()
	ResetModelCallErrors()
	ResetModelCalls()
	ResetToolCallErrors()
	ResetToolCalls()
	ResetTotalTokenUsage()
	ResetUserTurns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference
type jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) Duration() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterDurationOutputReference {
	var returns GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterDurationOutputReference
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) DurationInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterDuration {
	var returns *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterDuration
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ModelCallErrors() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCallErrorsOutputReference {
	var returns GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCallErrorsOutputReference
	_jsii_.Get(
		j,
		"modelCallErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ModelCallErrorsInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCallErrors {
	var returns *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCallErrors
	_jsii_.Get(
		j,
		"modelCallErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ModelCalls() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCallsOutputReference {
	var returns GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCallsOutputReference
	_jsii_.Get(
		j,
		"modelCalls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ModelCallsInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCalls {
	var returns *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCalls
	_jsii_.Get(
		j,
		"modelCallsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ToolCallErrors() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCallErrorsOutputReference {
	var returns GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCallErrorsOutputReference
	_jsii_.Get(
		j,
		"toolCallErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ToolCallErrorsInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCallErrors {
	var returns *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCallErrors
	_jsii_.Get(
		j,
		"toolCallErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ToolCalls() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCallsOutputReference {
	var returns GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCallsOutputReference
	_jsii_.Get(
		j,
		"toolCalls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ToolCallsInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCalls {
	var returns *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCalls
	_jsii_.Get(
		j,
		"toolCallsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) TotalTokenUsage() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterTotalTokenUsageOutputReference {
	var returns GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterTotalTokenUsageOutputReference
	_jsii_.Get(
		j,
		"totalTokenUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) TotalTokenUsageInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterTotalTokenUsage {
	var returns *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterTotalTokenUsage
	_jsii_.Get(
		j,
		"totalTokenUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) UserTurns() GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterUserTurnsOutputReference {
	var returns GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterUserTurnsOutputReference
	_jsii_.Get(
		j,
		"userTurns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) UserTurnsInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterUserTurns {
	var returns *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterUserTurns
	_jsii_.Get(
		j,
		"userTurnsInput",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiOnlineEvaluator.GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference_Override(g GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiOnlineEvaluator.GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) PutDuration(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterDuration) {
	if err := g.validatePutDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDuration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) PutModelCallErrors(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCallErrors) {
	if err := g.validatePutModelCallErrorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putModelCallErrors",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) PutModelCalls(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCalls) {
	if err := g.validatePutModelCallsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putModelCalls",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) PutToolCallErrors(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCallErrors) {
	if err := g.validatePutToolCallErrorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolCallErrors",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) PutToolCalls(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCalls) {
	if err := g.validatePutToolCallsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putToolCalls",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) PutTotalTokenUsage(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterTotalTokenUsage) {
	if err := g.validatePutTotalTokenUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTotalTokenUsage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) PutUserTurns(value *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterUserTurns) {
	if err := g.validatePutUserTurnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUserTurns",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ResetDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ResetModelCallErrors() {
	_jsii_.InvokeVoid(
		g,
		"resetModelCallErrors",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ResetModelCalls() {
	_jsii_.InvokeVoid(
		g,
		"resetModelCalls",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ResetToolCallErrors() {
	_jsii_.InvokeVoid(
		g,
		"resetToolCallErrors",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ResetToolCalls() {
	_jsii_.InvokeVoid(
		g,
		"resetToolCalls",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ResetTotalTokenUsage() {
	_jsii_.InvokeVoid(
		g,
		"resetTotalTokenUsage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ResetUserTurns() {
	_jsii_.InvokeVoid(
		g,
		"resetUserTurns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

