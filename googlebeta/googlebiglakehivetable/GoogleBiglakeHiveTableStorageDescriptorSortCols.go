// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakehivetable


type GoogleBiglakeHiveTableStorageDescriptorSortCols struct {
	// The column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_table#col GoogleBiglakeHiveTable#col}
	Col *string `field:"required" json:"col" yaml:"col"`
	// Sort order: 1 for Ascending, 0 for Descending.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_table#order GoogleBiglakeHiveTable#order}
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

