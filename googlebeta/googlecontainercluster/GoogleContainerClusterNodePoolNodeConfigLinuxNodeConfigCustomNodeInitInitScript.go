// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigCustomNodeInitInitScript struct {
	// The Secret Manager secret URI of the init script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_cluster#gcp_secret_manager_secret_uri GoogleContainerCluster#gcp_secret_manager_secret_uri}
	GcpSecretManagerSecretUri *string `field:"optional" json:"gcpSecretManagerSecretUri" yaml:"gcpSecretManagerSecretUri"`
	// The GCS generation of the init script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_cluster#gcs_generation GoogleContainerCluster#gcs_generation}
	GcsGeneration *float64 `field:"optional" json:"gcsGeneration" yaml:"gcsGeneration"`
	// The GCS URI of the init script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_cluster#gcs_uri GoogleContainerCluster#gcs_uri}
	GcsUri *string `field:"optional" json:"gcsUri" yaml:"gcsUri"`
}

