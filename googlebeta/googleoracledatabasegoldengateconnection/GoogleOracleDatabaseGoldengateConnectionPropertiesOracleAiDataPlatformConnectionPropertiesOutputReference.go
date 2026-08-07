// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleoracledatabasegoldengateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference interface {
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
	ConnectionUrl() *string
	SetConnectionUrl(val *string)
	ConnectionUrlInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties
	SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties)
	PrivateKeyFile() *string
	SetPrivateKeyFile(val *string)
	PrivateKeyFileInput() *string
	PrivateKeyPassphraseSecret() *string
	SetPrivateKeyPassphraseSecret(val *string)
	PrivateKeyPassphraseSecretInput() *string
	PublicKeyFingerprint() *string
	SetPublicKeyFingerprint(val *string)
	PublicKeyFingerprintInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	TechnologyType() *string
	SetTechnologyType(val *string)
	TechnologyTypeInput() *string
	TenancyId() *string
	SetTenancyId(val *string)
	TenancyIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UseResourcePrincipal() interface{}
	SetUseResourcePrincipal(val interface{})
	UseResourcePrincipalInput() interface{}
	UserId() *string
	SetUserId(val *string)
	UserIdInput() *string
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
	ResetConnectionUrl()
	ResetPrivateKeyFile()
	ResetPrivateKeyPassphraseSecret()
	ResetPublicKeyFingerprint()
	ResetRegion()
	ResetTechnologyType()
	ResetTenancyId()
	ResetUseResourcePrincipal()
	ResetUserId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference
type jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ConnectionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ConnectionUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) InternalValue() *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) PrivateKeyFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) PrivateKeyFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) PrivateKeyPassphraseSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPassphraseSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) PrivateKeyPassphraseSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPassphraseSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) PublicKeyFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) PublicKeyFingerprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyFingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) TechnologyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) TechnologyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) TenancyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) TenancyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) UseResourcePrincipal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourcePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) UseResourcePrincipalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourcePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) UserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference_Override(g GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference)SetConnectionUrl(val *string) {
	if err := j.validateSetConnectionUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionUrl",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference)SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference)SetPrivateKeyFile(val *string) {
	if err := j.validateSetPrivateKeyFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKeyFile",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference)SetPrivateKeyPassphraseSecret(val *string) {
	if err := j.validateSetPrivateKeyPassphraseSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKeyPassphraseSecret",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference)SetPublicKeyFingerprint(val *string) {
	if err := j.validateSetPublicKeyFingerprintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicKeyFingerprint",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference)SetTechnologyType(val *string) {
	if err := j.validateSetTechnologyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technologyType",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference)SetTenancyId(val *string) {
	if err := j.validateSetTenancyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancyId",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference)SetUseResourcePrincipal(val interface{}) {
	if err := j.validateSetUseResourcePrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useResourcePrincipal",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference)SetUserId(val *string) {
	if err := j.validateSetUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ResetConnectionUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ResetPrivateKeyFile() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateKeyFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ResetPrivateKeyPassphraseSecret() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateKeyPassphraseSecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ResetPublicKeyFingerprint() {
	_jsii_.InvokeVoid(
		g,
		"resetPublicKeyFingerprint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ResetTechnologyType() {
	_jsii_.InvokeVoid(
		g,
		"resetTechnologyType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ResetTenancyId() {
	_jsii_.InvokeVoid(
		g,
		"resetTenancyId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ResetUseResourcePrincipal() {
	_jsii_.InvokeVoid(
		g,
		"resetUseResourcePrincipal",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ResetUserId() {
	_jsii_.InvokeVoid(
		g,
		"resetUserId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

