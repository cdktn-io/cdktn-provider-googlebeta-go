// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigLinuxNodeConfigNodeKernelModuleLoading struct {
	// The policy for kernel module loading.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_container_node_pool#policy GoogleContainerNodePool#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

