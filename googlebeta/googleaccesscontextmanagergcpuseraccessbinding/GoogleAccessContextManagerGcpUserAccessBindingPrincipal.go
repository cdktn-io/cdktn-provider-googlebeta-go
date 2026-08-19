// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleaccesscontextmanagergcpuseraccessbinding


type GoogleAccessContextManagerGcpUserAccessBindingPrincipal struct {
	// Immutable.
	//
	// Service account email used to assign policies to a single service account.
	// If a service account is subject to multiple policies (e.g., if there is a policy for all
	// service accounts in a project and a policy for the service account), the closest (i.e.
	// the most specific) dry-run policy will be used for the dry-run functionality and the
	// closest policy will be used for the enforcement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_access_context_manager_gcp_user_access_binding#service_account GoogleAccessContextManagerGcpUserAccessBinding#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Immutable. Cloud project number used to assign policies to all service accounts owned by the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_access_context_manager_gcp_user_access_binding#service_account_project_number GoogleAccessContextManagerGcpUserAccessBinding#service_account_project_number}
	ServiceAccountProjectNumber *string `field:"optional" json:"serviceAccountProjectNumber" yaml:"serviceAccountProjectNumber"`
}

