// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNodePoolMaintenancePolicy struct {
	// exclusion_until_end_of_support block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_container_cluster#exclusion_until_end_of_support GoogleContainerCluster#exclusion_until_end_of_support}
	ExclusionUntilEndOfSupport interface{} `field:"optional" json:"exclusionUntilEndOfSupport" yaml:"exclusionUntilEndOfSupport"`
}

