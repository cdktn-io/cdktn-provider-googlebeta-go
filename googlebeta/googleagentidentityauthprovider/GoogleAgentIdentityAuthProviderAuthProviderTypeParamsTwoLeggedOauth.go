// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagentidentityauthprovider


type GoogleAgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth struct {
	// The client ID of the OAuth client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_agent_identity_auth_provider#client_id GoogleAgentIdentityAuthProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Input only. The client secret of the OAuth client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_agent_identity_auth_provider#client_secret GoogleAgentIdentityAuthProvider#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Input only. The client secret of the OAuth client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_agent_identity_auth_provider#client_secret_wo GoogleAgentIdentityAuthProvider#client_secret_wo}
	ClientSecretWo *string `field:"optional" json:"clientSecretWo" yaml:"clientSecretWo"`
	// Triggers update of 'client_secret_wo' write-only.
	//
	// Increment this value when an update to 'client_secret_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_agent_identity_auth_provider#client_secret_wo_version GoogleAgentIdentityAuthProvider#client_secret_wo_version}
	ClientSecretWoVersion *string `field:"optional" json:"clientSecretWoVersion" yaml:"clientSecretWoVersion"`
	// The token endpoint of the OAuth client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_agent_identity_auth_provider#token_url GoogleAgentIdentityAuthProvider#token_url}
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
}

