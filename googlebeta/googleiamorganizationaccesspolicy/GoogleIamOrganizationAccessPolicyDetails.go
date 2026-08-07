// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamorganizationaccesspolicy


type GoogleIamOrganizationAccessPolicyDetails struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_iam_organization_access_policy#rules GoogleIamOrganizationAccessPolicy#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

