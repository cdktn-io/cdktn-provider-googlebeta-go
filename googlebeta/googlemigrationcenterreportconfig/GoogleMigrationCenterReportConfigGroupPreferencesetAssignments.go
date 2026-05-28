// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlemigrationcenterreportconfig


type GoogleMigrationCenterReportConfigGroupPreferencesetAssignments struct {
	// Name of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_migration_center_report_config#group GoogleMigrationCenterReportConfig#group}
	Group *string `field:"required" json:"group" yaml:"group"`
	// Name of the Preference Set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.34.0/docs/resources/google_migration_center_report_config#preference_set GoogleMigrationCenterReportConfig#preference_set}
	PreferenceSet *string `field:"required" json:"preferenceSet" yaml:"preferenceSet"`
}

