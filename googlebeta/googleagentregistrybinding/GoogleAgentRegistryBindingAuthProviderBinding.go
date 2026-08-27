// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagentregistrybinding


type GoogleAgentRegistryBindingAuthProviderBinding struct {
	// The resource name of the target auth provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agent_registry_binding#auth_provider GoogleAgentRegistryBinding#auth_provider}
	AuthProvider *string `field:"required" json:"authProvider" yaml:"authProvider"`
	// The continue URI of the auth provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agent_registry_binding#continue_uri GoogleAgentRegistryBinding#continue_uri}
	ContinueUri *string `field:"optional" json:"continueUri" yaml:"continueUri"`
	// The list of OAuth2 scopes of the auth provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agent_registry_binding#scopes GoogleAgentRegistryBinding#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

