// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipersistentresource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlevertexaipersistentresource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
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
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList
type jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiPersistentResource.GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList_Override(g GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiPersistentResource.GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := g.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		g,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList) Get(index *float64) GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

