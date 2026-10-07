// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicesagentconnectivitytemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googlenetworkservicesagentconnectivitytemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference interface {
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
	DnsPeeringConfig() GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfigOutputReference
	DnsPeeringConfigInput() *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfig
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfig
	SetInternalValue(val *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfig)
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
	TlsConfig() GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfigOutputReference
	TlsConfigInput() *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfig
	VpcEgress() *string
	SetVpcEgress(val *string)
	VpcEgressInput() *string
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
	PutDnsPeeringConfig(value *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfig)
	PutTlsConfig(value *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfig)
	ResetDnsPeeringConfig()
	ResetNetworkAttachment()
	ResetTlsConfig()
	ResetVpcEgress()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference
type jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) DnsPeeringConfig() GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfigOutputReference {
	var returns GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfigOutputReference
	_jsii_.Get(
		j,
		"dnsPeeringConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) DnsPeeringConfigInput() *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfig {
	var returns *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfig
	_jsii_.Get(
		j,
		"dnsPeeringConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) InternalValue() *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfig {
	var returns *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) NetworkAttachment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) NetworkAttachmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) TlsConfig() GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfigOutputReference {
	var returns GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfigOutputReference
	_jsii_.Get(
		j,
		"tlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) TlsConfigInput() *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfig {
	var returns *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfig
	_jsii_.Get(
		j,
		"tlsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) VpcEgress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) VpcEgressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEgressInput",
		&returns,
	)
	return returns
}


func NewGoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkServicesAgentConnectivityTemplate.GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference_Override(g GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleNetworkServicesAgentConnectivityTemplate.GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference)SetInternalValue(val *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference)SetNetworkAttachment(val *string) {
	if err := j.validateSetNetworkAttachmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkAttachment",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference)SetVpcEgress(val *string) {
	if err := j.validateSetVpcEgressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcEgress",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) PutDnsPeeringConfig(value *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfig) {
	if err := g.validatePutDnsPeeringConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDnsPeeringConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) PutTlsConfig(value *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfig) {
	if err := g.validatePutTlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTlsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ResetDnsPeeringConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsPeeringConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ResetNetworkAttachment() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkAttachment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ResetTlsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetTlsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ResetVpcEgress() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcEgress",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

