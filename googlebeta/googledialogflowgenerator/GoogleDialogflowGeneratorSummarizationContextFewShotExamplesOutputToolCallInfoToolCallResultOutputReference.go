// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledialogflowgenerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googledialogflowgenerator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference interface {
	cdktn.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	Error() GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultErrorOutputReference
	ErrorInput() *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultError
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResult
	SetInternalValue(val *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResult)
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
	PutError(value *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultError)
	ResetAction()
	ResetError()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference
type jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) Error() GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultErrorOutputReference {
	var returns GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultErrorOutputReference
	_jsii_.Get(
		j,
		"error",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) ErrorInput() *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultError {
	var returns *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultError
	_jsii_.Get(
		j,
		"errorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) InternalValue() *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResult {
	var returns *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResult
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowGenerator.GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference_Override(g GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowGenerator.GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference)SetInternalValue(val *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResult) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) PutError(value *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultError) {
	if err := g.validatePutErrorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putError",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		g,
		"resetAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) ResetError() {
	_jsii_.InvokeVoid(
		g,
		"resetError",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

