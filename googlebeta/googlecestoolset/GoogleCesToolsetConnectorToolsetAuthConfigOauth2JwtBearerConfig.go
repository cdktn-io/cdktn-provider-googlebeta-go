// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestoolset


type GoogleCesToolsetConnectorToolsetAuthConfigOauth2JwtBearerConfig struct {
	// Client parameter name to pass through. Must be in the format '$context.variables.<name_of_variable>'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_ces_toolset#client_key GoogleCesToolset#client_key}
	ClientKey *string `field:"required" json:"clientKey" yaml:"clientKey"`
	// Issuer parameter name to pass through. Must be in the format '$context.variables.<name_of_variable>'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_ces_toolset#issuer GoogleCesToolset#issuer}
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Subject parameter name to pass through. Must be in the format '$context.variables.<name_of_variable>'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_ces_toolset#subject GoogleCesToolset#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

