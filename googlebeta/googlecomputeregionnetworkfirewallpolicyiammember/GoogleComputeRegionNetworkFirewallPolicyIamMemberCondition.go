// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeregionnetworkfirewallpolicyiammember


type GoogleComputeRegionNetworkFirewallPolicyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_compute_region_network_firewall_policy_iam_member#expression GoogleComputeRegionNetworkFirewallPolicyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_compute_region_network_firewall_policy_iam_member#title GoogleComputeRegionNetworkFirewallPolicyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_compute_region_network_firewall_policy_iam_member#description GoogleComputeRegionNetworkFirewallPolicyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

