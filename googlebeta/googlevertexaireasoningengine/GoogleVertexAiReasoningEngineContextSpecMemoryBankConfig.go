// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfig struct {
	// customization_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#customization_configs GoogleVertexAiReasoningEngine#customization_configs}
	CustomizationConfigs interface{} `field:"optional" json:"customizationConfigs" yaml:"customizationConfigs"`
	// If true, no memory revisions will be created for any requests to the Memory Bank.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#disable_memory_revisions GoogleVertexAiReasoningEngine#disable_memory_revisions}
	DisableMemoryRevisions interface{} `field:"optional" json:"disableMemoryRevisions" yaml:"disableMemoryRevisions"`
	// generation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#generation_config GoogleVertexAiReasoningEngine#generation_config}
	GenerationConfig *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig `field:"optional" json:"generationConfig" yaml:"generationConfig"`
	// similarity_search_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#similarity_search_config GoogleVertexAiReasoningEngine#similarity_search_config}
	SimilaritySearchConfig *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig `field:"optional" json:"similaritySearchConfig" yaml:"similaritySearchConfig"`
	// structured_memory_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#structured_memory_configs GoogleVertexAiReasoningEngine#structured_memory_configs}
	StructuredMemoryConfigs interface{} `field:"optional" json:"structuredMemoryConfigs" yaml:"structuredMemoryConfigs"`
	// ttl_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#ttl_config GoogleVertexAiReasoningEngine#ttl_config}
	TtlConfig *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig `field:"optional" json:"ttlConfig" yaml:"ttlConfig"`
}

