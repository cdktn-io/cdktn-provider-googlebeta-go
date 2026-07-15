// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_container_cluster#recurrence GoogleContainerCluster#recurrence}.
	Recurrence *string `field:"required" json:"recurrence" yaml:"recurrence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_container_cluster#window_duration GoogleContainerCluster#window_duration}.
	WindowDuration *string `field:"required" json:"windowDuration" yaml:"windowDuration"`
	// window_start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_container_cluster#window_start_time GoogleContainerCluster#window_start_time}
	WindowStartTime *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTime `field:"required" json:"windowStartTime" yaml:"windowStartTime"`
	// delay_until block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_container_cluster#delay_until GoogleContainerCluster#delay_until}
	DelayUntil *GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntil `field:"optional" json:"delayUntil" yaml:"delayUntil"`
}

