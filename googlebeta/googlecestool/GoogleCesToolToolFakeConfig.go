// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestool


type GoogleCesToolToolFakeConfig struct {
	// code_block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#code_block GoogleCesTool#code_block}
	CodeBlock *GoogleCesToolToolFakeConfigCodeBlock `field:"optional" json:"codeBlock" yaml:"codeBlock"`
	// Whether the tool is using fake mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_ces_tool#enable_fake_mode GoogleCesTool#enable_fake_mode}
	EnableFakeMode interface{} `field:"optional" json:"enableFakeMode" yaml:"enableFakeMode"`
}

