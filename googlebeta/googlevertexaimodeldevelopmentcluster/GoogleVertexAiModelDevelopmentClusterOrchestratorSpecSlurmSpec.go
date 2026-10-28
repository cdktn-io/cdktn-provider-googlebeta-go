// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaimodeldevelopmentcluster


type GoogleVertexAiModelDevelopmentClusterOrchestratorSpecSlurmSpec struct {
	// Full name of the Filestore to be mounted as /home directory on every node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#home_directory_storage GoogleVertexAiModelDevelopmentCluster#home_directory_storage}
	HomeDirectoryStorage *string `field:"required" json:"homeDirectoryStorage" yaml:"homeDirectoryStorage"`
	// The node pool id to be used as slurm login node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#login_node_pool_id GoogleVertexAiModelDevelopmentCluster#login_node_pool_id}
	LoginNodePoolId *string `field:"required" json:"loginNodePoolId" yaml:"loginNodePoolId"`
	// Cluster-scoped Slurm settings, emitted as global 'Key=Value' lines in slurm.conf.
	//
	// Keys are Slurm parameter names exactly as documented by SchedMD
	// (https://slurm.schedmd.com/slurm.conf.html), matched case-insensitively.
	// Values are written the way slurm.conf writes them, including composite
	// forms such as 'bf_window=1440,bf_resolution=600'.
	//
	// Example:
	// { "PriorityType":            "priority/multifactor",
	// "PriorityWeightAge":       "1000",
	// "PreemptMode":             "SUSPEND,GANG",
	// "SchedulerParameters":     "bf_window=1440,bf_resolution=600" }
	//
	// Unsupported, reserved, or misspelled parameters, and values that are not
	// legal for the parameter, are rejected at request time with the offending
	// parameter named; they are never silently dropped. Settings
	// that apply to a single partition or node pool belong in
	// 'SlurmPartition.config' or 'SlurmNodeSet.config' instead -
	// several parameters (e.g. 'DefMemPerCPU', 'OverTimeLimit', 'PreemptMode')
	// are valid at more than one scope with different meanings and different
	// legal values.
	//
	// On update this map is replaced wholesale; supply the complete desired set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#config GoogleVertexAiModelDevelopmentCluster#config}
	Config *map[string]*string `field:"optional" json:"config" yaml:"config"`
	// Id of the partition that Slurm jobs submitted without an explicit partition run on, for example 'sbatch' without '--partition'.
	//
	// Must match
	// the 'id' of one of 'partitions'.
	//
	// If unset, the service uses the first entry of 'partitions' and this field
	// stays empty; the cluster always has a default partition whether or not this
	// field is set. Setting it makes the choice independent of list order, which
	// matters because 'partitions' is replaced wholesale on update: reordering
	// the list, or prepending a partition to it, otherwise moves the default and
	// silently redirects every job submitted without '--partition'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#default_partition_id GoogleVertexAiModelDevelopmentCluster#default_partition_id}
	DefaultPartitionId *string `field:"optional" json:"defaultPartitionId" yaml:"defaultPartitionId"`
	// Custom bash scripts executed after each slurm job on every allocated node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#epilog_bash_scripts GoogleVertexAiModelDevelopmentCluster#epilog_bash_scripts}
	EpilogBashScripts *[]*string `field:"optional" json:"epilogBashScripts" yaml:"epilogBashScripts"`
	// node_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#node_sets GoogleVertexAiModelDevelopmentCluster#node_sets}
	NodeSets interface{} `field:"optional" json:"nodeSets" yaml:"nodeSets"`
	// partitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#partitions GoogleVertexAiModelDevelopmentCluster#partitions}
	Partitions interface{} `field:"optional" json:"partitions" yaml:"partitions"`
	// Custom bash scripts executed before each slurm job on every allocated node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_model_development_cluster#prolog_bash_scripts GoogleVertexAiModelDevelopmentCluster#prolog_bash_scripts}
	PrologBashScripts *[]*string `field:"optional" json:"prologBashScripts" yaml:"prologBashScripts"`
}

