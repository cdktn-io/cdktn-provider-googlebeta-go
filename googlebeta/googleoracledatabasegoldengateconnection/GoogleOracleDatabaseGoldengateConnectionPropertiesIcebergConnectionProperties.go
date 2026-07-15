// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionProperties struct {
	// catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#catalog GoogleOracleDatabaseGoldengateConnection#catalog}
	Catalog *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalog `field:"required" json:"catalog" yaml:"catalog"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#storage GoogleOracleDatabaseGoldengateConnection#storage}
	Storage *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesStorage `field:"required" json:"storage" yaml:"storage"`
	// The technology type of Iceberg connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_oracle_database_goldengate_connection#technology_type GoogleOracleDatabaseGoldengateConnection#technology_type}
	TechnologyType *string `field:"required" json:"technologyType" yaml:"technologyType"`
}

