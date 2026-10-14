// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaionlineevaluator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googlevertexaionlineevaluator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference interface {
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
	Duration() GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterDurationOutputReference
	DurationInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterDuration
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TotalTokenUsage() GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterTotalTokenUsageOutputReference
	TotalTokenUsageInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterTotalTokenUsage
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
	PutDuration(value *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterDuration)
	PutTotalTokenUsage(value *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterTotalTokenUsage)
	ResetDuration()
	ResetTotalTokenUsage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference
type jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) Duration() GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterDurationOutputReference {
	var returns GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterDurationOutputReference
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) DurationInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterDuration {
	var returns *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterDuration
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) TotalTokenUsage() GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterTotalTokenUsageOutputReference {
	var returns GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterTotalTokenUsageOutputReference
	_jsii_.Get(
		j,
		"totalTokenUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) TotalTokenUsageInput() *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterTotalTokenUsage {
	var returns *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterTotalTokenUsage
	_jsii_.Get(
		j,
		"totalTokenUsageInput",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiOnlineEvaluator.GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference_Override(g GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiOnlineEvaluator.GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) PutDuration(value *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterDuration) {
	if err := g.validatePutDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDuration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) PutTotalTokenUsage(value *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterTotalTokenUsage) {
	if err := g.validatePutTotalTokenUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTotalTokenUsage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) ResetDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) ResetTotalTokenUsage() {
	_jsii_.InvokeVoid(
		g,
		"resetTotalTokenUsage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

