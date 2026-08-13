// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlemodelarmortemplate


type GoogleModelArmorTemplateTemplateMetadataFilterVersionSelector struct {
	// A predefined filter version alias. The template automatically follows the version this alias points to. Possible values: FILTER_VERSION_ALIAS_STABLE FILTER_VERSION_ALIAS_LATEST.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_model_armor_template#alias GoogleModelArmorTemplate#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Pins the template to a specific, immutable filter version. Expected format is a case-sensitive string such as 'v1' or 'v2'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_model_armor_template#version GoogleModelArmorTemplate#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

