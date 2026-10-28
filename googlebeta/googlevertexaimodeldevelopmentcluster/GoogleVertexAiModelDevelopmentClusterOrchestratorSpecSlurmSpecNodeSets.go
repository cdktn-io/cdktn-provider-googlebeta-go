// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaimodeldevelopmentcluster


type GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecNodeSets struct {
	// Id of the node pool these settings apply to.
	//
	// Must match the 'id' of one of
	// the node pools in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#node_pool_id GoogleVertexAiModelDevelopmentCluster#node_pool_id}
	NodePoolId *string `field:"required" json:"nodePoolId" yaml:"nodePoolId"`
	// Node-scoped Slurm settings, emitted as additional 'Key=Value' pairs on the 'NodeName=' line for this node pool's nodes in slurm.conf.
	//
	// Example:
	// { "Features":     "a3mega,nvlink",
	// "Weight":       "10",
	// "MemSpecLimit": "4096" }
	//
	// On update this map is replaced wholesale; supply the complete desired set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#config GoogleVertexAiModelDevelopmentCluster#config}
	Config *map[string]*string `field:"optional" json:"config" yaml:"config"`
}

