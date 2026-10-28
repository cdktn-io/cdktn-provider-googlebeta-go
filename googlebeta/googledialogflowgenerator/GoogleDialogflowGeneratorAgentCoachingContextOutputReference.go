// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledialogflowgenerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googledialogflowgenerator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDialogflowGeneratorAgentCoachingContextOutputReference interface {
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
	// Experimental.
	Fqn() *string
	Instructions() GoogleDialogflowGeneratorAgentCoachingContextInstructionsList
	InstructionsInput() interface{}
	InternalValue() *GoogleDialogflowGeneratorAgentCoachingContext
	SetInternalValue(val *GoogleDialogflowGeneratorAgentCoachingContext)
	OutputLanguageCode() *string
	SetOutputLanguageCode(val *string)
	OutputLanguageCodeInput() *string
	OverarchingGuidance() *string
	SetOverarchingGuidance(val *string)
	OverarchingGuidanceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutInstructions(value interface{})
	ResetInstructions()
	ResetOutputLanguageCode()
	ResetOverarchingGuidance()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowGeneratorAgentCoachingContextOutputReference
type jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) Instructions() GoogleDialogflowGeneratorAgentCoachingContextInstructionsList {
	var returns GoogleDialogflowGeneratorAgentCoachingContextInstructionsList
	_jsii_.Get(
		j,
		"instructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) InstructionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instructionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) InternalValue() *GoogleDialogflowGeneratorAgentCoachingContext {
	var returns *GoogleDialogflowGeneratorAgentCoachingContext
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) OutputLanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputLanguageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) OutputLanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputLanguageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) OverarchingGuidance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overarchingGuidance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) OverarchingGuidanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overarchingGuidanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowGeneratorAgentCoachingContextOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowGeneratorAgentCoachingContextOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowGeneratorAgentCoachingContextOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowGenerator.GoogleDialogflowGeneratorAgentCoachingContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowGeneratorAgentCoachingContextOutputReference_Override(g GoogleDialogflowGeneratorAgentCoachingContextOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowGenerator.GoogleDialogflowGeneratorAgentCoachingContextOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference)SetInternalValue(val *GoogleDialogflowGeneratorAgentCoachingContext) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference)SetOutputLanguageCode(val *string) {
	if err := j.validateSetOutputLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputLanguageCode",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference)SetOverarchingGuidance(val *string) {
	if err := j.validateSetOverarchingGuidanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overarchingGuidance",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) PutInstructions(value interface{}) {
	if err := g.validatePutInstructionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstructions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) ResetInstructions() {
	_jsii_.InvokeVoid(
		g,
		"resetInstructions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) ResetOutputLanguageCode() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputLanguageCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) ResetOverarchingGuidance() {
	_jsii_.InvokeVoid(
		g,
		"resetOverarchingGuidance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

