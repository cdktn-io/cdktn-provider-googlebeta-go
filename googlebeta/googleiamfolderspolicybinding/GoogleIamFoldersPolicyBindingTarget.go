// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamfolderspolicybinding


type GoogleIamFoldersPolicyBindingTarget struct {
	// Immutable.
	//
	// Full Resource Name of the principal set used for principal access boundary policy bindings.
	// Examples for each one of the following supported principal set types:
	// * Folder: '//cloudresourcemanager.googleapis.com/folders/FOLDER_ID'
	// It must be parent by the policy binding's parent (the folder).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_iam_folders_policy_binding#principal_set GoogleIamFoldersPolicyBinding#principal_set}
	PrincipalSet *string `field:"optional" json:"principalSet" yaml:"principalSet"`
	// Immutable.
	//
	// Full Resource Name of the resource used for access policy bindings.
	// Use this together with 'policy_kind = "ACCESS"'. Examples:
	// * Folder: '//cloudresourcemanager.googleapis.com/folders/FOLDER_ID'
	// It must be the policy binding's parent (the folder).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_iam_folders_policy_binding#resource GoogleIamFoldersPolicyBinding#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

