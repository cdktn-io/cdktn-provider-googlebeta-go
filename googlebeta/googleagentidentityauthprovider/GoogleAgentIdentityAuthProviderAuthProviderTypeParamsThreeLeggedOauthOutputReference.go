// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagentidentityauthprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleagentidentityauthprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference interface {
	cdktn.ComplexObject
	AuthorizationUrl() *string
	SetAuthorizationUrl(val *string)
	AuthorizationUrlInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	ClientSecretWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetClientSecretWo(val *string)
	ClientSecretWoInput() *string
	ClientSecretWoVersion() *string
	SetClientSecretWoVersion(val *string)
	ClientSecretWoVersionInput() *string
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
	DefaultContinueUri() *string
	SetDefaultContinueUri(val *string)
	DefaultContinueUriInput() *string
	EnablePkce() interface{}
	SetEnablePkce(val interface{})
	EnablePkceInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth
	SetInternalValue(val *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth)
	RedirectUrl() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TokenUrl() *string
	SetTokenUrl(val *string)
	TokenUrlInput() *string
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
	ResetAuthorizationUrl()
	ResetClientId()
	ResetClientSecret()
	ResetClientSecretWo()
	ResetClientSecretWoVersion()
	ResetDefaultContinueUri()
	ResetEnablePkce()
	ResetTokenUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference
type jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) AuthorizationUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) AuthorizationUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientSecretWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientSecretWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientSecretWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ClientSecretWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) DefaultContinueUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultContinueUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) DefaultContinueUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultContinueUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) EnablePkce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePkce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) EnablePkceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePkceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) InternalValue() *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth {
	var returns *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) RedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) TokenUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrlInput",
		&returns,
	)
	return returns
}


func NewGoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAgentIdentityAuthProvider.GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference_Override(g GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAgentIdentityAuthProvider.GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetAuthorizationUrl(val *string) {
	if err := j.validateSetAuthorizationUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationUrl",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetClientSecretWo(val *string) {
	if err := j.validateSetClientSecretWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretWo",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetClientSecretWoVersion(val *string) {
	if err := j.validateSetClientSecretWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecretWoVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetDefaultContinueUri(val *string) {
	if err := j.validateSetDefaultContinueUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultContinueUri",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetEnablePkce(val interface{}) {
	if err := j.validateSetEnablePkceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePkce",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetInternalValue(val *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference)SetTokenUrl(val *string) {
	if err := j.validateSetTokenUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUrl",
		val,
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetAuthorizationUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorizationUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		g,
		"resetClientId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		g,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetClientSecretWo() {
	_jsii_.InvokeVoid(
		g,
		"resetClientSecretWo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetClientSecretWoVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetClientSecretWoVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetDefaultContinueUri() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultContinueUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetEnablePkce() {
	_jsii_.InvokeVoid(
		g,
		"resetEnablePkce",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ResetTokenUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetTokenUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

