// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicescdnedgeextension


type GoogleNetworkServicesCdnEdgeExtensionTarget struct {
	// A list of references to the resources that are targeted by CdnEdgeExtension.
	//
	// Types of resources supported: EdgeCacheService.
	// The format can be 'projects/{project}/locations/{location}/edgeCacheServices/{edgeCacheService}' or
	// '//networkservices.googleapis.com/projects/{project}/locations/{location}/edgeCacheServices/{edgeCacheService}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_services_cdn_edge_extension#resources GoogleNetworkServicesCdnEdgeExtension#resources}
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}

