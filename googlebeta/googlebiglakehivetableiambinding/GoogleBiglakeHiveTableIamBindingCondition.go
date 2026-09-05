// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakehivetableiambinding


type GoogleBiglakeHiveTableIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_hive_table_iam_binding#expression GoogleBiglakeHiveTableIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_hive_table_iam_binding#title GoogleBiglakeHiveTableIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_hive_table_iam_binding#description GoogleBiglakeHiveTableIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

