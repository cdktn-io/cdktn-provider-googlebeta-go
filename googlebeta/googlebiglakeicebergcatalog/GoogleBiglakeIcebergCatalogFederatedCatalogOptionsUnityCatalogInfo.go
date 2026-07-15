// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakeicebergcatalog


type GoogleBiglakeIcebergCatalogFederatedCatalogOptionsUnityCatalogInfo struct {
	// The name of the catalog within the Unity Catalog instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_biglake_iceberg_catalog#catalog_name GoogleBiglakeIcebergCatalog#catalog_name}
	CatalogName *string `field:"required" json:"catalogName" yaml:"catalogName"`
	// The Databricks workspace instance name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_biglake_iceberg_catalog#instance_name GoogleBiglakeIcebergCatalog#instance_name}
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// The application ID of the Databricks service principal for OIDC authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_biglake_iceberg_catalog#service_principal_application_id GoogleBiglakeIcebergCatalog#service_principal_application_id}
	ServicePrincipalApplicationId *string `field:"optional" json:"servicePrincipalApplicationId" yaml:"servicePrincipalApplicationId"`
}

