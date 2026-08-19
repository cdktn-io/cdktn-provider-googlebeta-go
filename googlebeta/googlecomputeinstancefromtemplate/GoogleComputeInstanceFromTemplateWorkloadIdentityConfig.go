// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeinstancefromtemplate


type GoogleComputeInstanceFromTemplateWorkloadIdentityConfig struct {
	// Identity SPIFFE id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_instance_from_template#identity GoogleComputeInstanceFromTemplate#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// Specifies whether identity certificates are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_instance_from_template#identity_certificate_enabled GoogleComputeInstanceFromTemplate#identity_certificate_enabled}
	IdentityCertificateEnabled interface{} `field:"optional" json:"identityCertificateEnabled" yaml:"identityCertificateEnabled"`
}

