// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNodePoolNodeDrainConfig struct {
	// The duration of the grace termination period for node drain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_cluster#grace_termination_duration GoogleContainerCluster#grace_termination_duration}
	GraceTerminationDuration *string `field:"optional" json:"graceTerminationDuration" yaml:"graceTerminationDuration"`
	// The duration of the PDB timeout period for node drain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_cluster#pdb_timeout_duration GoogleContainerCluster#pdb_timeout_duration}
	PdbTimeoutDuration *string `field:"optional" json:"pdbTimeoutDuration" yaml:"pdbTimeoutDuration"`
	// Whether to respect PodDisruptionBudget policy during node pool deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_cluster#respect_pdb_during_node_pool_deletion GoogleContainerCluster#respect_pdb_during_node_pool_deletion}
	RespectPdbDuringNodePoolDeletion interface{} `field:"optional" json:"respectPdbDuringNodePoolDeletion" yaml:"respectPdbDuringNodePoolDeletion"`
}

