// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamfolderaccesspolicy


type GoogleIamFolderAccessPolicyDetailsRulesConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_iam_folder_access_policy#service GoogleIamFolderAccessPolicy#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_iam_folder_access_policy#expression GoogleIamFolderAccessPolicy#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

