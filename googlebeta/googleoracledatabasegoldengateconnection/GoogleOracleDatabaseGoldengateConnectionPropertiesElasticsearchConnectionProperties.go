// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesElasticsearchConnectionProperties struct {
	// Authentication type for Elasticsearch. Possible values: NONE BASIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#authentication_type GoogleOracleDatabaseGoldengateConnection#authentication_type}
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Fingerprint required by TLS security protocol. Eg.: '6152b2dfbff200f973c5074a5b91d06ab3b472c07c09a1ea57bb7fd406cdce9c'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#fingerprint GoogleOracleDatabaseGoldengateConnection#fingerprint}
	Fingerprint *string `field:"optional" json:"fingerprint" yaml:"fingerprint"`
	// Input only. The password Oracle Goldengate uses for Elastic Search connection in plain text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#password GoogleOracleDatabaseGoldengateConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Input only.
	//
	// The resource name of a secret version in Secret Manager which contains
	// the password Oracle Goldengate uses for Elastic Search connection.
	// Format: projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#password_secret_version GoogleOracleDatabaseGoldengateConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
	// Security protocol for Elasticsearch. Possible values: PLAIN TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#security_protocol GoogleOracleDatabaseGoldengateConnection#security_protocol}
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
	// Comma separated list of Elasticsearch server addresses, specified as host:port entries, where :port is optional.
	//
	// If port is not specified, it
	// defaults to 9200. Example:
	// "server1.example.com:4000,server2.example.com:4000"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#servers GoogleOracleDatabaseGoldengateConnection#servers}
	Servers *string `field:"optional" json:"servers" yaml:"servers"`
	// The technology type of ElasticsearchConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// The username Oracle Goldengate uses to connect the associated system of the given technology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#username GoogleOracleDatabaseGoldengateConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

