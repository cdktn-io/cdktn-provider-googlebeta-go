// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecurityfirewallendpoint


type GoogleNetworkSecurityFirewallEndpointEndpointSettings struct {
	// The Palo Alto Networks content cloud region of the firewall endpoint, for example 'US_CENTRAL' or 'CANADA'.
	//
	// See
	// [API docs](https://docs.cloud.google.com/firewall/docs/reference/network-security/rest/v1beta1/organizations.locations.firewallEndpoints#contentcloudregion)
	// for the list of supported regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_network_security_firewall_endpoint#content_cloud_region GoogleNetworkSecurityFirewallEndpoint#content_cloud_region}
	ContentCloudRegion *string `field:"optional" json:"contentCloudRegion" yaml:"contentCloudRegion"`
	// Whether to block HTTP partial responses for the firewall endpoint.
	//
	// When 'true', resumption of blocked malicious HTTP file downloads is
	// blocked by the firewall. 'false' provides maximum availability, 'true'
	// provides maximum security.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_network_security_firewall_endpoint#http_partial_response_blocked GoogleNetworkSecurityFirewallEndpoint#http_partial_response_blocked}
	HttpPartialResponseBlocked interface{} `field:"optional" json:"httpPartialResponseBlocked" yaml:"httpPartialResponseBlocked"`
	// Indicates whether Jumbo Frames are enabled for the firewall endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_network_security_firewall_endpoint#jumbo_frames_enabled GoogleNetworkSecurityFirewallEndpoint#jumbo_frames_enabled}
	JumboFramesEnabled interface{} `field:"optional" json:"jumboFramesEnabled" yaml:"jumboFramesEnabled"`
}

