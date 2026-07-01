// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryroutineiammember


type GoogleBigqueryRoutineIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_bigquery_routine_iam_member#expression GoogleBigqueryRoutineIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_bigquery_routine_iam_member#title GoogleBigqueryRoutineIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_bigquery_routine_iam_member#description GoogleBigqueryRoutineIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

