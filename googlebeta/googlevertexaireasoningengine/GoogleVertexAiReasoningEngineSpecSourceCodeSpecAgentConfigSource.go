// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSource struct {
	// adk_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_vertex_ai_reasoning_engine#adk_config GoogleVertexAiReasoningEngine#adk_config}
	AdkConfig *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceAdkConfig `field:"optional" json:"adkConfig" yaml:"adkConfig"`
	// inline_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_vertex_ai_reasoning_engine#inline_source GoogleVertexAiReasoningEngine#inline_source}
	InlineSource *GoogleVertexAiReasoningEngineSpecSourceCodeSpecAgentConfigSourceInlineSource `field:"optional" json:"inlineSource" yaml:"inlineSource"`
}

