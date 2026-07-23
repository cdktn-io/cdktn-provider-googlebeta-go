// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakeicebergtable


type GoogleBiglakeIcebergTableSortOrderFields struct {
	// The sort direction for the sort field. Possible values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_iceberg_table#direction GoogleBiglakeIcebergTable#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The null ordering for the sort field. Possible values: "nulls-first", "nulls-last".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_iceberg_table#null_order GoogleBiglakeIcebergTable#null_order}
	NullOrder *string `field:"required" json:"nullOrder" yaml:"nullOrder"`
	// The source field ID for the sort field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_iceberg_table#source_id GoogleBiglakeIcebergTable#source_id}
	SourceId *float64 `field:"required" json:"sourceId" yaml:"sourceId"`
	// The transform to apply to the source field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_iceberg_table#transform GoogleBiglakeIcebergTable#transform}
	Transform *string `field:"required" json:"transform" yaml:"transform"`
}

