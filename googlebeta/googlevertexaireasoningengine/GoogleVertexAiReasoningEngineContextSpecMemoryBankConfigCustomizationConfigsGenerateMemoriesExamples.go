// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamples struct {
	// conversation_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#conversation_source GoogleVertexAiReasoningEngine#conversation_source}
	ConversationSource *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSource `field:"optional" json:"conversationSource" yaml:"conversationSource"`
	// generated_memories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#generated_memories GoogleVertexAiReasoningEngine#generated_memories}
	GeneratedMemories interface{} `field:"optional" json:"generatedMemories" yaml:"generatedMemories"`
}

