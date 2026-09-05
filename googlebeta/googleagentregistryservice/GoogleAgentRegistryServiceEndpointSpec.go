// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagentregistryservice


type GoogleAgentRegistryServiceEndpointSpec struct {
	// The type of the Endpoint spec content. Possible values: ["NO_SPEC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_registry_service#type GoogleAgentRegistryService#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

