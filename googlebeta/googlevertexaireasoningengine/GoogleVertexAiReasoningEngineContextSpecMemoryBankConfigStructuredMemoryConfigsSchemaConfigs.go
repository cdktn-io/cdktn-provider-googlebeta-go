// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigStructuredMemoryConfigsSchemaConfigs struct {
	// Required. Unique ID identifying the memory schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_reasoning_engine#id GoogleVertexAiReasoningEngine#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Optional. The memory schema defined as an OpenAPI Schema Object JSON string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_vertex_ai_reasoning_engine#memory_schema GoogleVertexAiReasoningEngine#memory_schema}
	MemorySchema *string `field:"optional" json:"memorySchema" yaml:"memorySchema"`
}

