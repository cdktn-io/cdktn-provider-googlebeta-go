// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterMaintenancePolicyRecurringMaintenanceWindowDelayUntil struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_container_cluster#day GoogleContainerCluster#day}.
	Day *float64 `field:"required" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_container_cluster#month GoogleContainerCluster#month}.
	Month *float64 `field:"required" json:"month" yaml:"month"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_container_cluster#year GoogleContainerCluster#year}.
	Year *float64 `field:"required" json:"year" yaml:"year"`
}

