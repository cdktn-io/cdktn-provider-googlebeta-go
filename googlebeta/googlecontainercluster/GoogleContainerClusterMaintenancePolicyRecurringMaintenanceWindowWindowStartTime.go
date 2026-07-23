// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowWindowStartTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_container_cluster#hours GoogleContainerCluster#hours}.
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_container_cluster#minutes GoogleContainerCluster#minutes}.
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_container_cluster#seconds GoogleContainerCluster#seconds}.
	Seconds *float64 `field:"required" json:"seconds" yaml:"seconds"`
}

