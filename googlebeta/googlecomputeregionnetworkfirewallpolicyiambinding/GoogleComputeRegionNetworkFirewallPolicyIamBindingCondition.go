// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputeregionnetworkfirewallpolicyiambinding


type GoogleComputeRegionNetworkFirewallPolicyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_compute_region_network_firewall_policy_iam_binding#expression GoogleComputeRegionNetworkFirewallPolicyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_compute_region_network_firewall_policy_iam_binding#title GoogleComputeRegionNetworkFirewallPolicyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_compute_region_network_firewall_policy_iam_binding#description GoogleComputeRegionNetworkFirewallPolicyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

