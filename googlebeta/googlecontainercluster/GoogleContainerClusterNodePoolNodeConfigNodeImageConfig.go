// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigNodeImageConfig struct {
	// The name of the image to use for this node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_container_cluster#image GoogleContainerCluster#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// The project containing the image to use for this node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_container_cluster#image_project GoogleContainerCluster#image_project}
	ImageProject *string `field:"optional" json:"imageProject" yaml:"imageProject"`
}

