// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaionlineevaluator


type GoogleVertexAiOnlineEvaluatorConfigA struct {
	// The maximum number of evaluations to perform per run. If set to 0, the number is unbounded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#max_evaluated_samples_per_run GoogleVertexAiOnlineEvaluator#max_evaluated_samples_per_run}
	MaxEvaluatedSamplesPerRun *string `field:"optional" json:"maxEvaluatedSamplesPerRun" yaml:"maxEvaluatedSamplesPerRun"`
	// random_sampling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#random_sampling GoogleVertexAiOnlineEvaluator#random_sampling}
	RandomSampling *GoogleVertexAiOnlineEvaluatorConfigRandomSampling `field:"optional" json:"randomSampling" yaml:"randomSampling"`
}

