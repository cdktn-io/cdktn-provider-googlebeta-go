// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesagent


type GoogleCesAgentRemoteA2AAgentA2AConfigAgentCardSupportedInterfaces struct {
	// The protocol binding supported at this URL. The core ones officially supported are JSONRPC, GRPC and HTTP+JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#protocol_binding GoogleCesAgent#protocol_binding}
	ProtocolBinding *string `field:"required" json:"protocolBinding" yaml:"protocolBinding"`
	// The version of the A2A protocol this interface exposes. Examples: "0.3", "1.0".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#protocol_version GoogleCesAgent#protocol_version}
	ProtocolVersion *string `field:"required" json:"protocolVersion" yaml:"protocolVersion"`
	// The URL where this interface is available. Must be a valid absolute HTTPS URL in production.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#url GoogleCesAgent#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Tenant ID to be used in the request when calling the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#tenant GoogleCesAgent#tenant}
	Tenant *string `field:"optional" json:"tenant" yaml:"tenant"`
}

