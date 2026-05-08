// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechroniclefeed


type GoogleChronicleFeedDetailsSalesforceSettings struct {
	// API hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#hostname GoogleChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// oauth_jwt_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#oauth_jwt_credentials GoogleChronicleFeed#oauth_jwt_credentials}
	OauthJwtCredentials *GoogleChronicleFeedDetailsSalesforceSettingsOauthJwtCredentials `field:"optional" json:"oauthJwtCredentials" yaml:"oauthJwtCredentials"`
	// oauth_password_grant_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#oauth_password_grant_auth GoogleChronicleFeed#oauth_password_grant_auth}
	OauthPasswordGrantAuth *GoogleChronicleFeedDetailsSalesforceSettingsOauthPasswordGrantAuth `field:"optional" json:"oauthPasswordGrantAuth" yaml:"oauthPasswordGrantAuth"`
}

