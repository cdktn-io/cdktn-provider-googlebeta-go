// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakeicebergcatalog


type GoogleBiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptionsRefreshSchedule struct {
	// The interval between metadata refreshes, expressed as a duration string (e.g., '300s'). The value must be at least 300s or 0s to disable refresh.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_biglake_iceberg_catalog#refresh_interval GoogleBiglakeIcebergCatalog#refresh_interval}
	RefreshInterval *string `field:"optional" json:"refreshInterval" yaml:"refreshInterval"`
}

