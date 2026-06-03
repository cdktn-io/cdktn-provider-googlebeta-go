// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesJavaMessageServiceConnectionProperties struct {
	// Authentication type for Java Message Service. Possible values: NONE BASIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#authentication_type GoogleOracleDatabaseGoldengateConnection#authentication_type}
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// The Java class implementing javax.jms.ConnectionFactory interface supplied by the JMS provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#connection_factory GoogleOracleDatabaseGoldengateConnection#connection_factory}
	ConnectionFactory *string `field:"optional" json:"connectionFactory" yaml:"connectionFactory"`
	// Connection URL of the Java Message Service, specifying the protocol, host, and port. e.g.: 'mq://myjms.host.domain:7676'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#connection_url GoogleOracleDatabaseGoldengateConnection#connection_url}
	ConnectionUrl *string `field:"optional" json:"connectionUrl" yaml:"connectionUrl"`
	// The Connection Factory can be looked up using this name. e.g.: 'ConnectionFactory'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#jndi_connection_factory GoogleOracleDatabaseGoldengateConnection#jndi_connection_factory}
	JndiConnectionFactory *string `field:"optional" json:"jndiConnectionFactory" yaml:"jndiConnectionFactory"`
	// The implementation of javax.naming.spi.InitialContextFactory interface used to obtain initial naming context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#jndi_initial_context_factory GoogleOracleDatabaseGoldengateConnection#jndi_initial_context_factory}
	JndiInitialContextFactory *string `field:"optional" json:"jndiInitialContextFactory" yaml:"jndiInitialContextFactory"`
	// The URL that Java Message Service will use to contact the JNDI provider. e.g.: 'tcp://myjms.host.domain:61616?jms.prefetchPolicy.all=1000'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#jndi_provider_url GoogleOracleDatabaseGoldengateConnection#jndi_provider_url}
	JndiProviderUrl *string `field:"optional" json:"jndiProviderUrl" yaml:"jndiProviderUrl"`
	// The password associated to the principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#jndi_security_credentials_secret GoogleOracleDatabaseGoldengateConnection#jndi_security_credentials_secret}
	JndiSecurityCredentialsSecret *string `field:"optional" json:"jndiSecurityCredentialsSecret" yaml:"jndiSecurityCredentialsSecret"`
	// Specifies the identity of the principal (user) to be authenticated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#jndi_security_principal GoogleOracleDatabaseGoldengateConnection#jndi_security_principal}
	JndiSecurityPrincipal *string `field:"optional" json:"jndiSecurityPrincipal" yaml:"jndiSecurityPrincipal"`
	// The content of the KeyStore file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#key_store_file GoogleOracleDatabaseGoldengateConnection#key_store_file}
	KeyStoreFile *string `field:"optional" json:"keyStoreFile" yaml:"keyStoreFile"`
	// Input only. The KeyStore password in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#key_store_password GoogleOracleDatabaseGoldengateConnection#key_store_password}
	KeyStorePassword *string `field:"optional" json:"keyStorePassword" yaml:"keyStorePassword"`
	// Input only. The resource name of a secret version in Secret Manager which contains the KeyStore password. Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#key_store_password_secret_version GoogleOracleDatabaseGoldengateConnection#key_store_password_secret_version}
	KeyStorePasswordSecretVersion *string `field:"optional" json:"keyStorePasswordSecretVersion" yaml:"keyStorePasswordSecretVersion"`
	// Input only. The password Oracle Goldengate uses to connect the Java Message Service in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#password GoogleOracleDatabaseGoldengateConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password Oracle Goldengate uses to connect the associated Java
	// Message Service.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#password_secret_version GoogleOracleDatabaseGoldengateConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
	// Security protocol for Java Message Service. Possible values: PLAIN TLS MTLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#security_protocol GoogleOracleDatabaseGoldengateConnection#security_protocol}
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
	// Input only. The password for the cert inside of the KeyStore in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#ssl_key_password GoogleOracleDatabaseGoldengateConnection#ssl_key_password}
	SslKeyPassword *string `field:"optional" json:"sslKeyPassword" yaml:"sslKeyPassword"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password for the cert inside of the KeyStore.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#ssl_key_password_secret_version GoogleOracleDatabaseGoldengateConnection#ssl_key_password_secret_version}
	SslKeyPasswordSecretVersion *string `field:"optional" json:"sslKeyPasswordSecretVersion" yaml:"sslKeyPasswordSecretVersion"`
	// The technology type of JavaMessageServiceConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// The content of the TrustStore file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#trust_store_file GoogleOracleDatabaseGoldengateConnection#trust_store_file}
	TrustStoreFile *string `field:"optional" json:"trustStoreFile" yaml:"trustStoreFile"`
	// Input only. The TrustStore password in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#trust_store_password GoogleOracleDatabaseGoldengateConnection#trust_store_password}
	TrustStorePassword *string `field:"optional" json:"trustStorePassword" yaml:"trustStorePassword"`
	// Input only. The resource name of a secret version in Secret Manager which contains the TrustStore password. Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#trust_store_password_secret_version GoogleOracleDatabaseGoldengateConnection#trust_store_password_secret_version}
	TrustStorePasswordSecretVersion *string `field:"optional" json:"trustStorePasswordSecretVersion" yaml:"trustStorePasswordSecretVersion"`
	// If set to true, Java Naming and Directory Interface (JNDI) properties should be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#use_jndi GoogleOracleDatabaseGoldengateConnection#use_jndi}
	UseJndi interface{} `field:"optional" json:"useJndi" yaml:"useJndi"`
	// The username Oracle Goldengate uses to connect to the Java Message Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#username GoogleOracleDatabaseGoldengateConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

