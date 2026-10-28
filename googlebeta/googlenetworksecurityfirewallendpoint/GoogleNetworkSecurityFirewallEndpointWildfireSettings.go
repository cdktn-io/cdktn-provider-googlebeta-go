// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecurityfirewallendpoint


type GoogleNetworkSecurityFirewallEndpointWildfireSettings struct {
	// Indicates whether WildFire is enabled for the firewall endpoint.
	//
	// Enabling WildFire requires accepting the WildFire end-user license
	// agreement (EULA) for the endpoint's billing project; see
	// [Configure WildFire](https://docs.cloud.google.com/firewall/docs/configure-wildfire).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_security_firewall_endpoint#enabled GoogleNetworkSecurityFirewallEndpoint#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// wildfire_inline_cloud_analysis_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_security_firewall_endpoint#wildfire_inline_cloud_analysis_settings GoogleNetworkSecurityFirewallEndpoint#wildfire_inline_cloud_analysis_settings}
	WildfireInlineCloudAnalysisSettings *GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings `field:"optional" json:"wildfireInlineCloudAnalysisSettings" yaml:"wildfireInlineCloudAnalysisSettings"`
	// The duration for which the firewall endpoint waits for a real-time signature lookup response from the WildFire cloud before taking the configured timeout action (1ms to 5000ms, default 1000ms).
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'.
	// Example: "1s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_security_firewall_endpoint#wildfire_realtime_lookup_duration GoogleNetworkSecurityFirewallEndpoint#wildfire_realtime_lookup_duration}
	WildfireRealtimeLookupDuration *string `field:"optional" json:"wildfireRealtimeLookupDuration" yaml:"wildfireRealtimeLookupDuration"`
	// The action to take when a WildFire real-time signature lookup times out. Possible values: ["ALLOW", "DENY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_security_firewall_endpoint#wildfire_realtime_lookup_timeout_action GoogleNetworkSecurityFirewallEndpoint#wildfire_realtime_lookup_timeout_action}
	WildfireRealtimeLookupTimeoutAction *string `field:"optional" json:"wildfireRealtimeLookupTimeoutAction" yaml:"wildfireRealtimeLookupTimeoutAction"`
	// The WildFire region for the firewall endpoint, for example 'UNITED_STATES' or 'CANADA'. See [API docs](https://docs.cloud.google.com/firewall/docs/reference/network-security/rest/v1beta1/organizations.locations.firewallEndpoints#wildfireregion) for the list of supported regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_security_firewall_endpoint#wildfire_region GoogleNetworkSecurityFirewallEndpoint#wildfire_region}
	WildfireRegion *string `field:"optional" json:"wildfireRegion" yaml:"wildfireRegion"`
}

