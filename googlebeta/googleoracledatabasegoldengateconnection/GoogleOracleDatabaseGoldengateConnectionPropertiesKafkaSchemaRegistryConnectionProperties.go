// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaSchemaRegistryConnectionProperties struct {
	// Used authentication mechanism to access Schema Registry. Possible values: NONE BASIC MUTUAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_oracle_database_goldengate_connection#authentication_type GoogleOracleDatabaseGoldengateConnection#authentication_type}
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// The content of the KeyStore file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_oracle_database_goldengate_connection#key_store_file GoogleOracleDatabaseGoldengateConnection#key_store_file}
	KeyStoreFile *string `field:"optional" json:"keyStoreFile" yaml:"keyStoreFile"`
	// Input only. The KeyStore password in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_oracle_database_goldengate_connection#key_store_password GoogleOracleDatabaseGoldengateConnection#key_store_password}
	KeyStorePassword *string `field:"optional" json:"keyStorePassword" yaml:"keyStorePassword"`
	// Input only. The resource name of a secret version in Secret Manager which contains the KeyStore password. Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_oracle_database_goldengate_connection#key_store_password_secret_version GoogleOracleDatabaseGoldengateConnection#key_store_password_secret_version}
	KeyStorePasswordSecretVersion *string `field:"optional" json:"keyStorePasswordSecretVersion" yaml:"keyStorePasswordSecretVersion"`
	// Input only. The password to access Schema Registry in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_oracle_database_goldengate_connection#password GoogleOracleDatabaseGoldengateConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password to access Schema Registry using basic authentication.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_oracle_database_goldengate_connection#password_secret_version GoogleOracleDatabaseGoldengateConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
	// Input only. The password for the cert inside the KeyStore in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_oracle_database_goldengate_connection#ssl_key_password GoogleOracleDatabaseGoldengateConnection#ssl_key_password}
	SslKeyPassword *string `field:"optional" json:"sslKeyPassword" yaml:"sslKeyPassword"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password for the cert inside the KeyStore.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_oracle_database_goldengate_connection#ssl_key_password_secret_version GoogleOracleDatabaseGoldengateConnection#ssl_key_password_secret_version}
	SslKeyPasswordSecretVersion *string `field:"optional" json:"sslKeyPasswordSecretVersion" yaml:"sslKeyPasswordSecretVersion"`
	// The technology type of KafkaSchemaRegistryConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// The content of the TrustStore file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_oracle_database_goldengate_connection#trust_store_file GoogleOracleDatabaseGoldengateConnection#trust_store_file}
	TrustStoreFile *string `field:"optional" json:"trustStoreFile" yaml:"trustStoreFile"`
	// Input only. The TrustStore password in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_oracle_database_goldengate_connection#trust_store_password GoogleOracleDatabaseGoldengateConnection#trust_store_password}
	TrustStorePassword *string `field:"optional" json:"trustStorePassword" yaml:"trustStorePassword"`
	// Input only. The resource name of a secret version in Secret Manager which contains the TrustStore password. Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_oracle_database_goldengate_connection#trust_store_password_secret_version GoogleOracleDatabaseGoldengateConnection#trust_store_password_secret_version}
	TrustStorePasswordSecretVersion *string `field:"optional" json:"trustStorePasswordSecretVersion" yaml:"trustStorePasswordSecretVersion"`
	// Kafka Schema Registry URL. e.g.: 'https://server1.us.oracle.com:8081'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_oracle_database_goldengate_connection#url GoogleOracleDatabaseGoldengateConnection#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The username to access Schema Registry using basic authentication. This value is injected into 'schema.registry.basic.auth.user.info=user:password' configuration property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_oracle_database_goldengate_connection#username GoogleOracleDatabaseGoldengateConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

