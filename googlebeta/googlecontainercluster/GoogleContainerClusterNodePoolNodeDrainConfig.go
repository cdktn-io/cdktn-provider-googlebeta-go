// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNodePoolNodeDrainConfig struct {
	// Whether to respect PodDisruptionBudget policy during node pool deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_container_cluster#respect_pdb_during_node_pool_deletion GoogleContainerCluster#respect_pdb_during_node_pool_deletion}
	RespectPdbDuringNodePoolDeletion interface{} `field:"optional" json:"respectPdbDuringNodePoolDeletion" yaml:"respectPdbDuringNodePoolDeletion"`
}

