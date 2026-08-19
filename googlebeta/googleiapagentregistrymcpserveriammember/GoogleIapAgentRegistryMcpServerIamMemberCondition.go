// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiapagentregistrymcpserveriammember


type GoogleIapAgentRegistryMcpServerIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_iap_agent_registry_mcp_server_iam_member#expression GoogleIapAgentRegistryMcpServerIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_iap_agent_registry_mcp_server_iam_member#title GoogleIapAgentRegistryMcpServerIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_iap_agent_registry_mcp_server_iam_member#description GoogleIapAgentRegistryMcpServerIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

