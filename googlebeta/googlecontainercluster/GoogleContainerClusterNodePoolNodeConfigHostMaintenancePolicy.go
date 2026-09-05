// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigHostMaintenancePolicy struct {
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_container_cluster#maintenance_interval GoogleContainerCluster#maintenance_interval}
	MaintenanceInterval *string `field:"required" json:"maintenanceInterval" yaml:"maintenanceInterval"`
	// opportunistic_maintenance_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_container_cluster#opportunistic_maintenance_strategy GoogleContainerCluster#opportunistic_maintenance_strategy}
	OpportunisticMaintenanceStrategy *GoogleContainerClusterNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategy `field:"optional" json:"opportunisticMaintenanceStrategy" yaml:"opportunisticMaintenanceStrategy"`
}

