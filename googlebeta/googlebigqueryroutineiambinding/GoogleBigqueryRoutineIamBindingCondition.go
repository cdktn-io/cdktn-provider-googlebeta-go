// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryroutineiambinding


type GoogleBigqueryRoutineIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_bigquery_routine_iam_binding#expression GoogleBigqueryRoutineIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_bigquery_routine_iam_binding#title GoogleBigqueryRoutineIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.36.0/docs/resources/google_bigquery_routine_iam_binding#description GoogleBigqueryRoutineIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

