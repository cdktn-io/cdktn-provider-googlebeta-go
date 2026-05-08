// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig struct {
	// The model used to generate memories. Format: projects/{project}/locations/{location}/publishers/google/models/{model}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#model GoogleVertexAiReasoningEngine#model}
	Model *string `field:"required" json:"model" yaml:"model"`
}

