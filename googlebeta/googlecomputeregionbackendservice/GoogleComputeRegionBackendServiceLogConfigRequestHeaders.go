// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeregionbackendservice


type GoogleComputeRegionBackendServiceLogConfigRequestHeaders struct {
	// The header name to match on for logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_compute_region_backend_service#header_name GoogleComputeRegionBackendService#header_name}
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
}

