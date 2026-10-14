// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaionlineevaluator


type GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScopeInactivityTrigger struct {
	// The amount of time that must pass with no new traces before a session is considered ready for evaluation.
	//
	// Must be a positive
	// duration no greater than 7 days (604800 seconds), e.g. "3600s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#threshold GoogleVertexAiOnlineEvaluator#threshold}
	Threshold *string `field:"required" json:"threshold" yaml:"threshold"`
}

