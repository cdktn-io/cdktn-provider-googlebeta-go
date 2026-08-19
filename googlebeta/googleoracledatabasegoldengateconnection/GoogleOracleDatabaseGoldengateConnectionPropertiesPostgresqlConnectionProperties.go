// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesPostgresqlConnectionProperties struct {
	// additional_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#additional_attributes GoogleOracleDatabaseGoldengateConnection#additional_attributes}
	AdditionalAttributes interface{} `field:"optional" json:"additionalAttributes" yaml:"additionalAttributes"`
	// The name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#database GoogleOracleDatabaseGoldengateConnection#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The OCID of the database system being referenced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#db_system_id GoogleOracleDatabaseGoldengateConnection#db_system_id}
	DbSystemId *string `field:"optional" json:"dbSystemId" yaml:"dbSystemId"`
	// The name or address of a host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#host GoogleOracleDatabaseGoldengateConnection#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Input only. The password Oracle Goldengate uses for PostgreSQL connection in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#password GoogleOracleDatabaseGoldengateConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password Oracle Goldengate uses for PostgreSQL connection.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#password_secret_version GoogleOracleDatabaseGoldengateConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
	// The port of an endpoint usually specified for a connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#port GoogleOracleDatabaseGoldengateConnection#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Security protocol for PostgreSQL. Possible values: PLAIN TLS MTLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#security_protocol GoogleOracleDatabaseGoldengateConnection#security_protocol}
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
	// The certificate of the trusted certificate authorities (Trusted CA) for PostgreSQL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#ssl_ca_file GoogleOracleDatabaseGoldengateConnection#ssl_ca_file}
	SslCaFile *string `field:"optional" json:"sslCaFile" yaml:"sslCaFile"`
	// The certificate of the PostgreSQL server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#ssl_cert_file GoogleOracleDatabaseGoldengateConnection#ssl_cert_file}
	SslCertFile *string `field:"optional" json:"sslCertFile" yaml:"sslCertFile"`
	// The list of certificates revoked by the trusted certificate authorities (Trusted CA).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#ssl_crl_file GoogleOracleDatabaseGoldengateConnection#ssl_crl_file}
	SslCrlFile *string `field:"optional" json:"sslCrlFile" yaml:"sslCrlFile"`
	// The private key of the PostgreSQL server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#ssl_key_file GoogleOracleDatabaseGoldengateConnection#ssl_key_file}
	SslKeyFile *string `field:"optional" json:"sslKeyFile" yaml:"sslKeyFile"`
	// SSL modes for PostgreSQL. Possible values: PREFER REQUIRE VERIFY_CA VERIFY_FULL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#ssl_mode GoogleOracleDatabaseGoldengateConnection#ssl_mode}
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	// The technology type of PostgresqlConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// The username Oracle Goldengate uses to connect the associated system of the given technology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_oracle_database_goldengate_connection#username GoogleOracleDatabaseGoldengateConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

