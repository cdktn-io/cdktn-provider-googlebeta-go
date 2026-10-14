// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherDefaultRouteActionDynamicCompressionPolicy struct {
	// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header. Possible values: ["AUTOMATIC", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_compute_url_map#compression_mode GoogleComputeUrlMap#compression_mode}
	CompressionMode *string `field:"optional" json:"compressionMode" yaml:"compressionMode"`
}

