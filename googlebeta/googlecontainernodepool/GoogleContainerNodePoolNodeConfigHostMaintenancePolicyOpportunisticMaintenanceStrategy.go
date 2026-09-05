// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigHostMaintenancePolicyOpportunisticMaintenanceStrategy struct {
	// The window of time that opportunistic maintenance can run.
	//
	// Example: A setting of 14 days implies that opportunistic maintenance can only be ran in the 2 weeks leading up to the scheduled maintenance date. Setting 28 days allows opportunistic maintenance to run at any time in the scheduled maintenance window (all PERIODIC maintenance is set 28 days in advance).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_container_node_pool#maintenance_availability_window GoogleContainerNodePool#maintenance_availability_window}
	MaintenanceAvailabilityWindow *string `field:"required" json:"maintenanceAvailabilityWindow" yaml:"maintenanceAvailabilityWindow"`
	// The minimum nodes required to be available in a pool.
	//
	// Blocks maintenance if it would cause the number of running nodes to dip below this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_container_node_pool#min_nodes_per_pool GoogleContainerNodePool#min_nodes_per_pool}
	MinNodesPerPool *float64 `field:"required" json:"minNodesPerPool" yaml:"minNodesPerPool"`
	// The amount of time that a node can remain idle (no customer owned workloads running), before triggering maintenance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_container_node_pool#node_idle_time_window GoogleContainerNodePool#node_idle_time_window}
	NodeIdleTimeWindow *string `field:"required" json:"nodeIdleTimeWindow" yaml:"nodeIdleTimeWindow"`
}

