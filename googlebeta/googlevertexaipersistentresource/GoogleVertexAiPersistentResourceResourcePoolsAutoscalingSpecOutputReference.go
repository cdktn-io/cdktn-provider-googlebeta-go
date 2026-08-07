// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipersistentresource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlevertexaipersistentresource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference interface {
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
	InternalValue() *GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpec
	SetInternalValue(val *GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpec)
	MaxReplicaCount() *string
	SetMaxReplicaCount(val *string)
	MaxReplicaCountInput() *string
	MinReplicaCount() *string
	SetMinReplicaCount(val *string)
	MinReplicaCountInput() *string
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
	ResetMaxReplicaCount()
	ResetMinReplicaCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference
type jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) InternalValue() *GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpec {
	var returns *GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) MaxReplicaCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxReplicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) MaxReplicaCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxReplicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) MinReplicaCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minReplicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) MinReplicaCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minReplicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiPersistentResource.GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference_Override(g GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiPersistentResource.GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference)SetInternalValue(val *GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference)SetMaxReplicaCount(val *string) {
	if err := j.validateSetMaxReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxReplicaCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference)SetMinReplicaCount(val *string) {
	if err := j.validateSetMinReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReplicaCount",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) ResetMaxReplicaCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxReplicaCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) ResetMinReplicaCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMinReplicaCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourceResourcePoolsAutoscalingSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

