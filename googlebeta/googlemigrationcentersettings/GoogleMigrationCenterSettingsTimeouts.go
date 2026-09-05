// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlemigrationcentersettings


type GoogleMigrationCenterSettingsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_migration_center_settings#create GoogleMigrationCenterSettings#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_migration_center_settings#delete GoogleMigrationCenterSettings#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_migration_center_settings#update GoogleMigrationCenterSettings#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

