// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterOrchestrator struct {
	// slurm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_hypercomputecluster_cluster#slurm GoogleHypercomputeclusterCluster#slurm}
	Slurm *GoogleHypercomputeclusterClusterOrchestratorSlurm `field:"optional" json:"slurm" yaml:"slurm"`
}

