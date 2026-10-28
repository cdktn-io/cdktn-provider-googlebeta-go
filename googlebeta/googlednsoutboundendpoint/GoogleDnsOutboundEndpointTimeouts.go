// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlednsoutboundendpoint


type GoogleDnsOutboundEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_dns_outbound_endpoint#create GoogleDnsOutboundEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_dns_outbound_endpoint#delete GoogleDnsOutboundEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_dns_outbound_endpoint#update GoogleDnsOutboundEndpoint#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

