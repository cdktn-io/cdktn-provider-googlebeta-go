// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledialogflowcxtool


type GoogleDialogflowCxToolConnectorSpecEndUserAuthConfigOauth2AuthCodeConfig struct {
	// Oauth token value or parameter name to pass it through.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_dialogflow_cx_tool#oauth_token GoogleDialogflowCxTool#oauth_token}
	OauthToken *string `field:"required" json:"oauthToken" yaml:"oauthToken"`
}

