// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecontainernodepool


type GoogleContainerNodePoolMaintenancePolicy struct {
	// exclusion_until_end_of_support block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_container_node_pool#exclusion_until_end_of_support GoogleContainerNodePool#exclusion_until_end_of_support}
	ExclusionUntilEndOfSupport interface{} `field:"optional" json:"exclusionUntilEndOfSupport" yaml:"exclusionUntilEndOfSupport"`
}

