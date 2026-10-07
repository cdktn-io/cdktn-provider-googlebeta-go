// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamprojectspolicybinding


type GoogleIamProjectsPolicyBindingTarget struct {
	// Immutable.
	//
	// Full Resource Name of the principal set used for principal access boundary policy bindings.
	// Examples for each one of the following supported principal set types:
	// * Project:
	//   * '//cloudresourcemanager.googleapis.com/projects/PROJECT_NUMBER'
	//   * '//cloudresourcemanager.googleapis.com/projects/PROJECT_ID'
	// * Workload Identity Pool: '//iam.googleapis.com/projects/PROJECT_NUMBER/locations/LOCATION/workloadIdentityPools/WORKLOAD_POOL_ID'
	// It must be parent by the policy binding's parent (the project).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_iam_projects_policy_binding#principal_set GoogleIamProjectsPolicyBinding#principal_set}
	PrincipalSet *string `field:"optional" json:"principalSet" yaml:"principalSet"`
	// Immutable.
	//
	// Full Resource Name of the resource used for access policy bindings.
	// Use this together with 'policy_kind = "ACCESS"'. Examples:
	// * Project:
	//   * '//cloudresourcemanager.googleapis.com/projects/PROJECT_NUMBER'
	//   * '//cloudresourcemanager.googleapis.com/projects/PROJECT_ID'
	// It must be the policy binding's parent (the project).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_iam_projects_policy_binding#resource GoogleIamProjectsPolicyBinding#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

