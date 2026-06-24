// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamprojectaccesspolicy


type GoogleIamProjectAccessPolicyDetailsRulesConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_iam_project_access_policy#service GoogleIamProjectAccessPolicy#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_iam_project_access_policy#expression GoogleIamProjectAccessPolicy#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

