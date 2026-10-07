// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicesagentconnectivitytemplate


type GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfig struct {
	// The URI of the target VPC network for DNS peering. Must be of the form 'projects/{project}/global/networks/{network}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_services_agent_connectivity_template#target_network GoogleNetworkServicesAgentConnectivityTemplate#target_network}
	TargetNetwork *string `field:"required" json:"targetNetwork" yaml:"targetNetwork"`
	// The domain name to peer for DNS resolution.
	//
	// Must be a fully
	// qualified domain name ending with a dot (for example, 'example.com.').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_services_agent_connectivity_template#domain GoogleNetworkServicesAgentConnectivityTemplate#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The list of domain names to peer for DNS resolution.
	//
	// Each entry
	// must be a fully qualified domain name ending with a dot
	// (for example, 'example.com.'). At least one domain must be
	// specified between 'domain' and 'domains'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_services_agent_connectivity_template#domains GoogleNetworkServicesAgentConnectivityTemplate#domains}
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
}

