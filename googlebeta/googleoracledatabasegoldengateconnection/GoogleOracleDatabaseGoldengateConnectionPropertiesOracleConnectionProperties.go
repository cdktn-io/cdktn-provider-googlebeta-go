// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesOracleConnectionProperties struct {
	// Authentication mode. Possible values: TLS MTLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#authentication_mode GoogleOracleDatabaseGoldengateConnection#authentication_mode}
	AuthenticationMode *string `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Connect descriptor or Easy Connect Naming method used to connect to a database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#connection_string GoogleOracleDatabaseGoldengateConnection#connection_string}
	ConnectionString *string `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Database instance id of database in Oracle Database @ Google Cloud. If gcp_oracle_database_id is provided, connection_string must be empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#gcp_oracle_database_id GoogleOracleDatabaseGoldengateConnection#gcp_oracle_database_id}
	GcpOracleDatabaseId *string `field:"optional" json:"gcpOracleDatabaseId" yaml:"gcpOracleDatabaseId"`
	// Input only. The password Oracle Goldengate uses in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#password GoogleOracleDatabaseGoldengateConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Input only. The resource name of a secret version in Secret Manager which contains the password Oracle Goldengate uses. Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#password_secret_version GoogleOracleDatabaseGoldengateConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
	// The mode of the database connection session to be established by the data client. Possible values: DIRECT REDIRECT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#session_mode GoogleOracleDatabaseGoldengateConnection#session_mode}
	SessionMode *string `field:"optional" json:"sessionMode" yaml:"sessionMode"`
	// The technology type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// The username Oracle Goldengate uses to connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#username GoogleOracleDatabaseGoldengateConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// The wallet contents Oracle Goldengate uses to make connections to a database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_oracle_database_goldengate_connection#wallet_file GoogleOracleDatabaseGoldengateConnection#wallet_file}
	WalletFile *string `field:"optional" json:"walletFile" yaml:"walletFile"`
}

