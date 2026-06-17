// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestool


type GoogleCesToolWidgetToolTextResponseConfig struct {
	// Optional. The static text response to return when type is STATIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_ces_tool#static_text GoogleCesTool#static_text}
	StaticText *string `field:"optional" json:"staticText" yaml:"staticText"`
	// Optional.
	//
	// Instruction for the LLM on how to generate the text response. Used as
	// the description for the text response parameter if type is LLM_GENERATED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_ces_tool#text_response_instruction GoogleCesTool#text_response_instruction}
	TextResponseInstruction *string `field:"optional" json:"textResponseInstruction" yaml:"textResponseInstruction"`
	// Optional. The strategy for providing the text response. Possible values: TYPE_UNSPECIFIED NONE LLM_GENERATED STATIC Possible values: ["TYPE_UNSPECIFIED", "NONE", "LLM_GENERATED", "STATIC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.37.0/docs/resources/google_ces_tool#type GoogleCesTool#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

