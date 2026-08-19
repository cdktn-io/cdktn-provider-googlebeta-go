// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecDeploymentSpecKeepAliveProbeHttpGet struct {
	// Required. Specifies the path of the HTTP GET request (e.g., '"/is_busy"').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#path GoogleVertexAiReasoningEngine#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Optional. Specifies the port number on the container to which the request is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#port GoogleVertexAiReasoningEngine#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

