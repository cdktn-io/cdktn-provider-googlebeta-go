// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlemigrationcenterassetsexportjob


type GoogleMigrationCenterAssetsExportJobPerformanceData struct {
	// When this value is set to a positive integer, performance data will be returned for the most recent days for which data is available.
	//
	// When this value is unset (or set to zero),
	// all available data is returned.
	// The maximum value is 420; values above 420 will be coerced to 420.
	// If unset (0 value) a default value of 40 will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_migration_center_assets_export_job#max_days GoogleMigrationCenterAssetsExportJob#max_days}
	MaxDays *float64 `field:"optional" json:"maxDays" yaml:"maxDays"`
}

