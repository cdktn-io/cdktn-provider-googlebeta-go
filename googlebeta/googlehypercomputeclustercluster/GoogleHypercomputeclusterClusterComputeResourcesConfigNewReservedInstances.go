// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlehypercomputeclustercluster


type GoogleHypercomputeclusterClusterComputeResourcesConfigNewReservedInstances struct {
	// A list of network tags to attach to the VM instances. A maximum of 64 network tags are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_hypercomputecluster_cluster#network_tags GoogleHypercomputeclusterCluster#network_tags}
	NetworkTags *[]*string `field:"optional" json:"networkTags" yaml:"networkTags"`
	// Name of the reservation from which VM instances should be created, in the format 'projects/{project}/zones/{zone}/reservations/{reservation}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_hypercomputecluster_cluster#reservation GoogleHypercomputeclusterCluster#reservation}
	Reservation *string `field:"optional" json:"reservation" yaml:"reservation"`
}

