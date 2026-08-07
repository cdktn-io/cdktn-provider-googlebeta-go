// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledialogflowconversationprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googledialogflowconversationprofile/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDialogflowConversationProfileSipConfigOutputReference interface {
	cdktn.ComplexObject
	AllowVirtualAgentInteraction() interface{}
	SetAllowVirtualAgentInteraction(val interface{})
	AllowVirtualAgentInteractionInput() interface{}
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
	CopyInboundCallLegHeaders() *[]*string
	SetCopyInboundCallLegHeaders(val *[]*string)
	CopyInboundCallLegHeadersInput() *[]*string
	CreateConversationOnTheFly() interface{}
	SetCreateConversationOnTheFly(val interface{})
	CreateConversationOnTheFlyInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IgnoreReinviteMediaDirection() interface{}
	SetIgnoreReinviteMediaDirection(val interface{})
	IgnoreReinviteMediaDirectionInput() interface{}
	InactiveStart() interface{}
	SetInactiveStart(val interface{})
	InactiveStartInput() interface{}
	InternalValue() *GoogleDialogflowConversationProfileSipConfig
	SetInternalValue(val *GoogleDialogflowConversationProfileSipConfig)
	KeepConversationRunning() interface{}
	SetKeepConversationRunning(val interface{})
	KeepConversationRunningInput() interface{}
	MaxAudioRecordingDuration() *string
	SetMaxAudioRecordingDuration(val *string)
	MaxAudioRecordingDurationInput() *string
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
	ResetAllowVirtualAgentInteraction()
	ResetCopyInboundCallLegHeaders()
	ResetCreateConversationOnTheFly()
	ResetIgnoreReinviteMediaDirection()
	ResetInactiveStart()
	ResetKeepConversationRunning()
	ResetMaxAudioRecordingDuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowConversationProfileSipConfigOutputReference
type jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) AllowVirtualAgentInteraction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVirtualAgentInteraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) AllowVirtualAgentInteractionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVirtualAgentInteractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) CopyInboundCallLegHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"copyInboundCallLegHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) CopyInboundCallLegHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"copyInboundCallLegHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) CreateConversationOnTheFly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createConversationOnTheFly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) CreateConversationOnTheFlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createConversationOnTheFlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) IgnoreReinviteMediaDirection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreReinviteMediaDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) IgnoreReinviteMediaDirectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreReinviteMediaDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) InactiveStart() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inactiveStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) InactiveStartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inactiveStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) InternalValue() *GoogleDialogflowConversationProfileSipConfig {
	var returns *GoogleDialogflowConversationProfileSipConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) KeepConversationRunning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepConversationRunning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) KeepConversationRunningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepConversationRunningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) MaxAudioRecordingDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxAudioRecordingDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) MaxAudioRecordingDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxAudioRecordingDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowConversationProfileSipConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowConversationProfileSipConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowConversationProfileSipConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfileSipConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowConversationProfileSipConfigOutputReference_Override(g GoogleDialogflowConversationProfileSipConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfileSipConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference)SetAllowVirtualAgentInteraction(val interface{}) {
	if err := j.validateSetAllowVirtualAgentInteractionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowVirtualAgentInteraction",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference)SetCopyInboundCallLegHeaders(val *[]*string) {
	if err := j.validateSetCopyInboundCallLegHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyInboundCallLegHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference)SetCreateConversationOnTheFly(val interface{}) {
	if err := j.validateSetCreateConversationOnTheFlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createConversationOnTheFly",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference)SetIgnoreReinviteMediaDirection(val interface{}) {
	if err := j.validateSetIgnoreReinviteMediaDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreReinviteMediaDirection",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference)SetInactiveStart(val interface{}) {
	if err := j.validateSetInactiveStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inactiveStart",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference)SetInternalValue(val *GoogleDialogflowConversationProfileSipConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference)SetKeepConversationRunning(val interface{}) {
	if err := j.validateSetKeepConversationRunningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepConversationRunning",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference)SetMaxAudioRecordingDuration(val *string) {
	if err := j.validateSetMaxAudioRecordingDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAudioRecordingDuration",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) ResetAllowVirtualAgentInteraction() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowVirtualAgentInteraction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) ResetCopyInboundCallLegHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetCopyInboundCallLegHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) ResetCreateConversationOnTheFly() {
	_jsii_.InvokeVoid(
		g,
		"resetCreateConversationOnTheFly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) ResetIgnoreReinviteMediaDirection() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreReinviteMediaDirection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) ResetInactiveStart() {
	_jsii_.InvokeVoid(
		g,
		"resetInactiveStart",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) ResetKeepConversationRunning() {
	_jsii_.InvokeVoid(
		g,
		"resetKeepConversationRunning",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) ResetMaxAudioRecordingDuration() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxAudioRecordingDuration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfileSipConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

