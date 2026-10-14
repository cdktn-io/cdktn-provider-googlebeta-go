// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputebackendservice


type GoogleComputeBackendServiceDynamicForwardingForwardProxy struct {
	// A boolean flag enabling dynamic forwarding proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_compute_backend_service#enabled GoogleComputeBackendService#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Determines the dynamic forwarding proxy mode Possible values: ["DIRECT_FORWARDING", "CLOUD_RUN"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_compute_backend_service#proxy_mode GoogleComputeBackendService#proxy_mode}
	ProxyMode *string `field:"required" json:"proxyMode" yaml:"proxyMode"`
}

