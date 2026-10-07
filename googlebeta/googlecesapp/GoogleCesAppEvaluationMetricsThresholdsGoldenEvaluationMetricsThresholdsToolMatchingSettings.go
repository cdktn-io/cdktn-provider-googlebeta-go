// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesapp


type GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsToolMatchingSettings struct {
	// Defines the behavior when an extra tool call is encountered.
	//
	// An extra
	// tool call is a tool call that is present in the execution but does
	// not match any tool call in the golden expectation. Possible values: ["FAIL", "ALLOW"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_app#extra_tool_call_behavior GoogleCesApp#extra_tool_call_behavior}
	ExtraToolCallBehavior *string `field:"optional" json:"extraToolCallBehavior" yaml:"extraToolCallBehavior"`
}

