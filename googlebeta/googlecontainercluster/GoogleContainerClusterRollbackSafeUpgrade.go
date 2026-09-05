// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterRollbackSafeUpgrade struct {
	// A user-defined period that the cluster remains in the rollbackable state.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "604800s" for 7 days. Minimum is 6 hours, maximum is 7 days. If omitted, the two-step upgrade is skipped and a standard one-step upgrade is performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_container_cluster#control_plane_soak_duration GoogleContainerCluster#control_plane_soak_duration}
	ControlPlaneSoakDuration *string `field:"optional" json:"controlPlaneSoakDuration" yaml:"controlPlaneSoakDuration"`
}

