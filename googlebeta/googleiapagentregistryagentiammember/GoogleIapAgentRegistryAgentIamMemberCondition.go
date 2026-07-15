// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiapagentregistryagentiammember


type GoogleIapAgentRegistryAgentIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_iap_agent_registry_agent_iam_member#expression GoogleIapAgentRegistryAgentIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_iap_agent_registry_agent_iam_member#title GoogleIapAgentRegistryAgentIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_iap_agent_registry_agent_iam_member#description GoogleIapAgentRegistryAgentIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

