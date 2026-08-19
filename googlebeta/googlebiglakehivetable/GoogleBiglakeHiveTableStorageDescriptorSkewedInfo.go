// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakehivetable


type GoogleBiglakeHiveTableStorageDescriptorSkewedInfo struct {
	// The column names that are skewed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_biglake_hive_table#skewed_col_names GoogleBiglakeHiveTable#skewed_col_names}
	SkewedColNames *[]*string `field:"required" json:"skewedColNames" yaml:"skewedColNames"`
	// skewed_col_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_biglake_hive_table#skewed_col_values GoogleBiglakeHiveTable#skewed_col_values}
	SkewedColValues interface{} `field:"required" json:"skewedColValues" yaml:"skewedColValues"`
	// skewed_key_values_locations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_biglake_hive_table#skewed_key_values_locations GoogleBiglakeHiveTable#skewed_key_values_locations}
	SkewedKeyValuesLocations interface{} `field:"required" json:"skewedKeyValuesLocations" yaml:"skewedKeyValuesLocations"`
}

