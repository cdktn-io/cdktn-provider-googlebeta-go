// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlemigrationcenterassetsexportjob


type GoogleMigrationCenterAssetsExportJobCondition struct {
	// Assets filter, supports the same syntax as asset listing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_migration_center_assets_export_job#filter GoogleMigrationCenterAssetsExportJob#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
}

