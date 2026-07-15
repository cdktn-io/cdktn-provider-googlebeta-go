// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigLinuxNodeConfigCustomNodeInit struct {
	// init_script block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_container_node_pool#init_script GoogleContainerNodePool#init_script}
	InitScript *GoogleContainerNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScript `field:"optional" json:"initScript" yaml:"initScript"`
}

