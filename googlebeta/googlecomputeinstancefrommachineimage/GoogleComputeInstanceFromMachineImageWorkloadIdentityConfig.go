// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeinstancefrommachineimage


type GoogleComputeInstanceFromMachineImageWorkloadIdentityConfig struct {
	// Identity SPIFFE id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_compute_instance_from_machine_image#identity GoogleComputeInstanceFromMachineImage#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// Specifies whether identity certificates are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_compute_instance_from_machine_image#identity_certificate_enabled GoogleComputeInstanceFromMachineImage#identity_certificate_enabled}
	IdentityCertificateEnabled interface{} `field:"optional" json:"identityCertificateEnabled" yaml:"identityCertificateEnabled"`
}

