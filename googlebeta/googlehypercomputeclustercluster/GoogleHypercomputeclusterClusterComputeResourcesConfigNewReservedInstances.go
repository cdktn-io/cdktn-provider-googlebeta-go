// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterComputeResourcesConfigNewReservedInstances struct {
	// Name of the reservation from which VM instances should be created, in the format 'projects/{project}/zones/{zone}/reservations/{reservation}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_hypercomputecluster_cluster#reservation GoogleHypercomputeclusterCluster#reservation}
	Reservation *string `field:"optional" json:"reservation" yaml:"reservation"`
}

