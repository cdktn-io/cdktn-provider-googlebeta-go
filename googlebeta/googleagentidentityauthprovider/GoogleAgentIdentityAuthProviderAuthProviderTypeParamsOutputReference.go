// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagentidentityauthprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleagentidentityauthprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference interface {
	cdktn.ComplexObject
	ApiKey() GoogleAgentIdentityAuthProviderAuthProviderTypeParamsApiKeyOutputReference
	ApiKeyInput() *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsApiKey
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
	GeAuthProvider() GoogleAgentIdentityAuthProviderAuthProviderTypeParamsGeAuthProviderList
	InternalValue() *GoogleAgentIdentityAuthProviderAuthProviderTypeParams
	SetInternalValue(val *GoogleAgentIdentityAuthProviderAuthProviderTypeParams)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThreeLeggedOauth() GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference
	ThreeLeggedOauthInput() *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth
	TwoLeggedOauth() GoogleAgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference
	TwoLeggedOauthInput() *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth
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
	PutApiKey(value *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsApiKey)
	PutThreeLeggedOauth(value *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth)
	PutTwoLeggedOauth(value *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth)
	ResetApiKey()
	ResetThreeLeggedOauth()
	ResetTwoLeggedOauth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference
type jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ApiKey() GoogleAgentIdentityAuthProviderAuthProviderTypeParamsApiKeyOutputReference {
	var returns GoogleAgentIdentityAuthProviderAuthProviderTypeParamsApiKeyOutputReference
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ApiKeyInput() *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsApiKey {
	var returns *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsApiKey
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GeAuthProvider() GoogleAgentIdentityAuthProviderAuthProviderTypeParamsGeAuthProviderList {
	var returns GoogleAgentIdentityAuthProviderAuthProviderTypeParamsGeAuthProviderList
	_jsii_.Get(
		j,
		"geAuthProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) InternalValue() *GoogleAgentIdentityAuthProviderAuthProviderTypeParams {
	var returns *GoogleAgentIdentityAuthProviderAuthProviderTypeParams
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ThreeLeggedOauth() GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference {
	var returns GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference
	_jsii_.Get(
		j,
		"threeLeggedOauth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ThreeLeggedOauthInput() *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth {
	var returns *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth
	_jsii_.Get(
		j,
		"threeLeggedOauthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) TwoLeggedOauth() GoogleAgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference {
	var returns GoogleAgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauthOutputReference
	_jsii_.Get(
		j,
		"twoLeggedOauth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) TwoLeggedOauthInput() *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth {
	var returns *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth
	_jsii_.Get(
		j,
		"twoLeggedOauthInput",
		&returns,
	)
	return returns
}


func NewGoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAgentIdentityAuthProvider.GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference_Override(g GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAgentIdentityAuthProvider.GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference)SetInternalValue(val *GoogleAgentIdentityAuthProviderAuthProviderTypeParams) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) PutApiKey(value *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsApiKey) {
	if err := g.validatePutApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) PutThreeLeggedOauth(value *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth) {
	if err := g.validatePutThreeLeggedOauthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putThreeLeggedOauth",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) PutTwoLeggedOauth(value *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth) {
	if err := g.validatePutTwoLeggedOauthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTwoLeggedOauth",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ResetApiKey() {
	_jsii_.InvokeVoid(
		g,
		"resetApiKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ResetThreeLeggedOauth() {
	_jsii_.InvokeVoid(
		g,
		"resetThreeLeggedOauth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ResetTwoLeggedOauth() {
	_jsii_.InvokeVoid(
		g,
		"resetTwoLeggedOauth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

