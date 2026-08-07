// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaipersistentresource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlevertexaipersistentresource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference interface {
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
	DnsPeeringConfigs() GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList
	DnsPeeringConfigsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleVertexAiPersistentResourcePscInterfaceConfig
	SetInternalValue(val *GoogleVertexAiPersistentResourcePscInterfaceConfig)
	NetworkAttachment() *string
	SetNetworkAttachment(val *string)
	NetworkAttachmentInput() *string
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
	PutDnsPeeringConfigs(value interface{})
	ResetDnsPeeringConfigs()
	ResetNetworkAttachment()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference
type jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) DnsPeeringConfigs() GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList {
	var returns GoogleVertexAiPersistentResourcePscInterfaceConfigDnsPeeringConfigsList
	_jsii_.Get(
		j,
		"dnsPeeringConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) DnsPeeringConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsPeeringConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) InternalValue() *GoogleVertexAiPersistentResourcePscInterfaceConfig {
	var returns *GoogleVertexAiPersistentResourcePscInterfaceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) NetworkAttachment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) NetworkAttachmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiPersistentResourcePscInterfaceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiPersistentResource.GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference_Override(g GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleVertexAiPersistentResource.GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference)SetInternalValue(val *GoogleVertexAiPersistentResourcePscInterfaceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference)SetNetworkAttachment(val *string) {
	if err := j.validateSetNetworkAttachmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkAttachment",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) PutDnsPeeringConfigs(value interface{}) {
	if err := g.validatePutDnsPeeringConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDnsPeeringConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) ResetDnsPeeringConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsPeeringConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) ResetNetworkAttachment() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkAttachment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiPersistentResourcePscInterfaceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

