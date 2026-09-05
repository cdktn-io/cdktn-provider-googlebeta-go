// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesGeneratedMemories struct {
	// Represents the fact to generate a memory from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_reasoning_engine#fact GoogleVertexAiReasoningEngine#fact}
	Fact *string `field:"required" json:"fact" yaml:"fact"`
	// topics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_reasoning_engine#topics GoogleVertexAiReasoningEngine#topics}
	Topics interface{} `field:"optional" json:"topics" yaml:"topics"`
}

