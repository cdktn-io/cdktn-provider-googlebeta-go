// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterNetworkResourcesConfigExistingNetwork struct {
	// Name of the network to import, in the format 'projects/{project}/global/networks/{network}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_hypercomputecluster_cluster#network GoogleHypercomputeclusterCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Particular subnetwork to use, in the format 'projects/{project}/regions/{region}/subnetworks/{subnetwork}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_hypercomputecluster_cluster#subnetwork GoogleHypercomputeclusterCluster#subnetwork}
	Subnetwork *string `field:"required" json:"subnetwork" yaml:"subnetwork"`
}

