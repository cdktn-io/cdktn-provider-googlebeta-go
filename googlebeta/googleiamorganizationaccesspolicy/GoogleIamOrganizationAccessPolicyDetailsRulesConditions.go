// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamorganizationaccesspolicy


type GoogleIamOrganizationAccessPolicyDetailsRulesConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_iam_organization_access_policy#service GoogleIamOrganizationAccessPolicy#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_iam_organization_access_policy#expression GoogleIamOrganizationAccessPolicy#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

