// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebiglakehivecatalogiammember


type GoogleBiglakeHiveCatalogIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_catalog_iam_member#expression GoogleBiglakeHiveCatalogIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_catalog_iam_member#title GoogleBiglakeHiveCatalogIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_biglake_hive_catalog_iam_member#description GoogleBiglakeHiveCatalogIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

