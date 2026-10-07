// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicesagentconnectivitytemplate


type GoogleNetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfig struct {
	// Defines whether additional roots should be trusted. Possible values: ["NO_ADDITIONAL_ROOTS", "PUBLICLY_TRUSTED_ROOTS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_services_agent_connectivity_template#additional_roots GoogleNetworkServicesAgentConnectivityTemplate#additional_roots}
	AdditionalRoots *string `field:"required" json:"additionalRoots" yaml:"additionalRoots"`
	// The trust config resource name. Format: projects/{project}/locations/{location}/trustConfigs/{trust_config}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_services_agent_connectivity_template#trust_config GoogleNetworkServicesAgentConnectivityTemplate#trust_config}
	TrustConfig *string `field:"optional" json:"trustConfig" yaml:"trustConfig"`
}

