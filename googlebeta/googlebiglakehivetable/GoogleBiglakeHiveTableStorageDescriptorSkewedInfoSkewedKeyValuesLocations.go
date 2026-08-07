// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakehivetable


type GoogleBiglakeHiveTableStorageDescriptorSkewedInfoSkewedKeyValuesLocations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_biglake_hive_table#location GoogleBiglakeHiveTable#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_biglake_hive_table#values GoogleBiglakeHiveTable#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

