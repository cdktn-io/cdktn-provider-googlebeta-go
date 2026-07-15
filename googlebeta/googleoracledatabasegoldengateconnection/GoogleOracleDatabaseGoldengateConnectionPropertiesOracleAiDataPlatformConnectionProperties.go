// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesOracleAiDataPlatformConnectionProperties struct {
	// Connection URL. It must start with 'jdbc:spark://'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#connection_url GoogleOracleDatabaseGoldengateConnection#connection_url}
	ConnectionUrl *string `field:"optional" json:"connectionUrl" yaml:"connectionUrl"`
	// The content of the private key file (PEM file) corresponding to the API key of the fingerprint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#private_key_file GoogleOracleDatabaseGoldengateConnection#private_key_file}
	PrivateKeyFile *string `field:"optional" json:"privateKeyFile" yaml:"privateKeyFile"`
	// The passphrase of the private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#private_key_passphrase_secret GoogleOracleDatabaseGoldengateConnection#private_key_passphrase_secret}
	PrivateKeyPassphraseSecret *string `field:"optional" json:"privateKeyPassphraseSecret" yaml:"privateKeyPassphraseSecret"`
	// The fingerprint of the API Key of the user specified by the user_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#public_key_fingerprint GoogleOracleDatabaseGoldengateConnection#public_key_fingerprint}
	PublicKeyFingerprint *string `field:"optional" json:"publicKeyFingerprint" yaml:"publicKeyFingerprint"`
	// The name of the region. e.g.: us-ashburn-1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#region GoogleOracleDatabaseGoldengateConnection#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The technology type of OracleAiDataPlatformConnection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"optional" json:"technologyType" yaml:"technologyType"`
	// The OCID of the related OCI tenancy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#tenancy_id GoogleOracleDatabaseGoldengateConnection#tenancy_id}
	TenancyId *string `field:"optional" json:"tenancyId" yaml:"tenancyId"`
	// Specifies that the user intends to authenticate to the instance using a resource principal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#use_resource_principal GoogleOracleDatabaseGoldengateConnection#use_resource_principal}
	UseResourcePrincipal interface{} `field:"optional" json:"useResourcePrincipal" yaml:"useResourcePrincipal"`
	// The OCID of the OCI user who will access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#user_id GoogleOracleDatabaseGoldengateConnection#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

