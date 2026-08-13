// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleoracledatabasegoldengateconnection


type GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalog struct {
	// The type of Iceberg catalog. Possible values: GLUE HADOOP NESSIE POLARIS REST.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#catalog_type GoogleOracleDatabaseGoldengateConnection#catalog_type}
	CatalogType *string `field:"required" json:"catalogType" yaml:"catalogType"`
	// glue_iceberg_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#glue_iceberg_catalog GoogleOracleDatabaseGoldengateConnection#glue_iceberg_catalog}
	GlueIcebergCatalog *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogGlueIcebergCatalog `field:"optional" json:"glueIcebergCatalog" yaml:"glueIcebergCatalog"`
	// nessie_iceberg_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#nessie_iceberg_catalog GoogleOracleDatabaseGoldengateConnection#nessie_iceberg_catalog}
	NessieIcebergCatalog *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogNessieIcebergCatalog `field:"optional" json:"nessieIcebergCatalog" yaml:"nessieIcebergCatalog"`
	// polaris_iceberg_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#polaris_iceberg_catalog GoogleOracleDatabaseGoldengateConnection#polaris_iceberg_catalog}
	PolarisIcebergCatalog *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogPolarisIcebergCatalog `field:"optional" json:"polarisIcebergCatalog" yaml:"polarisIcebergCatalog"`
	// rest_iceberg_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_oracle_database_goldengate_connection#rest_iceberg_catalog GoogleOracleDatabaseGoldengateConnection#rest_iceberg_catalog}
	RestIcebergCatalog *GoogleOracleDatabaseGoldengateConnectionPropertiesIcebergConnectionPropertiesCatalogRestIcebergCatalog `field:"optional" json:"restIcebergCatalog" yaml:"restIcebergCatalog"`
}

