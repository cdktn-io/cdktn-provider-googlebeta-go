// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagentidentityauthprovideriambinding


type GoogleAgentIdentityAuthProviderIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_agent_identity_auth_provider_iam_binding#expression GoogleAgentIdentityAuthProviderIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_agent_identity_auth_provider_iam_binding#title GoogleAgentIdentityAuthProviderIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_agent_identity_auth_provider_iam_binding#description GoogleAgentIdentityAuthProviderIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

