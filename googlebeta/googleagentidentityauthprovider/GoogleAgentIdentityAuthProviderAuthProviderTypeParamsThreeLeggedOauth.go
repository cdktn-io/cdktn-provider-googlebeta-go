// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagentidentityauthprovider


type GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth struct {
	// The authorization endpoint to send users to for consenting to delegate to the agent. eg. "https://auth.atlassian.com/authorize".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agent_identity_auth_provider#authorization_url GoogleAgentIdentityAuthProvider#authorization_url}
	AuthorizationUrl *string `field:"optional" json:"authorizationUrl" yaml:"authorizationUrl"`
	// The client ID of the OAuth client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agent_identity_auth_provider#client_id GoogleAgentIdentityAuthProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Input only. The client secret of the OAuth client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agent_identity_auth_provider#client_secret GoogleAgentIdentityAuthProvider#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Input only. The client secret of the OAuth client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agent_identity_auth_provider#client_secret_wo GoogleAgentIdentityAuthProvider#client_secret_wo}
	ClientSecretWo *string `field:"optional" json:"clientSecretWo" yaml:"clientSecretWo"`
	// Triggers update of 'client_secret_wo' write-only.
	//
	// Increment this value when an update to 'client_secret_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agent_identity_auth_provider#client_secret_wo_version GoogleAgentIdentityAuthProvider#client_secret_wo_version}
	ClientSecretWoVersion *string `field:"optional" json:"clientSecretWoVersion" yaml:"clientSecretWoVersion"`
	// The default continue URI for 3LO flow and it will be used when no continue URI is provided in the RetrieveCredentials request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agent_identity_auth_provider#default_continue_uri GoogleAgentIdentityAuthProvider#default_continue_uri}
	DefaultContinueUri *string `field:"optional" json:"defaultContinueUri" yaml:"defaultContinueUri"`
	// Enables Proof Key for Code Exchange (PKCE) for the OAuth flow to prevent authorization code interception attacks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agent_identity_auth_provider#enable_pkce GoogleAgentIdentityAuthProvider#enable_pkce}
	EnablePkce interface{} `field:"optional" json:"enablePkce" yaml:"enablePkce"`
	// The token endpoint for requesting tokens on behalf of an end user. eg. "https://auth.atlassian.com/oauth/token".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agent_identity_auth_provider#token_url GoogleAgentIdentityAuthProvider#token_url}
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
}

