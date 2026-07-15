// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagentregistryservice


type GoogleAgentRegistryServiceMcpServerSpec struct {
	// The type of the MCP Server spec content. Possible values: ["NO_SPEC", "TOOL_SPEC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_agent_registry_service#type GoogleAgentRegistryService#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The content of the MCP Server spec. This payload is validated against the schema for the specified type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_agent_registry_service#content GoogleAgentRegistryService#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
}

