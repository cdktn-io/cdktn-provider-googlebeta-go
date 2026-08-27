// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNodeConfigTaintConfig struct {
	// Architecture taint behavior. Controls, how we apply taints based on the node architecture.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_cluster#architecture_taint_behavior GoogleContainerCluster#architecture_taint_behavior}
	ArchitectureTaintBehavior *string `field:"required" json:"architectureTaintBehavior" yaml:"architectureTaintBehavior"`
}

