// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesagent


type GoogleCesAgentRemoteA2AAgentA2AConfig struct {
	// agent_card block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#agent_card GoogleCesAgent#agent_card}
	AgentCard *GoogleCesAgentRemoteA2AAgentA2AConfigAgentCard `field:"optional" json:"agentCard" yaml:"agentCard"`
	// Reference to the agent in the Agent Registry. Format: 'projects/{project}/locations/{location}/agents/{agent}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#agent_registry GoogleCesAgent#agent_registry}
	AgentRegistry *string `field:"optional" json:"agentRegistry" yaml:"agentRegistry"`
	// api_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#api_authentication GoogleCesAgent#api_authentication}
	ApiAuthentication *GoogleCesAgentRemoteA2AAgentA2AConfigApiAuthentication `field:"optional" json:"apiAuthentication" yaml:"apiAuthentication"`
	// If not empty, interactions with the remote A2A agent will use this context ID.
	//
	// This context_id field can refer to a session variable like
	// '$context.variables.order_agent_session_id'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#context_id GoogleCesAgent#context_id}
	ContextId *string `field:"optional" json:"contextId" yaml:"contextId"`
	// Mapping of input variable names of remote agent to GECX variable names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#input_variable_mapping GoogleCesAgent#input_variable_mapping}
	InputVariableMapping *map[string]*string `field:"optional" json:"inputVariableMapping" yaml:"inputVariableMapping"`
	// Mapping of output variable names of remote agent to GECX variable names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#output_variable_mapping GoogleCesAgent#output_variable_mapping}
	OutputVariableMapping *map[string]*string `field:"optional" json:"outputVariableMapping" yaml:"outputVariableMapping"`
	// Whether streaming is enabled for the remote agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#streaming_enabled GoogleCesAgent#streaming_enabled}
	StreamingEnabled interface{} `field:"optional" json:"streamingEnabled" yaml:"streamingEnabled"`
}

