// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigAgentToAnywhereConfig struct {
	// Required. The resource name of the Agent Gateway for outbound traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_reasoning_engine#agent_gateway GoogleVertexAiReasoningEngine#agent_gateway}
	AgentGateway *string `field:"required" json:"agentGateway" yaml:"agentGateway"`
}

