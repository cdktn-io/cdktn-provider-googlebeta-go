// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiamorganizationaccesspolicy


type GoogleIamOrganizationAccessPolicyDetailsRulesOperation struct {
	// The permissions that are explicitly affected by this rule.
	//
	// Each
	// permission uses the format '{service_fqdn}/{resource}.{verb}', where
	// '{service_fqdn}' is the fully qualified domain name for the service.
	// Currently supported permissions are as follows:
	// * 'eventarc.googleapis.com/messageBuses.publish'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_iam_organization_access_policy#permissions GoogleIamOrganizationAccessPolicy#permissions}
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// Specifies the permissions that this rule excludes from the set of affected permissions given by 'permissions'.
	//
	// If a permission appears in
	// 'permissions' _and_ in 'excluded_permissions' then it will _not_ be
	// subject to the policy effect.
	// The excluded permissions can be specified using the same syntax as
	// 'permissions'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_iam_organization_access_policy#excluded_permissions GoogleIamOrganizationAccessPolicy#excluded_permissions}
	ExcludedPermissions *[]*string `field:"optional" json:"excludedPermissions" yaml:"excludedPermissions"`
}

