// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleaccesscontextmanagergcpuseraccessbinding


type GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsScopeClientScopeRestrictedProject struct {
	// The Google Cloud project resource name. Format: 'projects/{project_number}'. Only the project number is supported. Example: 'projects/1234567890'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_access_context_manager_gcp_user_access_binding#name GoogleAccessContextManagerGcpUserAccessBinding#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

