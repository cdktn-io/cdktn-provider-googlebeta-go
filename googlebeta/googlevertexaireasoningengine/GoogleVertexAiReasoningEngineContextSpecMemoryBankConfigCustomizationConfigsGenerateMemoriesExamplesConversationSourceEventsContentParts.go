// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentParts struct {
	// code_execution_result block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#code_execution_result GoogleVertexAiReasoningEngine#code_execution_result}
	CodeExecutionResult *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsCodeExecutionResult `field:"optional" json:"codeExecutionResult" yaml:"codeExecutionResult"`
	// executable_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#executable_code GoogleVertexAiReasoningEngine#executable_code}
	ExecutableCode *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsExecutableCode `field:"optional" json:"executableCode" yaml:"executableCode"`
	// file_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#file_data GoogleVertexAiReasoningEngine#file_data}
	FileData *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsFileData `field:"optional" json:"fileData" yaml:"fileData"`
	// function_call block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#function_call GoogleVertexAiReasoningEngine#function_call}
	FunctionCall *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsFunctionCall `field:"optional" json:"functionCall" yaml:"functionCall"`
	// function_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#function_response GoogleVertexAiReasoningEngine#function_response}
	FunctionResponse *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsFunctionResponse `field:"optional" json:"functionResponse" yaml:"functionResponse"`
	// inline_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#inline_data GoogleVertexAiReasoningEngine#inline_data}
	InlineData *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsInlineData `field:"optional" json:"inlineData" yaml:"inlineData"`
	// The text content of the part.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#text GoogleVertexAiReasoningEngine#text}
	Text *string `field:"optional" json:"text" yaml:"text"`
	// Indicates whether the part represents the model's thought process or reasoning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#thought GoogleVertexAiReasoningEngine#thought}
	Thought interface{} `field:"optional" json:"thought" yaml:"thought"`
	// video_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_vertex_ai_reasoning_engine#video_metadata GoogleVertexAiReasoningEngine#video_metadata}
	VideoMetadata *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsVideoMetadata `field:"optional" json:"videoMetadata" yaml:"videoMetadata"`
}

