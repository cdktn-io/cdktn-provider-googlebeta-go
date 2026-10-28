// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputenetworkedgesecurityservice


type GoogleComputeNetworkEdgeSecurityServiceParams struct {
	// Resource manager tags to be bound to the networkEdgeSecurityService.
	//
	// Tag keys and values have the
	// same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id},
	// and values are in the format tagValues/456.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_compute_network_edge_security_service#resource_manager_tags GoogleComputeNetworkEdgeSecurityService#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
}

