// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googlecesagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference interface {
	cdktn.ComplexObject
	AgentCard() GoogleCesAgentRemoteA2AAgentA2AConfigAgentCardOutputReference
	AgentCardInput() *GoogleCesAgentRemoteA2AAgentA2AConfigAgentCard
	AgentRegistry() *string
	SetAgentRegistry(val *string)
	AgentRegistryInput() *string
	ApiAuthentication() GoogleCesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference
	ApiAuthenticationInput() *GoogleCesAgentRemoteA2AAgentA2AConfigApiAuthentication
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
	ContextId() *string
	SetContextId(val *string)
	ContextIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InputVariableMapping() *map[string]*string
	SetInputVariableMapping(val *map[string]*string)
	InputVariableMappingInput() *map[string]*string
	InternalValue() *GoogleCesAgentRemoteA2AAgentA2AConfig
	SetInternalValue(val *GoogleCesAgentRemoteA2AAgentA2AConfig)
	OutputVariableMapping() *map[string]*string
	SetOutputVariableMapping(val *map[string]*string)
	OutputVariableMappingInput() *map[string]*string
	StreamingEnabled() interface{}
	SetStreamingEnabled(val interface{})
	StreamingEnabledInput() interface{}
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
	PutAgentCard(value *GoogleCesAgentRemoteA2AAgentA2AConfigAgentCard)
	PutApiAuthentication(value *GoogleCesAgentRemoteA2AAgentA2AConfigApiAuthentication)
	ResetAgentCard()
	ResetAgentRegistry()
	ResetApiAuthentication()
	ResetContextId()
	ResetInputVariableMapping()
	ResetOutputVariableMapping()
	ResetStreamingEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference
type jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) AgentCard() GoogleCesAgentRemoteA2AAgentA2AConfigAgentCardOutputReference {
	var returns GoogleCesAgentRemoteA2AAgentA2AConfigAgentCardOutputReference
	_jsii_.Get(
		j,
		"agentCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) AgentCardInput() *GoogleCesAgentRemoteA2AAgentA2AConfigAgentCard {
	var returns *GoogleCesAgentRemoteA2AAgentA2AConfigAgentCard
	_jsii_.Get(
		j,
		"agentCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) AgentRegistry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) AgentRegistryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ApiAuthentication() GoogleCesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference {
	var returns GoogleCesAgentRemoteA2AAgentA2AConfigApiAuthenticationOutputReference
	_jsii_.Get(
		j,
		"apiAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ApiAuthenticationInput() *GoogleCesAgentRemoteA2AAgentA2AConfigApiAuthentication {
	var returns *GoogleCesAgentRemoteA2AAgentA2AConfigApiAuthentication
	_jsii_.Get(
		j,
		"apiAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ContextId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ContextIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) InputVariableMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"inputVariableMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) InputVariableMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"inputVariableMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) InternalValue() *GoogleCesAgentRemoteA2AAgentA2AConfig {
	var returns *GoogleCesAgentRemoteA2AAgentA2AConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) OutputVariableMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"outputVariableMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) OutputVariableMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"outputVariableMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) StreamingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) StreamingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesAgentRemoteA2AAgentA2AConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesAgentRemoteA2AAgentA2AConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesAgent.GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesAgentRemoteA2AAgentA2AConfigOutputReference_Override(g GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesAgent.GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference)SetAgentRegistry(val *string) {
	if err := j.validateSetAgentRegistryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentRegistry",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference)SetContextId(val *string) {
	if err := j.validateSetContextIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contextId",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference)SetInputVariableMapping(val *map[string]*string) {
	if err := j.validateSetInputVariableMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputVariableMapping",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference)SetInternalValue(val *GoogleCesAgentRemoteA2AAgentA2AConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference)SetOutputVariableMapping(val *map[string]*string) {
	if err := j.validateSetOutputVariableMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputVariableMapping",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference)SetStreamingEnabled(val interface{}) {
	if err := j.validateSetStreamingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamingEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) PutAgentCard(value *GoogleCesAgentRemoteA2AAgentA2AConfigAgentCard) {
	if err := g.validatePutAgentCardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAgentCard",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) PutApiAuthentication(value *GoogleCesAgentRemoteA2AAgentA2AConfigApiAuthentication) {
	if err := g.validatePutApiAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiAuthentication",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ResetAgentCard() {
	_jsii_.InvokeVoid(
		g,
		"resetAgentCard",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ResetAgentRegistry() {
	_jsii_.InvokeVoid(
		g,
		"resetAgentRegistry",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ResetApiAuthentication() {
	_jsii_.InvokeVoid(
		g,
		"resetApiAuthentication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ResetContextId() {
	_jsii_.InvokeVoid(
		g,
		"resetContextId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ResetInputVariableMapping() {
	_jsii_.InvokeVoid(
		g,
		"resetInputVariableMapping",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ResetOutputVariableMapping() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputVariableMapping",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ResetStreamingEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetStreamingEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesAgentRemoteA2AAgentA2AConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

