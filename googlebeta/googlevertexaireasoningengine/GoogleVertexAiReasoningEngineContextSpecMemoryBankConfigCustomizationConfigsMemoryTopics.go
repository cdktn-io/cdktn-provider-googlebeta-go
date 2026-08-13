// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopics struct {
	// custom_memory_topic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_reasoning_engine#custom_memory_topic GoogleVertexAiReasoningEngine#custom_memory_topic}
	CustomMemoryTopic *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsCustomMemoryTopic `field:"optional" json:"customMemoryTopic" yaml:"customMemoryTopic"`
	// managed_memory_topic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_reasoning_engine#managed_memory_topic GoogleVertexAiReasoningEngine#managed_memory_topic}
	ManagedMemoryTopic *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsMemoryTopicsManagedMemoryTopic `field:"optional" json:"managedMemoryTopic" yaml:"managedMemoryTopic"`
}

