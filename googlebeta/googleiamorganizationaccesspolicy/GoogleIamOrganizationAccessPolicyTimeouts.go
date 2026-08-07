// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamorganizationaccesspolicy


type GoogleIamOrganizationAccessPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_iam_organization_access_policy#create GoogleIamOrganizationAccessPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_iam_organization_access_policy#delete GoogleIamOrganizationAccessPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_iam_organization_access_policy#update GoogleIamOrganizationAccessPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

