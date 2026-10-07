// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledialogflowgenerator


type GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutput struct {
	// summary_suggestion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dialogflow_generator#summary_suggestion GoogleDialogflowGenerator#summary_suggestion}
	SummarySuggestion *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputSummarySuggestion `field:"optional" json:"summarySuggestion" yaml:"summarySuggestion"`
	// tool_call_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dialogflow_generator#tool_call_info GoogleDialogflowGenerator#tool_call_info}
	ToolCallInfo interface{} `field:"optional" json:"toolCallInfo" yaml:"toolCallInfo"`
}

