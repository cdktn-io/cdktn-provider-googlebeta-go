// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledialogflowgenerator


type GoogleDialogflowGeneratorSummarizationContextFewShotExamplesOutputToolCallInfoToolCall struct {
	// The name of the tool's action associated with this call.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dialogflow_generator#action GoogleDialogflowGenerator#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The tool associated with this call.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_dialogflow_generator#tool GoogleDialogflowGenerator#tool}
	Tool *string `field:"optional" json:"tool" yaml:"tool"`
}

