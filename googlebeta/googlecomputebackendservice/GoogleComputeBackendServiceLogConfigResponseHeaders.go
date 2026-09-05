// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputebackendservice


type GoogleComputeBackendServiceLogConfigResponseHeaders struct {
	// The header name to match on for logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_compute_backend_service#header_name GoogleComputeBackendService#header_name}
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
}

