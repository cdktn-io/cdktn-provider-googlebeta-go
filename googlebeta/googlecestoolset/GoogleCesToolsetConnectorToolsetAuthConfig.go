// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestoolset


type GoogleCesToolsetConnectorToolsetAuthConfig struct {
	// oauth2_auth_code_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_ces_toolset#oauth2_auth_code_config GoogleCesToolset#oauth2_auth_code_config}
	Oauth2AuthCodeConfig *GoogleCesToolsetConnectorToolsetAuthConfigOauth2AuthCodeConfig `field:"optional" json:"oauth2AuthCodeConfig" yaml:"oauth2AuthCodeConfig"`
	// oauth2_jwt_bearer_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_ces_toolset#oauth2_jwt_bearer_config GoogleCesToolset#oauth2_jwt_bearer_config}
	Oauth2JwtBearerConfig *GoogleCesToolsetConnectorToolsetAuthConfigOauth2JwtBearerConfig `field:"optional" json:"oauth2JwtBearerConfig" yaml:"oauth2JwtBearerConfig"`
}

