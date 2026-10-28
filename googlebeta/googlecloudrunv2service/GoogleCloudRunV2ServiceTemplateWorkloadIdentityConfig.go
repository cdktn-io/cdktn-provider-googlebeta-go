// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateWorkloadIdentityConfig struct {
	// The Revision's SPIFFE workload identity. Enables provisioning of SPIFFE workload certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_cloud_run_v2_service#identity GoogleCloudRunV2Service#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// Controls whether an instance receives a managed workload identity (MWLID) certificate. Defaults to true when 'identity_type' is 'IDENTITY_TYPE_AGENT_IDENTITY'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_cloud_run_v2_service#identity_certificate_enabled GoogleCloudRunV2Service#identity_certificate_enabled}
	IdentityCertificateEnabled interface{} `field:"optional" json:"identityCertificateEnabled" yaml:"identityCertificateEnabled"`
	// The type of identity to use.
	//
	// 'IDENTITY_TYPE_AGENT_IDENTITY' assigns a system-managed agent identity and
	// is required when the Service's 'functional_type' is 'FUNCTIONAL_TYPE_AGENT'. Once set, this field
	// cannot be changed or unset. Possible values: ["IDENTITY_TYPE_SERVICE_ACCOUNT", "IDENTITY_TYPE_WORKLOAD_IDENTITY", "IDENTITY_TYPE_AGENT_IDENTITY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_cloud_run_v2_service#identity_type GoogleCloudRunV2Service#identity_type}
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
}

