// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamprojectaccesspolicy


type GoogleIamProjectAccessPolicyDetailsRules struct {
	// The effect of the rule. Possible values: DENY ALLOW Possible values: ["DENY", "ALLOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_iam_project_access_policy#effect GoogleIamProjectAccessPolicy#effect}
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_iam_project_access_policy#operation GoogleIamProjectAccessPolicy#operation}
	Operation *GoogleIamProjectAccessPolicyDetailsRulesOperation `field:"required" json:"operation" yaml:"operation"`
	// The identities for which this rule's effect governs using one or more permissions on Google Cloud resources.
	//
	// This field can contain the
	// following values:
	// * 'principal://goog/subject/{email_id}': A specific Google Account.
	// Includes Gmail, Cloud Identity, and Google Workspace user accounts. For
	// example, 'principal://goog/subject/alice@example.com'.
	// * 'principal://iam.googleapis.com/projects/-/serviceAccounts/{service_account_id}':
	// A Google Cloud service account. For example,
	// 'principal://iam.googleapis.com/projects/-/serviceAccounts/my-service-account@iam.gserviceaccount.com'.
	// * 'principalSet://goog/group/{group_id}': A Google group. For example,
	// 'principalSet://goog/group/admins@example.com'.
	// * 'principalSet://goog/cloudIdentityCustomerId/{customer_id}': All of the
	// principals associated with the specified Google Workspace or Cloud
	// Identity customer ID. For example,
	// 'principalSet://goog/cloudIdentityCustomerId/C01Abc35'.
	// If an identifier that was previously set on a policy is soft deleted, then
	// calls to read that policy will return the identifier with a deleted
	// prefix. Users cannot set identifiers with this syntax.
	// * 'deleted:principal://goog/subject/{email_id}?uid={uid}': A specific
	// Google Account that was deleted recently. For example,
	// 'deleted:principal://goog/subject/alice@example.com?uid=1234567890'. If
	// the Google Account is recovered, this identifier reverts to the standard
	// identifier for a Google Account.
	// * 'deleted:principalSet://goog/group/{group_id}?uid={uid}': A Google group
	// that was deleted recently. For example,
	// 'deleted:principalSet://goog/group/admins@example.com?uid=1234567890'.
	// If the Google group is restored, this identifier reverts to the standard
	// identifier for a Google group.
	// * 'deleted:principal://iam.googleapis.com/projects/-/serviceAccounts/{service_account_id}?uid={uid}':
	// A Google Cloud service account that was deleted recently. For example,
	// 'deleted:principal://iam.googleapis.com/projects/-/serviceAccounts/my-service-account@iam.gserviceaccount.com?uid=1234567890'.
	// If the service account is undeleted, this identifier reverts to the
	// standard identifier for a service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_iam_project_access_policy#principals GoogleIamProjectAccessPolicy#principals}
	Principals *[]*string `field:"required" json:"principals" yaml:"principals"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_iam_project_access_policy#conditions GoogleIamProjectAccessPolicy#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Customer specified description of the rule. Must be less than or equal to 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_iam_project_access_policy#description GoogleIamProjectAccessPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identities that are excluded from the access policy rule, even if they are listed in the 'principals'.
	//
	// For example, you could add a Google
	// group to the 'principals', then exclude specific users who belong to
	// that group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_iam_project_access_policy#excluded_principals GoogleIamProjectAccessPolicy#excluded_principals}
	ExcludedPrincipals *[]*string `field:"optional" json:"excludedPrincipals" yaml:"excludedPrincipals"`
}

