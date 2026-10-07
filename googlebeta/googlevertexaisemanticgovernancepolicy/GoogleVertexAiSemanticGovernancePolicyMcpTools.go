// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaisemanticgovernancepolicy


type GoogleVertexAiSemanticGovernancePolicyMcpTools struct {
	// The resource name of the McpServer in Agent Registry that is affected by this policy. Format: 'projects/{project}/locations/{location}/mcpServers/{mcpServer}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_semantic_governance_policy#mcp_server GoogleVertexAiSemanticGovernancePolicy#mcp_server}
	McpServer *string `field:"required" json:"mcpServer" yaml:"mcpServer"`
	// The resource names of the McpTools used by the Agent that is affected by this policy.
	//
	// At least one tool must be listed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_semantic_governance_policy#tools GoogleVertexAiSemanticGovernancePolicy#tools}
	Tools *[]*string `field:"required" json:"tools" yaml:"tools"`
}

