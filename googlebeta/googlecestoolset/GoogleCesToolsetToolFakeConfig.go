// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestoolset


type GoogleCesToolsetToolFakeConfig struct {
	// code_block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_toolset#code_block GoogleCesToolset#code_block}
	CodeBlock *GoogleCesToolsetToolFakeConfigCodeBlock `field:"optional" json:"codeBlock" yaml:"codeBlock"`
	// Whether the tool is using fake mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_toolset#enable_fake_mode GoogleCesToolset#enable_fake_mode}
	EnableFakeMode interface{} `field:"optional" json:"enableFakeMode" yaml:"enableFakeMode"`
}

