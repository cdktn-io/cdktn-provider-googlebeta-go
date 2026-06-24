// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicesagentgateway


type GoogleNetworkServicesAgentGatewayNetworkConfigDnsPeeringConfig struct {
	// The list of domain names to peer for DNS resolution.
	//
	// Each entry
	// must be a fully qualified domain name ending with a dot
	// (for example, 'example.com.').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_network_services_agent_gateway#domains GoogleNetworkServicesAgentGateway#domains}
	Domains *[]*string `field:"required" json:"domains" yaml:"domains"`
	// The URI of the target VPC network for DNS peering. Must be of the form 'projects/{project}/global/networks/{network}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_network_services_agent_gateway#target_network GoogleNetworkServicesAgentGateway#target_network}
	TargetNetwork *string `field:"required" json:"targetNetwork" yaml:"targetNetwork"`
	// The ID of the project that hosts the target VPC network for DNS peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_network_services_agent_gateway#target_project GoogleNetworkServicesAgentGateway#target_project}
	TargetProject *string `field:"required" json:"targetProject" yaml:"targetProject"`
}

