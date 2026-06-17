// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigNodeImageConfig struct {
	// The name of the image to use for this node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_container_node_pool#image GoogleContainerNodePool#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// The project containing the image to use for this node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_container_node_pool#image_project GoogleContainerNodePool#image_project}
	ImageProject *string `field:"optional" json:"imageProject" yaml:"imageProject"`
}

