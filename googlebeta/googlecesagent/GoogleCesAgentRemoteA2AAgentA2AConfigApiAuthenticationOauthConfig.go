// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesagent


type GoogleCesAgentRemoteA2AAgentA2AConfigApiAuthenticationOauthConfig struct {
	// The client ID from the OAuth provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#client_id GoogleCesAgent#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The name of the SecretManager secret version resource storing the client secret. Format: 'projects/{project}/secrets/{secret}/versions/{version}'.
	//
	// Note: You should grant 'roles/secretmanager.secretAccessor' role to the CES
	// service agent
	// 'service-@gcp-sa-ces.iam.gserviceaccount.com'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#client_secret_version GoogleCesAgent#client_secret_version}
	ClientSecretVersion *string `field:"required" json:"clientSecretVersion" yaml:"clientSecretVersion"`
	// OAuth grant types. Possible values: CLIENT_CREDENTIAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#oauth_grant_type GoogleCesAgent#oauth_grant_type}
	OauthGrantType *string `field:"required" json:"oauthGrantType" yaml:"oauthGrantType"`
	// The token endpoint in the OAuth provider to exchange for an access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#token_endpoint GoogleCesAgent#token_endpoint}
	TokenEndpoint *string `field:"required" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The OAuth scopes to grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_agent#scopes GoogleCesAgent#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

