// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledialogflowgenerator


type GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResult struct {
	// The name of the tool's action associated with this call.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dialogflow_generator#action GoogleDialogflowGenerator#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// error block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dialogflow_generator#error GoogleDialogflowGenerator#error}
	Error *GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultError `field:"optional" json:"error" yaml:"error"`
}

