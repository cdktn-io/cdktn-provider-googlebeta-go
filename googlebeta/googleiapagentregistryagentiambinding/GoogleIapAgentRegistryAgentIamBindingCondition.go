// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiapagentregistryagentiambinding


type GoogleIapAgentRegistryAgentIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_iap_agent_registry_agent_iam_binding#expression GoogleIapAgentRegistryAgentIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_iap_agent_registry_agent_iam_binding#title GoogleIapAgentRegistryAgentIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_iap_agent_registry_agent_iam_binding#description GoogleIapAgentRegistryAgentIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

