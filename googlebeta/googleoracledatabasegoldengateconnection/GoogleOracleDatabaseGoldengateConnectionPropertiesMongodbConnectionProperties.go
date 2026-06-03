// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesMongodbConnectionProperties struct {
	// MongoDB connection string. e.g.: 'mongodb://mongodb0.example.com:27017/recordsrecords'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#connection_string GoogleOracleDatabaseGoldengateConnection#connection_string}
	ConnectionString *string `field:"optional" json:"connectionString" yaml:"connectionString"`
	// The OCID of the Oracle Autonomous Json Database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#database_id GoogleOracleDatabaseGoldengateConnection#database_id}
	DatabaseId *string `field:"optional" json:"databaseId" yaml:"databaseId"`
	// Input only. The password Oracle Goldengate uses to connect the Mongodb connection in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#password GoogleOracleDatabaseGoldengateConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password Oracle Goldengate uses to connect the Mongodb connection.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#password_secret_version GoogleOracleDatabaseGoldengateConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
	// Security Type for MongoDB. Possible values: PLAIN TLS MTLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#security_protocol GoogleOracleDatabaseGoldengateConnection#security_protocol}
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
	// The technology type of MongodbConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// Database Certificate - The content of a .pem file, containing the server public key (for 1 and 2-way SSL).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#tls_ca_file GoogleOracleDatabaseGoldengateConnection#tls_ca_file}
	TlsCaFile *string `field:"optional" json:"tlsCaFile" yaml:"tlsCaFile"`
	// Client Certificate - The content of a .pem file, containing the client public key (for 2-way SSL).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#tls_certificate_key_file GoogleOracleDatabaseGoldengateConnection#tls_certificate_key_file}
	TlsCertificateKeyFile *string `field:"optional" json:"tlsCertificateKeyFile" yaml:"tlsCertificateKeyFile"`
	// Input only. The Client Certificate key file password in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#tls_certificate_key_file_password GoogleOracleDatabaseGoldengateConnection#tls_certificate_key_file_password}
	TlsCertificateKeyFilePassword *string `field:"optional" json:"tlsCertificateKeyFilePassword" yaml:"tlsCertificateKeyFilePassword"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the Client Certificate key file password in Secret Manager.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#tls_certificate_key_file_password_secret_version GoogleOracleDatabaseGoldengateConnection#tls_certificate_key_file_password_secret_version}
	TlsCertificateKeyFilePasswordSecretVersion *string `field:"optional" json:"tlsCertificateKeyFilePasswordSecretVersion" yaml:"tlsCertificateKeyFilePasswordSecretVersion"`
	// The username Oracle Goldengate uses to connect to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.35.0/docs/resources/google_oracle_database_goldengate_connection#username GoogleOracleDatabaseGoldengateConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

