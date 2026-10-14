// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaionlineevaluator


type GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilter struct {
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#duration GoogleVertexAiOnlineEvaluator#duration}
	Duration *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterDuration `field:"optional" json:"duration" yaml:"duration"`
	// model_call_errors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#model_call_errors GoogleVertexAiOnlineEvaluator#model_call_errors}
	ModelCallErrors *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCallErrors `field:"optional" json:"modelCallErrors" yaml:"modelCallErrors"`
	// model_calls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#model_calls GoogleVertexAiOnlineEvaluator#model_calls}
	ModelCalls *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterModelCalls `field:"optional" json:"modelCalls" yaml:"modelCalls"`
	// tool_call_errors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#tool_call_errors GoogleVertexAiOnlineEvaluator#tool_call_errors}
	ToolCallErrors *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCallErrors `field:"optional" json:"toolCallErrors" yaml:"toolCallErrors"`
	// tool_calls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#tool_calls GoogleVertexAiOnlineEvaluator#tool_calls}
	ToolCalls *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterToolCalls `field:"optional" json:"toolCalls" yaml:"toolCalls"`
	// total_token_usage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#total_token_usage GoogleVertexAiOnlineEvaluator#total_token_usage}
	TotalTokenUsage *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterTotalTokenUsage `field:"optional" json:"totalTokenUsage" yaml:"totalTokenUsage"`
	// user_turns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#user_turns GoogleVertexAiOnlineEvaluator#user_turns}
	UserTurns *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeFilterUserTurns `field:"optional" json:"userTurns" yaml:"userTurns"`
}

