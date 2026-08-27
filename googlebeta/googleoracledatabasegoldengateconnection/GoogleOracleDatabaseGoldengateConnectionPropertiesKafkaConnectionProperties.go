// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesKafkaConnectionProperties struct {
	// bootstrap_servers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#bootstrap_servers GoogleOracleDatabaseGoldengateConnection#bootstrap_servers}
	BootstrapServers interface{} `field:"optional" json:"bootstrapServers" yaml:"bootstrapServers"`
	// The OCID of the Kafka cluster being referenced from OCI Streaming with Apache Kafka.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#cluster_id GoogleOracleDatabaseGoldengateConnection#cluster_id}
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// The content of the consumer.properties file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#consumer_properties_file GoogleOracleDatabaseGoldengateConnection#consumer_properties_file}
	ConsumerPropertiesFile *string `field:"optional" json:"consumerPropertiesFile" yaml:"consumerPropertiesFile"`
	// The content of the KeyStore file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#key_store_file GoogleOracleDatabaseGoldengateConnection#key_store_file}
	KeyStoreFile *string `field:"optional" json:"keyStoreFile" yaml:"keyStoreFile"`
	// Input only. The KeyStore password in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#key_store_password GoogleOracleDatabaseGoldengateConnection#key_store_password}
	KeyStorePassword *string `field:"optional" json:"keyStorePassword" yaml:"keyStorePassword"`
	// Input only. The resource name of a secret version in Secret Manager which contains the KeyStore password. Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#key_store_password_secret_version GoogleOracleDatabaseGoldengateConnection#key_store_password_secret_version}
	KeyStorePasswordSecretVersion *string `field:"optional" json:"keyStorePasswordSecretVersion" yaml:"keyStorePasswordSecretVersion"`
	// Input only. The password for Kafka basic/SASL auth in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#password GoogleOracleDatabaseGoldengateConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password for Kafka basic/SASL auth.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#password_secret_version GoogleOracleDatabaseGoldengateConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
	// The content of the producer.properties file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#producer_properties_file GoogleOracleDatabaseGoldengateConnection#producer_properties_file}
	ProducerPropertiesFile *string `field:"optional" json:"producerPropertiesFile" yaml:"producerPropertiesFile"`
	// Security Type for Kafka. Possible values: SSL SASL_SSL PLAINTEXT SASL_PLAINTEXT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#security_protocol GoogleOracleDatabaseGoldengateConnection#security_protocol}
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
	// Input only. The password for the cert inside of the KeyStore in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#ssl_key_password GoogleOracleDatabaseGoldengateConnection#ssl_key_password}
	SslKeyPassword *string `field:"optional" json:"sslKeyPassword" yaml:"sslKeyPassword"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password for the cert inside of the KeyStore.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#ssl_key_password_secret_version GoogleOracleDatabaseGoldengateConnection#ssl_key_password_secret_version}
	SslKeyPasswordSecretVersion *string `field:"optional" json:"sslKeyPasswordSecretVersion" yaml:"sslKeyPasswordSecretVersion"`
	// The OCID of the stream pool being referenced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#stream_pool_id GoogleOracleDatabaseGoldengateConnection#stream_pool_id}
	StreamPoolId *string `field:"optional" json:"streamPoolId" yaml:"streamPoolId"`
	// The technology type of KafkaConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// The content of the TrustStore file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#trust_store_file GoogleOracleDatabaseGoldengateConnection#trust_store_file}
	TrustStoreFile *string `field:"optional" json:"trustStoreFile" yaml:"trustStoreFile"`
	// Input only. The TrustStore password in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#trust_store_password GoogleOracleDatabaseGoldengateConnection#trust_store_password}
	TrustStorePassword *string `field:"optional" json:"trustStorePassword" yaml:"trustStorePassword"`
	// Input only. The resource name of a secret version in Secret Manager which contains the TrustStore password. Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#trust_store_password_secret_version GoogleOracleDatabaseGoldengateConnection#trust_store_password_secret_version}
	TrustStorePasswordSecretVersion *string `field:"optional" json:"trustStorePasswordSecretVersion" yaml:"trustStorePasswordSecretVersion"`
	// Specifies that the user intends to authenticate to the instance using a resource principal. Applicable only for OCI Streaming connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#use_resource_principal GoogleOracleDatabaseGoldengateConnection#use_resource_principal}
	UseResourcePrincipal interface{} `field:"optional" json:"useResourcePrincipal" yaml:"useResourcePrincipal"`
	// The username Oracle Goldengate uses to connect the associated system of the given technology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_oracle_database_goldengate_connection#username GoogleOracleDatabaseGoldengateConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

