// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesRedisConnectionProperties struct {
	// Authentication type for Redis. Possible values: NONE BASIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#authentication_type GoogleOracleDatabaseGoldengateConnection#authentication_type}
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// The content of the KeyStore file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#key_store_file GoogleOracleDatabaseGoldengateConnection#key_store_file}
	KeyStoreFile *string `field:"optional" json:"keyStoreFile" yaml:"keyStoreFile"`
	// Input only. The KeyStore password in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#key_store_password GoogleOracleDatabaseGoldengateConnection#key_store_password}
	KeyStorePassword *string `field:"optional" json:"keyStorePassword" yaml:"keyStorePassword"`
	// Input only. The resource name of a secret version in Secret Manager which contains the KeyStore password. Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#key_store_password_secret_version GoogleOracleDatabaseGoldengateConnection#key_store_password_secret_version}
	KeyStorePasswordSecretVersion *string `field:"optional" json:"keyStorePasswordSecretVersion" yaml:"keyStorePasswordSecretVersion"`
	// Input only. The password Oracle Goldengate uses for Redis connection in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#password GoogleOracleDatabaseGoldengateConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password Oracle Goldengate uses for Redis connection.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#password_secret_version GoogleOracleDatabaseGoldengateConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
	// The OCID of the Redis cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#redis_cluster_id GoogleOracleDatabaseGoldengateConnection#redis_cluster_id}
	RedisClusterId *string `field:"optional" json:"redisClusterId" yaml:"redisClusterId"`
	// Security protocol for Redis. Possible values: PLAIN TLS MTLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#security_protocol GoogleOracleDatabaseGoldengateConnection#security_protocol}
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
	// Comma separated list of Redis server addresses, specified as host:port entries, where :port is optional.
	//
	// If port is not specified, it defaults
	// to 6379. Example: "server1.example.com:6379,server2.example.com:6379"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#servers GoogleOracleDatabaseGoldengateConnection#servers}
	Servers *string `field:"optional" json:"servers" yaml:"servers"`
	// The technology type of RedisConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// The content of the TrustStore file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#trust_store_file GoogleOracleDatabaseGoldengateConnection#trust_store_file}
	TrustStoreFile *string `field:"optional" json:"trustStoreFile" yaml:"trustStoreFile"`
	// Input only. The TrustStore password in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#trust_store_password GoogleOracleDatabaseGoldengateConnection#trust_store_password}
	TrustStorePassword *string `field:"optional" json:"trustStorePassword" yaml:"trustStorePassword"`
	// Input only. The resource name of a secret version in Secret Manager which contains the TrustStore password. Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#trust_store_password_secret_version GoogleOracleDatabaseGoldengateConnection#trust_store_password_secret_version}
	TrustStorePasswordSecretVersion *string `field:"optional" json:"trustStorePasswordSecretVersion" yaml:"trustStorePasswordSecretVersion"`
	// The username Oracle Goldengate uses to connect the associated system of the given technology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#username GoogleOracleDatabaseGoldengateConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

