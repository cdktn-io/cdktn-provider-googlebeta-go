// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigs struct {
	// consolidation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_reasoning_engine#consolidation_config GoogleVertexAiReasoningEngine#consolidation_config}
	ConsolidationConfig *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsConsolidationConfig `field:"optional" json:"consolidationConfig" yaml:"consolidationConfig"`
	// Optional. Generate memories in the third person if set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_reasoning_engine#enable_third_person_memories GoogleVertexAiReasoningEngine#enable_third_person_memories}
	EnableThirdPersonMemories interface{} `field:"optional" json:"enableThirdPersonMemories" yaml:"enableThirdPersonMemories"`
	// memory_topics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_reasoning_engine#memory_topics GoogleVertexAiReasoningEngine#memory_topics}
	MemoryTopics interface{} `field:"optional" json:"memoryTopics" yaml:"memoryTopics"`
	// Optional. List of scope keys that this customization config applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_reasoning_engine#scope_keys GoogleVertexAiReasoningEngine#scope_keys}
	ScopeKeys *[]*string `field:"optional" json:"scopeKeys" yaml:"scopeKeys"`
}

