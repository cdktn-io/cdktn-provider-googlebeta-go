// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v21/googlecesdeployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleCesDeploymentWhatsappCredentialsOutputReference interface {
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
	BusinessAccountId() *string
	SetBusinessAccountId(val *string)
	BusinessAccountIdInput() *string
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
	InternalValue() *GoogleCesDeploymentWhatsappCredentials
	SetInternalValue(val *GoogleCesDeploymentWhatsappCredentials)
	PhoneNumber() *string
	SetPhoneNumber(val *string)
	PhoneNumberInput() *string
	Pin() *string
	SetPin(val *string)
	PinInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	PinWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetPinWo(val *string)
	PinWoInput() *string
	PinWoVersion() *string
	SetPinWoVersion(val *string)
	PinWoVersionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WabaId() *string
	SetWabaId(val *string)
	WabaIdInput() *string
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
	ResetPin()
	ResetPinWo()
	ResetPinWoVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesDeploymentWhatsappCredentialsOutputReference
type jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) AuthCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) AuthCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) AuthCodeWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) AuthCodeWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) AuthCodeWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) AuthCodeWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authCodeWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) BusinessAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) BusinessAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) ConversationProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conversationProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) ConversationProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conversationProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) InternalValue() *GoogleCesDeploymentWhatsappCredentials {
	var returns *GoogleCesDeploymentWhatsappCredentials
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) PhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) Pin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) PinInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) PinWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) PinWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) PinWoVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) PinWoVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pinWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) WabaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wabaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) WabaIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wabaIdInput",
		&returns,
	)
	return returns
}


func NewGoogleCesDeploymentWhatsappCredentialsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesDeploymentWhatsappCredentialsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesDeploymentWhatsappCredentialsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesDeployment.GoogleCesDeploymentWhatsappCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesDeploymentWhatsappCredentialsOutputReference_Override(g GoogleCesDeploymentWhatsappCredentialsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesDeployment.GoogleCesDeploymentWhatsappCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetAuthCode(val *string) {
	if err := j.validateSetAuthCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCode",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetAuthCodeWo(val *string) {
	if err := j.validateSetAuthCodeWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCodeWo",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetAuthCodeWoVersion(val *string) {
	if err := j.validateSetAuthCodeWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authCodeWoVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetBusinessAccountId(val *string) {
	if err := j.validateSetBusinessAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessAccountId",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetConversationProfileId(val *string) {
	if err := j.validateSetConversationProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conversationProfileId",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetInternalValue(val *GoogleCesDeploymentWhatsappCredentials) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetPhoneNumber(val *string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetPin(val *string) {
	if err := j.validateSetPinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pin",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetPinWo(val *string) {
	if err := j.validateSetPinWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pinWo",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetPinWoVersion(val *string) {
	if err := j.validateSetPinWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pinWoVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference)SetWabaId(val *string) {
	if err := j.validateSetWabaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wabaId",
		val,
	)
}

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) ResetAuthCode() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) ResetAuthCodeWo() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthCodeWo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) ResetAuthCodeWoVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthCodeWoVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) ResetConversationProfileId() {
	_jsii_.InvokeVoid(
		g,
		"resetConversationProfileId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) ResetPin() {
	_jsii_.InvokeVoid(
		g,
		"resetPin",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) ResetPinWo() {
	_jsii_.InvokeVoid(
		g,
		"resetPinWo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) ResetPinWoVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetPinWoVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesDeploymentWhatsappCredentialsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

