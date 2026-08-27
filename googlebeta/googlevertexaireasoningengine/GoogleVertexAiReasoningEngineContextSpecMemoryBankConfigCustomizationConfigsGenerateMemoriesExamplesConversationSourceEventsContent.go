// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContent struct {
	// parts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#parts GoogleVertexAiReasoningEngine#parts}
	Parts interface{} `field:"required" json:"parts" yaml:"parts"`
	// The producer of the content.
	//
	// Must be either 'user' or 'model'. If not set, the service will default to 'user'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#role GoogleVertexAiReasoningEngine#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

