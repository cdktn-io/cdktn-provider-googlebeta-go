// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogNessieIcebergCatalog struct {
	// The Nessie branch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#branch GoogleOracleDatabaseGoldengateConnection#branch}
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// The Nessie uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_oracle_database_goldengate_connection#uri GoogleOracleDatabaseGoldengateConnection#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

