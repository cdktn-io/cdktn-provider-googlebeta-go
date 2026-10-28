// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlednsmanagedzone


type GoogleDnsManagedZoneForwardingConfigOutboundEndpoints struct {
	// Specified as a resource path including project and region. This should be formatted like 'projects/{project}/locations/{location}/outboundEndpoints/{endpoint}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_dns_managed_zone#name GoogleDnsManagedZone#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

