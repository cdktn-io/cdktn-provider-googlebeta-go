// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaionlineevaluator


type GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilter struct {
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#duration GoogleVertexAiOnlineEvaluator#duration}
	Duration *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterDuration `field:"optional" json:"duration" yaml:"duration"`
	// total_token_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#total_token_usage GoogleVertexAiOnlineEvaluator#total_token_usage}
	TotalTokenUsage *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScopeFilterTotalTokenUsage `field:"optional" json:"totalTokenUsage" yaml:"totalTokenUsage"`
}

