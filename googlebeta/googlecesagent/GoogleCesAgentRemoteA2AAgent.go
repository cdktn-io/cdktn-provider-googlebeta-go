// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesagent


type GoogleCesAgentRemoteA2AAgent struct {
	// a2a_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#a2a_config GoogleCesAgent#a2a_config}
	A2AConfig *GoogleCesAgentRemoteA2AAgentA2AConfig `field:"required" json:"a2AConfig" yaml:"a2AConfig"`
}

