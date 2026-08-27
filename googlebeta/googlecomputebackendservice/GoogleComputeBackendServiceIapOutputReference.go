// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputebackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googlecomputebackendservice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleComputeBackendServiceIapOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeBackendServiceIap
	SetInternalValue(val *GoogleComputeBackendServiceIap)
	Oauth2ClientId() *string
	SetOauth2ClientId(val *string)
	Oauth2ClientIdInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	Oauth2ClientIdWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetOauth2ClientIdWo(val *string)
	Oauth2ClientIdWoInput() *string
	Oauth2ClientIdWoVersion() *string
	SetOauth2ClientIdWoVersion(val *string)
	Oauth2ClientIdWoVersionInput() *string
	Oauth2ClientSecret() *string
	SetOauth2ClientSecret(val *string)
	Oauth2ClientSecretInput() *string
	Oauth2ClientSecretSha256() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	Oauth2ClientSecretWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetOauth2ClientSecretWo(val *string)
	Oauth2ClientSecretWoInput() *string
	Oauth2ClientSecretWoVersion() *string
	SetOauth2ClientSecretWoVersion(val *string)
	Oauth2ClientSecretWoVersionInput() *string
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
	ResetOauth2ClientId()
	ResetOauth2ClientIdWo()
	ResetOauth2ClientIdWoVersion()
	ResetOauth2ClientSecret()
	ResetOauth2ClientSecretWo()
	ResetOauth2ClientSecretWoVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeBackendServiceIapOutputReference
type jsiiProxy_GoogleComputeBackendServiceIapOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) InternalValue() *GoogleComputeBackendServiceIap {
	var returns *GoogleComputeBackendServiceIap
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Oauth2ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2ClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Oauth2ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2ClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Oauth2ClientIdWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2ClientIdWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Oauth2ClientIdWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2ClientIdWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Oauth2ClientIdWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2ClientIdWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Oauth2ClientIdWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2ClientIdWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Oauth2ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2ClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Oauth2ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2ClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Oauth2ClientSecretSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2ClientSecretSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Oauth2ClientSecretWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2ClientSecretWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Oauth2ClientSecretWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2ClientSecretWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Oauth2ClientSecretWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2ClientSecretWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Oauth2ClientSecretWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2ClientSecretWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeBackendServiceIapOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleComputeBackendServiceIapOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeBackendServiceIapOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeBackendServiceIapOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeBackendService.GoogleComputeBackendServiceIapOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeBackendServiceIapOutputReference_Override(g GoogleComputeBackendServiceIapOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleComputeBackendService.GoogleComputeBackendServiceIapOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference)SetInternalValue(val *GoogleComputeBackendServiceIap) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference)SetOauth2ClientId(val *string) {
	if err := j.validateSetOauth2ClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauth2ClientId",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference)SetOauth2ClientIdWo(val *string) {
	if err := j.validateSetOauth2ClientIdWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauth2ClientIdWo",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference)SetOauth2ClientIdWoVersion(val *string) {
	if err := j.validateSetOauth2ClientIdWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauth2ClientIdWoVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference)SetOauth2ClientSecret(val *string) {
	if err := j.validateSetOauth2ClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauth2ClientSecret",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference)SetOauth2ClientSecretWo(val *string) {
	if err := j.validateSetOauth2ClientSecretWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauth2ClientSecretWo",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference)SetOauth2ClientSecretWoVersion(val *string) {
	if err := j.validateSetOauth2ClientSecretWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauth2ClientSecretWoVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendServiceIapOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) ResetOauth2ClientId() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2ClientId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) ResetOauth2ClientIdWo() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2ClientIdWo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) ResetOauth2ClientIdWoVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2ClientIdWoVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) ResetOauth2ClientSecret() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2ClientSecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) ResetOauth2ClientSecretWo() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2ClientSecretWo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) ResetOauth2ClientSecretWoVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetOauth2ClientSecretWoVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeBackendServiceIapOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

