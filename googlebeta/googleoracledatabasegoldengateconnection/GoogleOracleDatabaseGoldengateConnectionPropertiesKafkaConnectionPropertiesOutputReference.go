// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/jsii"

	"github.com/cdktn-io/cdktn-provider-googlebeta-go/googlebeta/v20/googleoracledatabasegoldengateconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference interface {
	cdktn.ComplexObject
	BootstrapServers() GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesBootstrapServersList
	BootstrapServersInput() interface{}
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
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
	ConsumerPropertiesFile() *string
	SetConsumerPropertiesFile(val *string)
	ConsumerPropertiesFileInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties
	SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties)
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
	ProducerPropertiesFile() *string
	SetProducerPropertiesFile(val *string)
	ProducerPropertiesFileInput() *string
	SecurityProtocol() *string
	SetSecurityProtocol(val *string)
	SecurityProtocolInput() *string
	SslKeyPassword() *string
	SetSslKeyPassword(val *string)
	SslKeyPasswordInput() *string
	SslKeyPasswordSecretVersion() *string
	SetSslKeyPasswordSecretVersion(val *string)
	SslKeyPasswordSecretVersionInput() *string
	StreamPoolId() *string
	SetStreamPoolId(val *string)
	StreamPoolIdInput() *string
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
	UseResourcePrincipal() interface{}
	SetUseResourcePrincipal(val interface{})
	UseResourcePrincipalInput() interface{}
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
	PutBootstrapServers(value interface{})
	ResetBootstrapServers()
	ResetClusterId()
	ResetConsumerPropertiesFile()
	ResetKeyStoreFile()
	ResetKeyStorePassword()
	ResetKeyStorePasswordSecretVersion()
	ResetPassword()
	ResetPasswordSecretVersion()
	ResetProducerPropertiesFile()
	ResetSecurityProtocol()
	ResetSslKeyPassword()
	ResetSslKeyPasswordSecretVersion()
	ResetStreamPoolId()
	ResetTechnologyType()
	ResetTrustStoreFile()
	ResetTrustStorePassword()
	ResetTrustStorePasswordSecretVersion()
	ResetUseResourcePrincipal()
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

// The jsii proxy struct for GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference
type jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) BootstrapServers() GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesBootstrapServersList {
	var returns GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesBootstrapServersList
	_jsii_.Get(
		j,
		"bootstrapServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) BootstrapServersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ConsumerPropertiesFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerPropertiesFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ConsumerPropertiesFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerPropertiesFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) InternalValue() *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties {
	var returns *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) KeyStoreFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStoreFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) KeyStoreFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStoreFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) KeyStorePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) KeyStorePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) KeyStorePasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePasswordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) KeyStorePasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStorePasswordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) PasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) PasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ProducerPropertiesFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"producerPropertiesFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ProducerPropertiesFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"producerPropertiesFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) SecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) SslKeyPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) SslKeyPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) SslKeyPasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyPasswordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) SslKeyPasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyPasswordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) StreamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) StreamPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) TechnologyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) TechnologyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technologyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) TrustStoreFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) TrustStoreFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) TrustStorePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStorePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) TrustStorePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStorePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) TrustStorePasswordSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStorePasswordSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) TrustStorePasswordSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStorePasswordSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) UseResourcePrincipal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourcePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) UseResourcePrincipalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourcePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference_Override(g GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseGoldengateConnection.GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetConsumerPropertiesFile(val *string) {
	if err := j.validateSetConsumerPropertiesFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerPropertiesFile",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetInternalValue(val *GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetKeyStoreFile(val *string) {
	if err := j.validateSetKeyStoreFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyStoreFile",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetKeyStorePassword(val *string) {
	if err := j.validateSetKeyStorePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyStorePassword",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetKeyStorePasswordSecretVersion(val *string) {
	if err := j.validateSetKeyStorePasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyStorePasswordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetPasswordSecretVersion(val *string) {
	if err := j.validateSetPasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetProducerPropertiesFile(val *string) {
	if err := j.validateSetProducerPropertiesFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"producerPropertiesFile",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetSecurityProtocol(val *string) {
	if err := j.validateSetSecurityProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProtocol",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetSslKeyPassword(val *string) {
	if err := j.validateSetSslKeyPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslKeyPassword",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetSslKeyPasswordSecretVersion(val *string) {
	if err := j.validateSetSslKeyPasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslKeyPasswordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetStreamPoolId(val *string) {
	if err := j.validateSetStreamPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamPoolId",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetTechnologyType(val *string) {
	if err := j.validateSetTechnologyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technologyType",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetTrustStoreFile(val *string) {
	if err := j.validateSetTrustStoreFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustStoreFile",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetTrustStorePassword(val *string) {
	if err := j.validateSetTrustStorePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustStorePassword",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetTrustStorePasswordSecretVersion(val *string) {
	if err := j.validateSetTrustStorePasswordSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustStorePasswordSecretVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetUseResourcePrincipal(val interface{}) {
	if err := j.validateSetUseResourcePrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useResourcePrincipal",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) PutBootstrapServers(value interface{}) {
	if err := g.validatePutBootstrapServersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBootstrapServers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetBootstrapServers() {
	_jsii_.InvokeVoid(
		g,
		"resetBootstrapServers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetClusterId() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetConsumerPropertiesFile() {
	_jsii_.InvokeVoid(
		g,
		"resetConsumerPropertiesFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetKeyStoreFile() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyStoreFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetKeyStorePassword() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyStorePassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetKeyStorePasswordSecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyStorePasswordSecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetPasswordSecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetPasswordSecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetProducerPropertiesFile() {
	_jsii_.InvokeVoid(
		g,
		"resetProducerPropertiesFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetSecurityProtocol() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityProtocol",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetSslKeyPassword() {
	_jsii_.InvokeVoid(
		g,
		"resetSslKeyPassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetSslKeyPasswordSecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetSslKeyPasswordSecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetStreamPoolId() {
	_jsii_.InvokeVoid(
		g,
		"resetStreamPoolId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetTechnologyType() {
	_jsii_.InvokeVoid(
		g,
		"resetTechnologyType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetTrustStoreFile() {
	_jsii_.InvokeVoid(
		g,
		"resetTrustStoreFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetTrustStorePassword() {
	_jsii_.InvokeVoid(
		g,
		"resetTrustStorePassword",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetTrustStorePasswordSecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetTrustStorePasswordSecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetUseResourcePrincipal() {
	_jsii_.InvokeVoid(
		g,
		"resetUseResourcePrincipal",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		g,
		"resetUsername",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

