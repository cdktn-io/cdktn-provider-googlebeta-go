// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleoracledatabasegoldengateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference interface {
	cdktn.ComplexObject
	AdditionalAttributes() GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesAdditionalAttributesList
	AdditionalAttributesInput() interface{}
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
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	// Experimental.
	Fqn() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() *GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties
	SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PasswordSecretVersion() *string
	SetPasswordSecretVersion(val *string)
	PasswordSecretVersionInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	SecurityProtocol() *string
	SetSecurityProtocol(val *string)
	SecurityProtocolInput() *string
	SslClientKeystashFile() *string
	SetSslClientKeystashFile(val *string)
	SslClientKeystashFileInput() *string
	SslClientKeystoredbFile() *string
	SetSslClientKeystoredbFile(val *string)
	SslClientKeystoredbFileInput() *string
	SslServerCertificateFile() *string
	SetSslServerCertificateFile(val *string)
	SslServerCertificateFileInput() *string
	TechnologyType() *string
	SetTechnologyType(val *string)
	TechnologyTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	PutAdditionalAttributes(value interface{})
	ResetAdditionalAttributes()
	ResetDatabase()
	ResetHost()
	ResetPassword()
	ResetPasswordSecretVersion()
	ResetPort()
	ResetSecurityProtocol()
	ResetSslClientKeystashFile()
	ResetSslClientKeystoredbFile()
	ResetSslServerCertificateFile()
	ResetTechnologyType()
	ResetUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference
type jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) AdditionalAttributes() GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesAdditionalAttributesList {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesAdditionalAttributesList
	_jsii_.Get(
		j,
		"additionalAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) AdditionalAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) InternalValue() *GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) PasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) PasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) SecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) SslClientKeystashFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeystashFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) SslClientKeystashFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeystashFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) SslClientKeystoredbFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeystoredbFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) SslClientKeystoredbFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeystoredbFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) SslServerCertificateFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslServerCertificateFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) SslServerCertificateFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslServerCertificateFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) TechnologyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) TechnologyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference_Override(g GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetPasswordSecretVersion(val *string) {
	if err := j.validateSetPasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetSecurityProtocol(val *string) {
	if err := j.validateSetSecurityProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProtocol",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetSslClientKeystashFile(val *string) {
	if err := j.validateSetSslClientKeystashFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslClientKeystashFile",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetSslClientKeystoredbFile(val *string) {
	if err := j.validateSetSslClientKeystoredbFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslClientKeystoredbFile",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetSslServerCertificateFile(val *string) {
	if err := j.validateSetSslServerCertificateFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslServerCertificateFile",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetTechnologyType(val *string) {
	if err := j.validateSetTechnologyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technologyType",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) PutAdditionalAttributes(value interface{}) {
	if err := g.validatePutAdditionalAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdditionalAttributes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ResetAdditionalAttributes() {
	_jsii_.InvokeVoid(
		g,
		"resetAdditionalAttributes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		g,
		"resetHost",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ResetPasswordSecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetPasswordSecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		g,
		"resetPort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ResetSecurityProtocol() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityProtocol",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ResetSslClientKeystashFile() {
	_jsii_.InvokeVoid(
		g,
		"resetSslClientKeystashFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ResetSslClientKeystoredbFile() {
	_jsii_.InvokeVoid(
		g,
		"resetSslClientKeystoredbFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ResetSslServerCertificateFile() {
	_jsii_.InvokeVoid(
		g,
		"resetSslServerCertificateFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ResetTechnologyType() {
	_jsii_.InvokeVoid(
		g,
		"resetTechnologyType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		g,
		"resetUsername",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

