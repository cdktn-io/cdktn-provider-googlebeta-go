// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigHostMaintenancePolicy struct {
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#maintenance_interval GoogleContainerNodePool#maintenance_interval}
	MaintenanceInterval *string `field:"required" json:"maintenanceInterval" yaml:"maintenanceInterval"`
	// opportunistic_maintenance_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_container_node_pool#opportunistic_maintenance_strategy GoogleContainerNodePool#opportunistic_maintenance_strategy}
	OpportunisticMaintenanceStrategy *GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategy `field:"optional" json:"opportunisticMaintenanceStrategy" yaml:"opportunisticMaintenanceStrategy"`
}

