// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicesagentconnectivitytemplate


type GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfig struct {
	// dns_peering_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_services_agent_connectivity_template#dns_peering_config GoogleNetworkServicesAgentConnectivityTemplate#dns_peering_config}
	DnsPeeringConfig *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfig `field:"optional" json:"dnsPeeringConfig" yaml:"dnsPeeringConfig"`
	// The network attachment resource name. Format: projects/{project}/regions/{region}/networkAttachments/{network_attachment_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_services_agent_connectivity_template#network_attachment GoogleNetworkServicesAgentConnectivityTemplate#network_attachment}
	NetworkAttachment *string `field:"optional" json:"networkAttachment" yaml:"networkAttachment"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_services_agent_connectivity_template#tls_config GoogleNetworkServicesAgentConnectivityTemplate#tls_config}
	TlsConfig *GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
	// The VPC egress setting. Possible values: ["ALL_TRAFFIC", "PRIVATE_RANGES_ONLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_services_agent_connectivity_template#vpc_egress GoogleNetworkServicesAgentConnectivityTemplate#vpc_egress}
	VpcEgress *string `field:"optional" json:"vpcEgress" yaml:"vpcEgress"`
}

