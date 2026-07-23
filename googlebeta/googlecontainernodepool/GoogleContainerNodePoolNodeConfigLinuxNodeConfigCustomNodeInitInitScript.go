// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScript struct {
	// The Secret Manager secret URI of the init script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_container_node_pool#gcp_secret_manager_secret_uri GoogleContainerNodePool#gcp_secret_manager_secret_uri}
	GcpSecretManagerSecretUri *string `field:"optional" json:"gcpSecretManagerSecretUri" yaml:"gcpSecretManagerSecretUri"`
	// The GCS generation of the init script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_container_node_pool#gcs_generation GoogleContainerNodePool#gcs_generation}
	GcsGeneration *float64 `field:"optional" json:"gcsGeneration" yaml:"gcsGeneration"`
	// The GCS URI of the init script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_container_node_pool#gcs_uri GoogleContainerNodePool#gcs_uri}
	GcsUri *string `field:"optional" json:"gcsUri" yaml:"gcsUri"`
}

