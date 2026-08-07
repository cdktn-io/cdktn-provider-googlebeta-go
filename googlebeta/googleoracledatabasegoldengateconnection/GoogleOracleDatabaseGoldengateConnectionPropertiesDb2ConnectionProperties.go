// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesDb2ConnectionProperties struct {
	// additional_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_oracle_database_goldengate_connection#additional_attributes GoogleOracleDatabaseGoldengateConnection#additional_attributes}
	AdditionalAttributes interface{} `field:"optional" json:"additionalAttributes" yaml:"additionalAttributes"`
	// The name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_oracle_database_goldengate_connection#database GoogleOracleDatabaseGoldengateConnection#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The name or address of a host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_oracle_database_goldengate_connection#host GoogleOracleDatabaseGoldengateConnection#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Input only. The password Oracle Goldengate uses for Db2 connection in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_oracle_database_goldengate_connection#password GoogleOracleDatabaseGoldengateConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password Oracle Goldengate uses for Db2 connection.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_oracle_database_goldengate_connection#password_secret_version GoogleOracleDatabaseGoldengateConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
	// The port of an endpoint usually specified for a connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_oracle_database_goldengate_connection#port GoogleOracleDatabaseGoldengateConnection#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Security protocol for the DB2 database. Possible values: PLAIN TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_oracle_database_goldengate_connection#security_protocol GoogleOracleDatabaseGoldengateConnection#security_protocol}
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
	// The keystash file which contains the encrypted password to the key database file. Not supported for IBM Db2 for i.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_oracle_database_goldengate_connection#ssl_client_keystash_file GoogleOracleDatabaseGoldengateConnection#ssl_client_keystash_file}
	SslClientKeystashFile *string `field:"optional" json:"sslClientKeystashFile" yaml:"sslClientKeystashFile"`
	// The keystore file created at the client containing the server certificate / CA root certificate.
	//
	// Not supported for IBM Db2 for i.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_oracle_database_goldengate_connection#ssl_client_keystoredb_file GoogleOracleDatabaseGoldengateConnection#ssl_client_keystoredb_file}
	SslClientKeystoredbFile *string `field:"optional" json:"sslClientKeystoredbFile" yaml:"sslClientKeystoredbFile"`
	// The file which contains the self-signed server certificate / Certificate Authority (CA) certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_oracle_database_goldengate_connection#ssl_server_certificate_file GoogleOracleDatabaseGoldengateConnection#ssl_server_certificate_file}
	SslServerCertificateFile *string `field:"optional" json:"sslServerCertificateFile" yaml:"sslServerCertificateFile"`
	// The technology type of Db2Connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// The username Oracle Goldengate uses to connect to the DB2 database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_oracle_database_goldengate_connection#username GoogleOracleDatabaseGoldengateConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

