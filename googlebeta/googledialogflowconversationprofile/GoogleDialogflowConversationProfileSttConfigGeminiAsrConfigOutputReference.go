// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledialogflowconversationprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v22/googledialogflowconversationprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference interface {
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
	EndOfSpeechSensitivity() *string
	SetEndOfSpeechSensitivity(val *string)
	EndOfSpeechSensitivityInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDialogflowConversationProfileSttConfigGeminiAsrConfig
	SetInternalValue(val *GoogleDialogflowConversationProfileSttConfigGeminiAsrConfig)
	ModelId() *string
	SetModelId(val *string)
	ModelIdInput() *string
	PrefixPaddingMs() *float64
	SetPrefixPaddingMs(val *float64)
	PrefixPaddingMsInput() *float64
	SilenceDurationMs() *float64
	SetSilenceDurationMs(val *float64)
	SilenceDurationMsInput() *float64
	StartOfSpeechSensitivity() *string
	SetStartOfSpeechSensitivity(val *string)
	StartOfSpeechSensitivityInput() *string
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
	ResetEndOfSpeechSensitivity()
	ResetModelId()
	ResetPrefixPaddingMs()
	ResetSilenceDurationMs()
	ResetStartOfSpeechSensitivity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference
type jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) EndOfSpeechSensitivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endOfSpeechSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) EndOfSpeechSensitivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endOfSpeechSensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) InternalValue() *GoogleDialogflowConversationProfileSttConfigGeminiAsrConfig {
	var returns *GoogleDialogflowConversationProfileSttConfigGeminiAsrConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ModelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) PrefixPaddingMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"prefixPaddingMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) PrefixPaddingMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"prefixPaddingMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) SilenceDurationMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"silenceDurationMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) SilenceDurationMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"silenceDurationMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) StartOfSpeechSensitivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startOfSpeechSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) StartOfSpeechSensitivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startOfSpeechSensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference_Override(g GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetEndOfSpeechSensitivity(val *string) {
	if err := j.validateSetEndOfSpeechSensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endOfSpeechSensitivity",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetInternalValue(val *GoogleDialogflowConversationProfileSttConfigGeminiAsrConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetModelId(val *string) {
	if err := j.validateSetModelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelId",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetPrefixPaddingMs(val *float64) {
	if err := j.validateSetPrefixPaddingMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixPaddingMs",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetSilenceDurationMs(val *float64) {
	if err := j.validateSetSilenceDurationMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"silenceDurationMs",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetStartOfSpeechSensitivity(val *string) {
	if err := j.validateSetStartOfSpeechSensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startOfSpeechSensitivity",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ResetEndOfSpeechSensitivity() {
	_jsii_.InvokeVoid(
		g,
		"resetEndOfSpeechSensitivity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ResetModelId() {
	_jsii_.InvokeVoid(
		g,
		"resetModelId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ResetPrefixPaddingMs() {
	_jsii_.InvokeVoid(
		g,
		"resetPrefixPaddingMs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ResetSilenceDurationMs() {
	_jsii_.InvokeVoid(
		g,
		"resetSilenceDurationMs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ResetStartOfSpeechSensitivity() {
	_jsii_.InvokeVoid(
		g,
		"resetStartOfSpeechSensitivity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSttConfigGeminiAsrConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

