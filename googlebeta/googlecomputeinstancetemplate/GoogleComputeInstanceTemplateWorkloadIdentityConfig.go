// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeinstancetemplate


type GoogleComputeInstanceTemplateWorkloadIdentityConfig struct {
	// Identity SPIFFE id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_compute_instance_template#identity GoogleComputeInstanceTemplate#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// Specifies whether identity certificates are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_compute_instance_template#identity_certificate_enabled GoogleComputeInstanceTemplate#identity_certificate_enabled}
	IdentityCertificateEnabled interface{} `field:"optional" json:"identityCertificateEnabled" yaml:"identityCertificateEnabled"`
}

