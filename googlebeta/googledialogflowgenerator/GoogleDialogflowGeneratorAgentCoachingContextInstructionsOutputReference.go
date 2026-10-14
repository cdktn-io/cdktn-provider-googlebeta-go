// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledialogflowgenerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googledialogflowgenerator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference interface {
	cdktn.ComplexObject
	AgentAction() *string
	SetAgentAction(val *string)
	AgentActionInput() *string
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
	Condition() *string
	SetCondition(val *string)
	ConditionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisplayDetails() *string
	SetDisplayDetails(val *string)
	DisplayDetailsInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SystemAction() *string
	SetSystemAction(val *string)
	SystemActionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TriggeringEvent() *string
	SetTriggeringEvent(val *string)
	TriggeringEventInput() *string
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
	ResetAgentAction()
	ResetCondition()
	ResetDisplayDetails()
	ResetDisplayName()
	ResetSystemAction()
	ResetTriggeringEvent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference
type jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) AgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) AgentActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) Condition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) DisplayDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) DisplayDetailsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) SystemAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) SystemActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) TriggeringEvent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggeringEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) TriggeringEventInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggeringEventInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowGenerator.GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference_Override(g GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowGenerator.GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetAgentAction(val *string) {
	if err := j.validateSetAgentActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentAction",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetCondition(val *string) {
	if err := j.validateSetConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"condition",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetDisplayDetails(val *string) {
	if err := j.validateSetDisplayDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayDetails",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetSystemAction(val *string) {
	if err := j.validateSetSystemActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemAction",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference)SetTriggeringEvent(val *string) {
	if err := j.validateSetTriggeringEventParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggeringEvent",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ResetAgentAction() {
	_jsii_.InvokeVoid(
		g,
		"resetAgentAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		g,
		"resetCondition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ResetDisplayDetails() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayDetails",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ResetSystemAction() {
	_jsii_.InvokeVoid(
		g,
		"resetSystemAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ResetTriggeringEvent() {
	_jsii_.InvokeVoid(
		g,
		"resetTriggeringEvent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowGeneratorAgentCoachingContextInstructionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

