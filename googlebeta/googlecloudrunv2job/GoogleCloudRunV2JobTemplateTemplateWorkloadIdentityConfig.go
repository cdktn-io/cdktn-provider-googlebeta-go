// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecloudrunv2job


type GoogleCloudRunV2JobTemplateTemplateWorkloadIdentityConfig struct {
	// The Task's SPIFFE workload identity. Enables provisioning of SPIFFE workload certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_cloud_run_v2_job#identity GoogleCloudRunV2Job#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// Controls whether an instance receives a managed workload identity (MWLID) certificate. Defaults to true when 'identity_type' is 'IDENTITY_TYPE_AGENT_IDENTITY'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_cloud_run_v2_job#identity_certificate_enabled GoogleCloudRunV2Job#identity_certificate_enabled}
	IdentityCertificateEnabled interface{} `field:"optional" json:"identityCertificateEnabled" yaml:"identityCertificateEnabled"`
	// The type of identity to use.
	//
	// 'IDENTITY_TYPE_AGENT_IDENTITY' assigns a system-managed agent identity
	// and is required when the Job's 'functional_type' is 'FUNCTIONAL_TYPE_AGENT'. Once set, this field
	// cannot be changed or unset. Possible values: ["IDENTITY_TYPE_SERVICE_ACCOUNT", "IDENTITY_TYPE_AGENT_IDENTITY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_cloud_run_v2_job#identity_type GoogleCloudRunV2Job#identity_type}
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
}

