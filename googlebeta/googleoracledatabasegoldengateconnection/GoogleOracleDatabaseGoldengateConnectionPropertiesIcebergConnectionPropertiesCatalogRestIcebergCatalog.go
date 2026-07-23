// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogRestIcebergCatalog struct {
	// The REST uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#uri GoogleOracleDatabaseGoldengateConnection#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The content of the configuration file containing additional properties for the REST catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_oracle_database_goldengate_connection#properties GoogleOracleDatabaseGoldengateConnection#properties}
	Properties *string `field:"optional" json:"properties" yaml:"properties"`
}

