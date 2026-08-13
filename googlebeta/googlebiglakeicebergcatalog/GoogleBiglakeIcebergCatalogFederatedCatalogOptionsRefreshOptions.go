// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakeicebergcatalog


type GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptions struct {
	// refresh_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_biglake_iceberg_catalog#refresh_schedule GoogleBiglakeIcebergCatalog#refresh_schedule}
	RefreshSchedule *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptionsRefreshSchedule `field:"optional" json:"refreshSchedule" yaml:"refreshSchedule"`
	// refresh_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_biglake_iceberg_catalog#refresh_scope GoogleBiglakeIcebergCatalog#refresh_scope}
	RefreshScope *GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptionsRefreshScope `field:"optional" json:"refreshScope" yaml:"refreshScope"`
}

