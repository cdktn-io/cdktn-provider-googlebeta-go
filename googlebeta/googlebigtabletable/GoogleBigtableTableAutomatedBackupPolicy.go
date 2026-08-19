// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigtabletable


type GoogleBigtableTableAutomatedBackupPolicy struct {
	// How frequently automated backups should occur.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_bigtable_table#frequency GoogleBigtableTable#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// A list of Cloud Bigtable zones where automated backups are allowed to be created.
	//
	// If empty, automated backups will be created in all zones of the instance. Locations are in the format projects/{project}/locations/{zone}. This field can only be set for tables in Enterprise Plus instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_bigtable_table#locations GoogleBigtableTable#locations}
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
	// How long the automated backups should be retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_bigtable_table#retention_period GoogleBigtableTable#retention_period}
	RetentionPeriod *string `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
}

