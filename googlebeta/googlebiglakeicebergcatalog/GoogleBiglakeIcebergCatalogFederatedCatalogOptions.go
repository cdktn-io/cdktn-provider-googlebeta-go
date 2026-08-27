// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakeicebergcatalog


type GoogleBiglakeIcebergCatalogFederatedCatalogOptions struct {
	// glue_catalog_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_biglake_iceberg_catalog#glue_catalog_info GoogleBiglakeIcebergCatalog#glue_catalog_info}
	GlueCatalogInfo *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsGlueCatalogInfo `field:"optional" json:"glueCatalogInfo" yaml:"glueCatalogInfo"`
	// refresh_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_biglake_iceberg_catalog#refresh_options GoogleBiglakeIcebergCatalog#refresh_options}
	RefreshOptions *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptions `field:"optional" json:"refreshOptions" yaml:"refreshOptions"`
	// The secret resource name in Secret Manager, in the format 'projects/{projectId}/locations/{location}/secrets/{secret_id}'. Used to store credentials for authenticating with the remote catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_biglake_iceberg_catalog#secret_name GoogleBiglakeIcebergCatalog#secret_name}
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// The Service Directory service name for private network connectivity through Cross-Cloud Interconnect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_biglake_iceberg_catalog#service_directory_name GoogleBiglakeIcebergCatalog#service_directory_name}
	ServiceDirectoryName *string `field:"optional" json:"serviceDirectoryName" yaml:"serviceDirectoryName"`
	// unity_catalog_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_biglake_iceberg_catalog#unity_catalog_info GoogleBiglakeIcebergCatalog#unity_catalog_info}
	UnityCatalogInfo *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo `field:"optional" json:"unityCatalogInfo" yaml:"unityCatalogInfo"`
}

