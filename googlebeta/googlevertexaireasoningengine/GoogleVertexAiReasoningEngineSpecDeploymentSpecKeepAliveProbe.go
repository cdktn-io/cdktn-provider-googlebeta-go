// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecDeploymentSpecKeepAliveProbe struct {
	// http_get block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_reasoning_engine#http_get GoogleVertexAiReasoningEngine#http_get}
	HttpGet *GoogleVertexAiReasoningEngineSpecDeploymentSpecKeepAliveProbeHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// Optional.
	//
	// Specifies the maximum duration (in seconds) to keep the instance alive via this probe. Can be a maximum of 3600 seconds (1 hour).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_reasoning_engine#max_seconds GoogleVertexAiReasoningEngine#max_seconds}
	MaxSeconds *float64 `field:"optional" json:"maxSeconds" yaml:"maxSeconds"`
}

