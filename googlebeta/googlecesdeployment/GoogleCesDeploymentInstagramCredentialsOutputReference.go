// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googlecesdeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleCesDeploymentInstagramCredentialsOutputReference interface {
	cdktn.ComplexObject
	AuthCode() *string
	SetAuthCode(val *string)
	AuthCodeInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	AuthCodeWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetAuthCodeWo(val *string)
	AuthCodeWoInput() *string
	AuthCodeWoVersion() *string
	SetAuthCodeWoVersion(val *string)
	AuthCodeWoVersionInput() *string
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
	ConversationProfileId() *string
	SetConversationProfileId(val *string)
	ConversationProfileIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleCesDeploymentInstagramCredentials
	SetInternalValue(val *GoogleCesDeploymentInstagramCredentials)
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
	ResetAuthCode()
	ResetAuthCodeWo()
	ResetAuthCodeWoVersion()
	ResetConversationProfileId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesDeploymentInstagramCredentialsOutputReference
type jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) AuthCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) AuthCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) AuthCodeWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) AuthCodeWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) AuthCodeWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) AuthCodeWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) ConversationProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conversationProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) ConversationProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conversationProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) InternalValue() *GoogleCesDeploymentInstagramCredentials {
	var returns *GoogleCesDeploymentInstagramCredentials
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesDeploymentInstagramCredentialsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesDeploymentInstagramCredentialsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesDeploymentInstagramCredentialsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesDeployment.GoogleCesDeploymentInstagramCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesDeploymentInstagramCredentialsOutputReference_Override(g GoogleCesDeploymentInstagramCredentialsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesDeployment.GoogleCesDeploymentInstagramCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference)SetAuthCode(val *string) {
	if err := j.validateSetAuthCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCode",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference)SetAuthCodeWo(val *string) {
	if err := j.validateSetAuthCodeWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCodeWo",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference)SetAuthCodeWoVersion(val *string) {
	if err := j.validateSetAuthCodeWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCodeWoVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference)SetConversationProfileId(val *string) {
	if err := j.validateSetConversationProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conversationProfileId",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference)SetInternalValue(val *GoogleCesDeploymentInstagramCredentials) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) ResetAuthCode() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) ResetAuthCodeWo() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthCodeWo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) ResetAuthCodeWoVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthCodeWoVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) ResetConversationProfileId() {
	_jsii_.InvokeVoid(
		g,
		"resetConversationProfileId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesDeploymentInstagramCredentialsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

