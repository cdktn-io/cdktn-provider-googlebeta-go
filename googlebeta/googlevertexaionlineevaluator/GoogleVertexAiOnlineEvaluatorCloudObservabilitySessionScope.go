// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaionlineevaluator


type GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScope struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_online_evaluator#filter GoogleVertexAiOnlineEvaluator#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// inactivity_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_online_evaluator#inactivity_trigger GoogleVertexAiOnlineEvaluator#inactivity_trigger}
	InactivityTrigger *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeInactivityTrigger `field:"optional" json:"inactivityTrigger" yaml:"inactivityTrigger"`
}

