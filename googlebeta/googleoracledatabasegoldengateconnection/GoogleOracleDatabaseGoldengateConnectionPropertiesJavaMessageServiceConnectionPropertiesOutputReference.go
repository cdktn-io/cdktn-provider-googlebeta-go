// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleoracledatabasegoldengateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference interface {
	cdktn.ComplexObject
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
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
	ConnectionFactory() *string
	SetConnectionFactory(val *string)
	ConnectionFactoryInput() *string
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
	InternalValue() *GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties
	SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties)
	JndiConnectionFactory() *string
	SetJndiConnectionFactory(val *string)
	JndiConnectionFactoryInput() *string
	JndiInitialContextFactory() *string
	SetJndiInitialContextFactory(val *string)
	JndiInitialContextFactoryInput() *string
	JndiProviderUrl() *string
	SetJndiProviderUrl(val *string)
	JndiProviderUrlInput() *string
	JndiSecurityCredentialsSecret() *string
	SetJndiSecurityCredentialsSecret(val *string)
	JndiSecurityCredentialsSecretInput() *string
	JndiSecurityPrincipal() *string
	SetJndiSecurityPrincipal(val *string)
	JndiSecurityPrincipalInput() *string
	KeyStoreFile() *string
	SetKeyStoreFile(val *string)
	KeyStoreFileInput() *string
	KeyStorePassword() *string
	SetKeyStorePassword(val *string)
	KeyStorePasswordInput() *string
	KeyStorePasswordSecretVersion() *string
	SetKeyStorePasswordSecretVersion(val *string)
	KeyStorePasswordSecretVersionInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PasswordSecretVersion() *string
	SetPasswordSecretVersion(val *string)
	PasswordSecretVersionInput() *string
	SecurityProtocol() *string
	SetSecurityProtocol(val *string)
	SecurityProtocolInput() *string
	SslKeyPassword() *string
	SetSslKeyPassword(val *string)
	SslKeyPasswordInput() *string
	SslKeyPasswordSecretVersion() *string
	SetSslKeyPasswordSecretVersion(val *string)
	SslKeyPasswordSecretVersionInput() *string
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
	TrustStoreFile() *string
	SetTrustStoreFile(val *string)
	TrustStoreFileInput() *string
	TrustStorePassword() *string
	SetTrustStorePassword(val *string)
	TrustStorePasswordInput() *string
	TrustStorePasswordSecretVersion() *string
	SetTrustStorePasswordSecretVersion(val *string)
	TrustStorePasswordSecretVersionInput() *string
	UseJndi() interface{}
	SetUseJndi(val interface{})
	UseJndiInput() interface{}
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
	ResetAuthenticationType()
	ResetConnectionFactory()
	ResetConnectionUrl()
	ResetJndiConnectionFactory()
	ResetJndiInitialContextFactory()
	ResetJndiProviderUrl()
	ResetJndiSecurityCredentialsSecret()
	ResetJndiSecurityPrincipal()
	ResetKeyStoreFile()
	ResetKeyStorePassword()
	ResetKeyStorePasswordSecretVersion()
	ResetPassword()
	ResetPasswordSecretVersion()
	ResetSecurityProtocol()
	ResetSslKeyPassword()
	ResetSslKeyPasswordSecretVersion()
	ResetTechnologyType()
	ResetTrustStoreFile()
	ResetTrustStorePassword()
	ResetTrustStorePasswordSecretVersion()
	ResetUseJndi()
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

