// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledataplexdataproductiammember


type GoogleDataplexDataProductIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_dataplex_data_product_iam_member#expression GoogleDataplexDataProductIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_dataplex_data_product_iam_member#title GoogleDataplexDataProductIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_dataplex_data_product_iam_member#description GoogleDataplexDataProductIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

