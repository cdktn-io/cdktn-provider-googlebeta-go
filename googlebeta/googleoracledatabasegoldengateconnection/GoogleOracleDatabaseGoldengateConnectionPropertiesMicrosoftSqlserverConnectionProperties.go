// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesMicrosoftSqlserverConnectionProperties struct {
	// additional_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#additional_attributes GoogleOracleDatabaseGoldengateConnection#additional_attributes}
	AdditionalAttributes interface{} `field:"optional" json:"additionalAttributes" yaml:"additionalAttributes"`
	// The name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#database GoogleOracleDatabaseGoldengateConnection#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The name or address of a host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#host GoogleOracleDatabaseGoldengateConnection#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Input only. The password Oracle Goldengate uses for Microsoft SQL Server connection in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#password GoogleOracleDatabaseGoldengateConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password Oracle Goldengate uses for Microsoft SQL Server
	// connection.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#password_secret_version GoogleOracleDatabaseGoldengateConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
	// The port of an endpoint usually specified for a connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#port GoogleOracleDatabaseGoldengateConnection#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Security Type for Microsoft SQL Server. Possible values: PLAIN TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#security_protocol GoogleOracleDatabaseGoldengateConnection#security_protocol}
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
	// If set to true, the driver validates the certificate that is sent by the database server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#server_certificate_validation_required GoogleOracleDatabaseGoldengateConnection#server_certificate_validation_required}
	ServerCertificateValidationRequired interface{} `field:"optional" json:"serverCertificateValidationRequired" yaml:"serverCertificateValidationRequired"`
	// Database Certificate - The content of a .pem or .crt file containing the server public key (for 1-way SSL).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#ssl_ca_file GoogleOracleDatabaseGoldengateConnection#ssl_ca_file}
	SslCaFile *string `field:"optional" json:"sslCaFile" yaml:"sslCaFile"`
	// The technology type of MicrosoftSqlserverConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// The username Oracle Goldengate uses to connect to the Microsoft SQL Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#username GoogleOracleDatabaseGoldengateConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

