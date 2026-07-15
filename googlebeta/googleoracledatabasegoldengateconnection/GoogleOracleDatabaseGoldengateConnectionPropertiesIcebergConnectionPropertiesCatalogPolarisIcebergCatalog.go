// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogPolarisIcebergCatalog struct {
	// The Polaris client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#client_id GoogleOracleDatabaseGoldengateConnection#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The catalog name within Polaris.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#polaris_catalog GoogleOracleDatabaseGoldengateConnection#polaris_catalog}
	PolarisCatalog *string `field:"required" json:"polarisCatalog" yaml:"polarisCatalog"`
	// The Polaris principal role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#principal_role GoogleOracleDatabaseGoldengateConnection#principal_role}
	PrincipalRole *string `field:"required" json:"principalRole" yaml:"principalRole"`
	// The Polaris uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#uri GoogleOracleDatabaseGoldengateConnection#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The Polaris client secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#client_secret GoogleOracleDatabaseGoldengateConnection#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