// The jsii proxy struct for GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference
type jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ConnectionFactory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ConnectionFactoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionFactoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ConnectionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ConnectionUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) InternalValue() *GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) JndiConnectionFactory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jndiConnectionFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) JndiConnectionFactoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jndiConnectionFactoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) JndiInitialContextFactory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jndiInitialContextFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) JndiInitialContextFactoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jndiInitialContextFactoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) JndiProviderUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jndiProviderUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) JndiProviderUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jndiProviderUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) JndiSecurityCredentialsSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jndiSecurityCredentialsSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) JndiSecurityCredentialsSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jndiSecurityCredentialsSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) JndiSecurityPrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jndiSecurityPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) JndiSecurityPrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jndiSecurityPrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) KeyStoreFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStoreFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) KeyStoreFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStoreFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) KeyStorePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) KeyStorePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) KeyStorePasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePasswordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) KeyStorePasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePasswordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) PasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) PasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) SecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) SslKeyPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) SslKeyPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) SslKeyPasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyPasswordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) SslKeyPasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyPasswordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) TechnologyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) TechnologyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) TrustStoreFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) TrustStoreFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) TrustStorePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStorePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) TrustStorePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStorePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) TrustStorePasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStorePasswordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) TrustStorePasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStorePasswordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) UseJndi() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useJndi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) UseJndiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useJndiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference_Override(g GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetConnectionFactory(val *string) {
	if err := j.validateSetConnectionFactoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionFactory",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetConnectionUrl(val *string) {
	if err := j.validateSetConnectionUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionUrl",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetJndiConnectionFactory(val *string) {
	if err := j.validateSetJndiConnectionFactoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jndiConnectionFactory",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetJndiInitialContextFactory(val *string) {
	if err := j.validateSetJndiInitialContextFactoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jndiInitialContextFactory",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetJndiProviderUrl(val *string) {
	if err := j.validateSetJndiProviderUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jndiProviderUrl",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetJndiSecurityCredentialsSecret(val *string) {
	if err := j.validateSetJndiSecurityCredentialsSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jndiSecurityCredentialsSecret",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetJndiSecurityPrincipal(val *string) {
	if err := j.validateSetJndiSecurityPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jndiSecurityPrincipal",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetKeyStoreFile(val *string) {
	if err := j.validateSetKeyStoreFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyStoreFile",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetKeyStorePassword(val *string) {
	if err := j.validateSetKeyStorePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyStorePassword",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetKeyStorePasswordSecretVersion(val *string) {
	if err := j.validateSetKeyStorePasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyStorePasswordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetPasswordSecretVersion(val *string) {
	if err := j.validateSetPasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetSecurityProtocol(val *string) {
	if err := j.validateSetSecurityProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProtocol",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetSslKeyPassword(val *string) {
	if err := j.validateSetSslKeyPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslKeyPassword",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetSslKeyPasswordSecretVersion(val *string) {
	if err := j.validateSetSslKeyPasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslKeyPasswordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetTechnologyType(val *string) {
	if err := j.validateSetTechnologyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technologyType",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetTrustStoreFile(val *string) {
	if err := j.validateSetTrustStoreFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustStoreFile",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetTrustStorePassword(val *string) {
	if err := j.validateSetTrustStorePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustStorePassword",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetTrustStorePasswordSecretVersion(val *string) {
	if err := j.validateSetTrustStorePasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustStorePasswordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetUseJndi(val interface{}) {
	if err := j.validateSetUseJndiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useJndi",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetAuthenticationType() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthenticationType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetConnectionFactory() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionFactory",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetConnectionUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetJndiConnectionFactory() {
	_jsii_.InvokeVoid(
		g,
		"resetJndiConnectionFactory",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetJndiInitialContextFactory() {
	_jsii_.InvokeVoid(
		g,
		"resetJndiInitialContextFactory",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetJndiProviderUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetJndiProviderUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetJndiSecurityCredentialsSecret() {
	_jsii_.InvokeVoid(
		g,
		"resetJndiSecurityCredentialsSecret",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetJndiSecurityPrincipal() {
	_jsii_.InvokeVoid(
		g,
		"resetJndiSecurityPrincipal",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetKeyStoreFile() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyStoreFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetKeyStorePassword() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyStorePassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetKeyStorePasswordSecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyStorePasswordSecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetPasswordSecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetPasswordSecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetSecurityProtocol() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityProtocol",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetSslKeyPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetSslKeyPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetSslKeyPasswordSecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetSslKeyPasswordSecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetTechnologyType() {
	_jsii_.InvokeVoid(
		g,
		"resetTechnologyType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetTrustStoreFile() {
	_jsii_.InvokeVoid(
		g,
		"resetTrustStoreFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetTrustStorePassword() {
	_jsii_.InvokeVoid(
		g,
		"resetTrustStorePassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetTrustStorePasswordSecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetTrustStorePasswordSecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetUseJndi() {
	_jsii_.InvokeVoid(
		g,
		"resetUseJndi",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		g,
		"resetUsername",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

