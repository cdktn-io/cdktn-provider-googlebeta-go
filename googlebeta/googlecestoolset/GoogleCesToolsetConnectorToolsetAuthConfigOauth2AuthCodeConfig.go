// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestoolset


type GoogleCesToolsetConnectorToolsetAuthConfigOauth2AuthCodeConfig struct {
	// Oauth token parameter name to pass through. Must be in the format '$context.variables.<name_of_variable>'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_ces_toolset#oauth_token GoogleCesToolset#oauth_token}
	OauthToken *string `field:"required" json:"oauthToken" yaml:"oauthToken"`
}

