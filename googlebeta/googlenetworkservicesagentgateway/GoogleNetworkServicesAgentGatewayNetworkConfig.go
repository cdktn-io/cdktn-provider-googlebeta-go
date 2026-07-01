// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicesagentgateway


type GoogleNetworkServicesAgentGatewayNetworkConfig struct {
	// egress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_network_services_agent_gateway#egress GoogleNetworkServicesAgentGateway#egress}
	Egress *GoogleNetworkServicesAgentGatewayNetworkConfigEgress `field:"required" json:"egress" yaml:"egress"`
	// dns_peering_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_network_services_agent_gateway#dns_peering_config GoogleNetworkServicesAgentGateway#dns_peering_config}
	DnsPeeringConfig *GoogleNetworkServicesAgentGatewayNetworkConfigDnsPeeringConfig `field:"optional" json:"dnsPeeringConfig" yaml:"dnsPeeringConfig"`
}

