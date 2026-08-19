// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfig struct {
	// agent_to_anywhere_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#agent_to_anywhere_config GoogleVertexAiReasoningEngine#agent_to_anywhere_config}
	AgentToAnywhereConfig *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigAgentToAnywhereConfig `field:"optional" json:"agentToAnywhereConfig" yaml:"agentToAnywhereConfig"`
	// client_to_agent_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#client_to_agent_config GoogleVertexAiReasoningEngine#client_to_agent_config}
	ClientToAgentConfig *GoogleVertexAiReasoningEngineSpecDeploymentSpecAgentGatewayConfigClientToAgentConfig `field:"optional" json:"clientToAgentConfig" yaml:"clientToAgentConfig"`
}

