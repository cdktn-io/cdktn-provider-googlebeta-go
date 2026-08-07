// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlevertexaireasoningengine/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference interface {
	cdktn.ComplexObject
	AgentToAnywhereConfig() GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigAgentToAnywhereConfigOutputReference
	AgentToAnywhereConfigInput() *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigAgentToAnywhereConfig
	ClientToAgentConfig() GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigClientToAgentConfigOutputReference
	ClientToAgentConfigInput() *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigClientToAgentConfig
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
	InternalValue() *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfig
	SetInternalValue(val *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfig)
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
	PutAgentToAnywhereConfig(value *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigAgentToAnywhereConfig)
	PutClientToAgentConfig(value *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigClientToAgentConfig)
	ResetAgentToAnywhereConfig()
	ResetClientToAgentConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference
type jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) AgentToAnywhereConfig() GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigAgentToAnywhereConfigOutputReference {
	var returns GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigAgentToAnywhereConfigOutputReference
	_jsii_.Get(
		j,
		"agentToAnywhereConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) AgentToAnywhereConfigInput() *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigAgentToAnywhereConfig {
	var returns *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigAgentToAnywhereConfig
	_jsii_.Get(
		j,
		"agentToAnywhereConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) ClientToAgentConfig() GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigClientToAgentConfigOutputReference {
	var returns GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigClientToAgentConfigOutputReference
	_jsii_.Get(
		j,
		"clientToAgentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) ClientToAgentConfigInput() *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigClientToAgentConfig {
	var returns *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigClientToAgentConfig
	_jsii_.Get(
		j,
		"clientToAgentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) InternalValue() *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfig {
	var returns *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiReasoningEngine.GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference_Override(g GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiReasoningEngine.GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference)SetInternalValue(val *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) PutAgentToAnywhereConfig(value *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigAgentToAnywhereConfig) {
	if err := g.validatePutAgentToAnywhereConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAgentToAnywhereConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) PutClientToAgentConfig(value *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigClientToAgentConfig) {
	if err := g.validatePutClientToAgentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientToAgentConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) ResetAgentToAnywhereConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAgentToAnywhereConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) ResetClientToAgentConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetClientToAgentConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

