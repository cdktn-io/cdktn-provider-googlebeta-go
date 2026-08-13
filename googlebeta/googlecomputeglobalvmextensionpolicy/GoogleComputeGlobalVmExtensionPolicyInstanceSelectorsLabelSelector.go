// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeglobalvmextensionpolicy


type GoogleComputeGlobalVmExtensionPolicyInstanceSelectorsLabelSelector struct {
	// Labels as key value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_compute_global_vm_extension_policy#inclusion_labels GoogleComputeGlobalVmExtensionPolicy#inclusion_labels}
	InclusionLabels *map[string]*string `field:"optional" json:"inclusionLabels" yaml:"inclusionLabels"`
}

