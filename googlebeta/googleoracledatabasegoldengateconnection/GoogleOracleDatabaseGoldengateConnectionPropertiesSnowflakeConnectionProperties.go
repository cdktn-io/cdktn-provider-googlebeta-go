// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesSnowflakeConnectionProperties struct {
	// Used authentication mechanism to access Snowflake. Possible values: BASIC KEY_PAIR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#authentication_type GoogleOracleDatabaseGoldengateConnection#authentication_type}
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// JDBC connection URL. e.g.: 'jdbc:snowflake://.snowflakecomputing.com/?warehouse=&db='.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#connection_url GoogleOracleDatabaseGoldengateConnection#connection_url}
	ConnectionUrl *string `field:"optional" json:"connectionUrl" yaml:"connectionUrl"`
	// Input only. The password Oracle Goldengate uses to connect to Snowflake platform in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#password GoogleOracleDatabaseGoldengateConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password Oracle Goldengate uses to connect to Snowflake platform.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#password_secret_version GoogleOracleDatabaseGoldengateConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
	// The content of private key file in PEM format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#private_key_file GoogleOracleDatabaseGoldengateConnection#private_key_file}
	PrivateKeyFile *string `field:"optional" json:"privateKeyFile" yaml:"privateKeyFile"`
	// Password if the private key file is encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#private_key_passphrase_secret GoogleOracleDatabaseGoldengateConnection#private_key_passphrase_secret}
	PrivateKeyPassphraseSecret *string `field:"optional" json:"privateKeyPassphraseSecret" yaml:"privateKeyPassphraseSecret"`
	// The technology type of SnowflakeConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// The username Oracle Goldengate uses to connect to Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#username GoogleOracleDatabaseGoldengateConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

