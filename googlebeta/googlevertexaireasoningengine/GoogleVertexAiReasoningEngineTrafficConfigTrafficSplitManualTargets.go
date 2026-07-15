// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineTrafficConfigTrafficSplitManualTargets struct {
	// Required. Specifies percent of the traffic to this Runtime Revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_reasoning_engine#percent GoogleVertexAiReasoningEngine#percent}
	Percent *float64 `field:"required" json:"percent" yaml:"percent"`
	// Required. The Runtime Revision name to which to send this portion of traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_vertex_ai_reasoning_engine#runtime_revision_name GoogleVertexAiReasoningEngine#runtime_revision_name}
	RuntimeRevisionName *string `field:"required" json:"runtimeRevisionName" yaml:"runtimeRevisionName"`
}

