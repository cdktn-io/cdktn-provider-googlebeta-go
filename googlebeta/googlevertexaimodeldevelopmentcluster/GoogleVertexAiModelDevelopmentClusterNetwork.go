// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaimodeldevelopmentcluster


type GoogleVertexAiModelDevelopmentClusterNetwork struct {
	// The name of the VPC network where to create the cluster. The 'default' network is used if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#network GoogleVertexAiModelDevelopmentCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The name of the subnetwork in the above network to create the cluster.
	//
	// If not specified:
	// * for auto mode VPC, the subnet in corresponding region is auto-selected;
	// * for custom mode VPC, a new subnet will be automatically created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#subnetwork GoogleVertexAiModelDevelopmentCluster#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
}

