// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagentregistrybinding


type GoogleAgentRegistryBindingTarget struct {
	// The identifier of the target Agent, MCP Server, or Endpoint. Format: * 'urn:agent:{publisher}:{namespace}:{name}' * 'urn:mcp:{publisher}:{namespace}:{name}' * 'urn:endpoint:{publisher}:{namespace}:{name}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_agent_registry_binding#identifier GoogleAgentRegistryBinding#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

