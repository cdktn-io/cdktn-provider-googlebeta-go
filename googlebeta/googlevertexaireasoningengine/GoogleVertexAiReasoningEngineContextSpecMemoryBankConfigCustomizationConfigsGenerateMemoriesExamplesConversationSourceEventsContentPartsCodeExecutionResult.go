// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsCodeExecutionResult struct {
	// Outcome of the code execution. Possible values: ["OUTCOME_UNSPECIFIED", "OUTCOME_OK", "OUTCOME_FAILED", "OUTCOME_DEADLINE_EXCEEDED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_reasoning_engine#outcome GoogleVertexAiReasoningEngine#outcome}
	Outcome *string `field:"required" json:"outcome" yaml:"outcome"`
	// The identifier of the ExecutableCode part this result is for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_reasoning_engine#id GoogleVertexAiReasoningEngine#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Contains stdout when code execution is successful, stderr or other description otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_vertex_ai_reasoning_engine#output GoogleVertexAiReasoningEngine#output}
	Output *string `field:"optional" json:"output" yaml:"output"`
}

