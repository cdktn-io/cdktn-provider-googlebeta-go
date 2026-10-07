// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesapp


type GoogleCesAppVpcScSettings struct {
	// The allowed HTTP(s) origins that OpenAPI tools in the App are able to directly call when VPC Service Controls are enabled.
	//
	// These strings
	// must match the origin exactly, including the port if specified. For
	// example, "https://example.com" or "https://example.com:443". This list does
	// not yet apply to Python tools that may make direct HTTP calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_app#allowed_origins GoogleCesApp#allowed_origins}
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
}

