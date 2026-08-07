// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleoracledatabasegoldengateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference interface {
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
	InternalValue() *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties
	SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties)
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

// The jsii proxy struct for GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference
type jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) InternalValue() *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) PrivateKeyFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) PrivateKeyFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) PrivateKeyPassphraseSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPassphraseSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) PrivateKeyPassphraseSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPassphraseSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) PublicKeyFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) PublicKeyFingerprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKeyFingerprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) TechnologyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) TechnologyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) TenancyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) TenancyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) UseResourcePrincipal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourcePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) UseResourcePrincipalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourcePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) UserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference_Override(g GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetPrivateKeyFile(val *string) {
	if err := j.validateSetPrivateKeyFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKeyFile",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetPrivateKeyPassphraseSecret(val *string) {
	if err := j.validateSetPrivateKeyPassphraseSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKeyPassphraseSecret",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetPublicKeyFingerprint(val *string) {
	if err := j.validateSetPublicKeyFingerprintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicKeyFingerprint",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetTechnologyType(val *string) {
	if err := j.validateSetTechnologyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technologyType",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetTenancyId(val *string) {
	if err := j.validateSetTenancyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancyId",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetUseResourcePrincipal(val interface{}) {
	if err := j.validateSetUseResourcePrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useResourcePrincipal",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference)SetUserId(val *string) {
	if err := j.validateSetUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetPrivateKeyFile() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateKeyFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetPrivateKeyPassphraseSecret() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateKeyPassphraseSecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetPublicKeyFingerprint() {
	_jsii_.InvokeVoid(
		g,
		"resetPublicKeyFingerprint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetTechnologyType() {
	_jsii_.InvokeVoid(
		g,
		"resetTechnologyType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetTenancyId() {
	_jsii_.InvokeVoid(
		g,
		"resetTenancyId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetUseResourcePrincipal() {
	_jsii_.InvokeVoid(
		g,
		"resetUseResourcePrincipal",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ResetUserId() {
	_jsii_.InvokeVoid(
		g,
		"resetUserId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesOracleNosqlConnectionPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

