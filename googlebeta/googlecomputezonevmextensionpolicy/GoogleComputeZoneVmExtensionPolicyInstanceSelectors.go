// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputezonevmextensionpolicy


type GoogleComputeZoneVmExtensionPolicyInstanceSelectors struct {
	// label_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_compute_zone_vm_extension_policy#label_selector GoogleComputeZoneVmExtensionPolicy#label_selector}
	LabelSelector *GoogleComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
}

