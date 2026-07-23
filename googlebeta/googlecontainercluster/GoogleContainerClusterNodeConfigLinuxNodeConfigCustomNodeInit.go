// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInit struct {
	// init_script block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_container_cluster#init_script GoogleContainerCluster#init_script}
	InitScript *GoogleContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScript `field:"optional" json:"initScript" yaml:"initScript"`
}

