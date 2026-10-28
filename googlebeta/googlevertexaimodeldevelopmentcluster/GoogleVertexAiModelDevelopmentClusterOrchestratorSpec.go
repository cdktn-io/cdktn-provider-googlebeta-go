// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaimodeldevelopmentcluster


type GoogleVertexAiModelDevelopmentClusterOrchestratorSpec struct {
	// slurm_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#slurm_spec GoogleVertexAiModelDevelopmentCluster#slurm_spec}
	SlurmSpec *GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpec `field:"optional" json:"slurmSpec" yaml:"slurmSpec"`
}

