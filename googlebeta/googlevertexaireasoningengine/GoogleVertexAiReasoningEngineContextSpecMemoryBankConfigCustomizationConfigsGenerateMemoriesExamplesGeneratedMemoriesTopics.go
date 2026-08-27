// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesGeneratedMemoriesTopics struct {
	// Represents the custom memory topic label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#custom_memory_topic_label GoogleVertexAiReasoningEngine#custom_memory_topic_label}
	CustomMemoryTopicLabel *string `field:"optional" json:"customMemoryTopicLabel" yaml:"customMemoryTopicLabel"`
	// Represents the managed memory topic. Possible values: ["USER_PERSONAL_INFO", "USER_PREFERENCES", "KEY_CONVERSATION_DETAILS", "EXPLICIT_INSTRUCTIONS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#managed_memory_topic GoogleVertexAiReasoningEngine#managed_memory_topic}
	ManagedMemoryTopic *string `field:"optional" json:"managedMemoryTopic" yaml:"managedMemoryTopic"`
}

