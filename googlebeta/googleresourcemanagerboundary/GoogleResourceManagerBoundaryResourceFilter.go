// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleresourcemanagerboundary


type GoogleResourceManagerBoundaryResourceFilter struct {
	// tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_resource_manager_boundary#tag_filter GoogleResourceManagerBoundary#tag_filter}
	TagFilter *GoogleResourceManagerBoundaryResourceFilterTagFilter `field:"required" json:"tagFilter" yaml:"tagFilter"`
}

