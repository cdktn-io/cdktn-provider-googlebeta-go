// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleaccesscontextmanagergcpuseraccessbinding


type GoogleAccessContextManagerGcpUserAccessBindingPrincipal struct {
	// Immutable.
	//
	// The IAM principal identifier of the federated workforce or workload to
	// assign the policy to. Examples include the following:
	// * Single principal:
	// 'principal://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/subject/{subject_attribute_value}'
	// * All workloads in a workload identity pool:
	// 'principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/*'
	// * All Workforce Pools in a Google Cloud organization:
	// 'principalSet://cloudresourcemanager.googleapis.com/organizations/{organization_id}/type/WorkforcePool'
	//
	// Bindings created for all Workforce Pools in a Google Cloud organization
	// support only 'scoped_access_settings' with the 'restricted_project'
	// client scope and active 'session_settings'. No other configurations are
	// allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_access_context_manager_gcp_user_access_binding#federated_principal GoogleAccessContextManagerGcpUserAccessBinding#federated_principal}
	FederatedPrincipal *string `field:"optional" json:"federatedPrincipal" yaml:"federatedPrincipal"`
	// Immutable.
	//
	// Service account email used to assign policies to a single service account.
	// If a service account is subject to multiple policies (e.g., if there is a policy for all
	// service accounts in a project and a policy for the service account), the closest (i.e.
	// the most specific) dry-run policy will be used for the dry-run functionality and the
	// closest policy will be used for the enforcement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_access_context_manager_gcp_user_access_binding#service_account GoogleAccessContextManagerGcpUserAccessBinding#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Immutable. Cloud project number used to assign policies to all service accounts owned by the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_access_context_manager_gcp_user_access_binding#service_account_project_number GoogleAccessContextManagerGcpUserAccessBinding#service_account_project_number}
	ServiceAccountProjectNumber *string `field:"optional" json:"serviceAccountProjectNumber" yaml:"serviceAccountProjectNumber"`
}

