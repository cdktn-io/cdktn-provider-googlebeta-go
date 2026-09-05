// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagentidentityauthprovider


type GoogleAgentIdentityAuthProviderAuthProviderTypeParams struct {
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_identity_auth_provider#api_key GoogleAgentIdentityAuthProvider#api_key}
	ApiKey *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// three_legged_oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_identity_auth_provider#three_legged_oauth GoogleAgentIdentityAuthProvider#three_legged_oauth}
	ThreeLeggedOauth *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth `field:"optional" json:"threeLeggedOauth" yaml:"threeLeggedOauth"`
	// two_legged_oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agent_identity_auth_provider#two_legged_oauth GoogleAgentIdentityAuthProvider#two_legged_oauth}
	TwoLeggedOauth *GoogleAgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth `field:"optional" json:"twoLeggedOauth" yaml:"twoLeggedOauth"`
}

