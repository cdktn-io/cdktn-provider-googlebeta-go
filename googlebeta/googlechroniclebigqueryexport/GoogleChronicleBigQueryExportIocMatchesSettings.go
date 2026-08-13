// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechroniclebigqueryexport


type GoogleChronicleBigQueryExportIocMatchesSettings struct {
	// Whether the data source is enabled for export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_big_query_export#enabled GoogleChronicleBigQueryExport#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The retention period for the data source in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_chronicle_big_query_export#retention_days GoogleChronicleBigQueryExport#retention_days}
	RetentionDays *float64 `field:"required" json:"retentionDays" yaml:"retentionDays"`
}

