// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicesedgecacheorigin


type GoogleNetworkServicesEdgeCacheOriginFlexShielding struct {
	// Whenever possible, content will be fetched from origin and cached in or near the specified origin. Best effort.
	//
	// You must specify exactly one FlexShieldingRegion. Possible values include
	// 'AFRICA_SOUTH1', 'ME_CENTRAL1', 'EUROPE_WEST3', and 'US_EAST5'. See
	// [FlexShieldingRegion](https://cloud.google.com/media-cdn/docs/reference/rest/v1/projects.locations.edgeCacheOrigins#FlexShieldingRegion)
	// for the full list of supported regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_network_services_edge_cache_origin#flex_shielding_regions GoogleNetworkServicesEdgeCacheOrigin#flex_shielding_regions}
	FlexShieldingRegions *[]*string `field:"optional" json:"flexShieldingRegions" yaml:"flexShieldingRegions"`
}

