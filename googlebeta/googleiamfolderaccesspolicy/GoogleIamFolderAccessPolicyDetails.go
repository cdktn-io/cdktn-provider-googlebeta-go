// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamfolderaccesspolicy


type GoogleIamFolderAccessPolicyDetails struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_iam_folder_access_policy#rules GoogleIamFolderAccessPolicy#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

