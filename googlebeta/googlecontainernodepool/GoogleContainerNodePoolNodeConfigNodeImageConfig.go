// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigNodeImageConfig struct {
	// The Operating System image for the node pool.
	//
	// This is a private feature, please contact your Google account team for allowlisting this feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_container_node_pool#image GoogleContainerNodePool#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// The GCP project storing the Operating System image for the node pool.
	//
	// This is a private feature, please contact your Google account team for allowlisting this feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_container_node_pool#image_project GoogleContainerNodePool#image_project}
	ImageProject *string `field:"optional" json:"imageProject" yaml:"imageProject"`
}

