// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledialogflowgenerator


type GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCallResultError struct {
	// The error message of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dialogflow_generator#message GoogleDialogflowGenerator#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Specifies whether the tool call is retryable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dialogflow_generator#retryable GoogleDialogflowGenerator#retryable}
	Retryable interface{} `field:"optional" json:"retryable" yaml:"retryable"`
}

