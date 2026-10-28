// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleresourcemanagerboundary


type GoogleResourceManagerBoundaryResourceFilterTagFilter struct {
	// The namespaced name of the tag key, in the form '{parent_id}/{tag_key_short_name}', for example '123456789012/env'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_resource_manager_boundary#tag_key GoogleResourceManagerBoundary#tag_key}
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// The allowed short name of the tag value, for example 'production'.
	//
	// A project matches this tag filter if its effective tag value for
	// the given 'tagKey' matches this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_resource_manager_boundary#tag_value GoogleResourceManagerBoundary#tag_value}
	TagValue *string `field:"required" json:"tagValue" yaml:"tagValue"`
}

