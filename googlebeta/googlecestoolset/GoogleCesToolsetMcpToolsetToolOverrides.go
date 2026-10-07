// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestoolset


type GoogleCesToolsetMcpToolsetToolOverrides struct {
	// The name of the tool to be overridden.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_toolset#tool GoogleCesToolset#tool}
	Tool *string `field:"required" json:"tool" yaml:"tool"`
	// The description override for the tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_toolset#description_override GoogleCesToolset#description_override}
	DescriptionOverride *string `field:"optional" json:"descriptionOverride" yaml:"descriptionOverride"`
	// The name override for the tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_toolset#name_override GoogleCesToolset#name_override}
	NameOverride *string `field:"optional" json:"nameOverride" yaml:"nameOverride"`
}

