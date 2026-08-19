// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakeicebergcatalog


type GoogleBiglakeIcebergCatalogRestrictedLocationsConfig struct {
	// A list of GCS locations (e.g., 'gs://my-other-bucket/...') that are permitted for use by resources within this catalog. Each entry can be either a GCS bucket or a path within it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_biglake_iceberg_catalog#restricted_locations GoogleBiglakeIcebergCatalog#restricted_locations}
	RestrictedLocations *[]*string `field:"optional" json:"restrictedLocations" yaml:"restrictedLocations"`
}

