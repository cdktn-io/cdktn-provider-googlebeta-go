// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleiapagentregistryendpointiammember


type GoogleIapAgentRegistryEndpointIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_iap_agent_registry_endpoint_iam_member#expression GoogleIapAgentRegistryEndpointIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_iap_agent_registry_endpoint_iam_member#title GoogleIapAgentRegistryEndpointIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.41.0/docs/resources/google_iap_agent_registry_endpoint_iam_member#description GoogleIapAgentRegistryEndpointIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

