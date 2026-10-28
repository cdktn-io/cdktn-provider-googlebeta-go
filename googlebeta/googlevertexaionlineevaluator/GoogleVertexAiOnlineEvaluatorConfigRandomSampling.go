// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaionlineevaluator


type GoogleVertexAiOnlineEvaluatorConfigRandomSampling struct {
	// The percentage of traces to sample for evaluation. Must be an integer between 1 and 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_online_evaluator#percentage GoogleVertexAiOnlineEvaluator#percentage}
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
}

