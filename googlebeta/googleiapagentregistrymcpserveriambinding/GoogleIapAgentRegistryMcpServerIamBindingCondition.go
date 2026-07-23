// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiapagentregistrymcpserveriambinding


type GoogleIapAgentRegistryMcpServerIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_iap_agent_registry_mcp_server_iam_binding#expression GoogleIapAgentRegistryMcpServerIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_iap_agent_registry_mcp_server_iam_binding#title GoogleIapAgentRegistryMcpServerIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_iap_agent_registry_mcp_server_iam_binding#description GoogleIapAgentRegistryMcpServerIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

