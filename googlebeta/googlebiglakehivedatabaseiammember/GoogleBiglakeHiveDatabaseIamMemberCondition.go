// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakehivedatabaseiammember


type GoogleBiglakeHiveDatabaseIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_hive_database_iam_member#expression GoogleBiglakeHiveDatabaseIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_hive_database_iam_member#title GoogleBiglakeHiveDatabaseIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_biglake_hive_database_iam_member#description GoogleBiglakeHiveDatabaseIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

