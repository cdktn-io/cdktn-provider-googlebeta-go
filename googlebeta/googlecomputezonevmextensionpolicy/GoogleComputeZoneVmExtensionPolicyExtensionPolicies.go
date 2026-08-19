// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputezonevmextensionpolicy


type GoogleComputeZoneVmExtensionPolicyExtensionPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_zone_vm_extension_policy#extension_name GoogleComputeZoneVmExtensionPolicy#extension_name}.
	ExtensionName *string `field:"required" json:"extensionName" yaml:"extensionName"`
	// The specific version of the extension to install.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_zone_vm_extension_policy#pinned_version GoogleComputeZoneVmExtensionPolicy#pinned_version}
	PinnedVersion *string `field:"optional" json:"pinnedVersion" yaml:"pinnedVersion"`
	// String-based configuration data for the extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_compute_zone_vm_extension_policy#string_config GoogleComputeZoneVmExtensionPolicy#string_config}
	StringConfig *string `field:"optional" json:"stringConfig" yaml:"stringConfig"`
}

