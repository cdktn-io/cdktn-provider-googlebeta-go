// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaimodeldevelopmentcluster


type GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpecPartitions struct {
	// Id of the slurm partiton.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#id GoogleVertexAiModelDevelopmentCluster#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Node pools for the slurm partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#node_pool_ids GoogleVertexAiModelDevelopmentCluster#node_pool_ids}
	NodePoolIds *[]*string `field:"required" json:"nodePoolIds" yaml:"nodePoolIds"`
	// Partition-scoped Slurm settings, emitted as additional 'Key=Value' pairs on this partition's 'PartitionName=' line in slurm.conf.
	//
	// Example:
	// { "MaxTime":       "4-00:00:00",
	// "DenyAccounts":  "low-pri",
	// "PriorityTier":  "20",
	// "OverSubscribe": "YES:4" }
	//
	// Note: values set here take precedence over service-computed defaults for
	// the same parameter (for example 'OverSubscribe' and 'DefMemPerCPU', which
	// the service otherwise derives from the partition's exclusivity and from
	// the memory per CPU of its nodes).
	//
	// 'Default' is rejected here: which partition is the default is a
	// cluster-scoped choice, so it belongs in 'SlurmSpec.default_partition_id'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#config GoogleVertexAiModelDevelopmentCluster#config}
	Config *map[string]*string `field:"optional" json:"config" yaml:"config"`
}

