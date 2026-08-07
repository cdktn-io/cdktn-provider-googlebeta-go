// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecomputefirewallpolicyiambinding


type GoogleComputeFirewallPolicyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_compute_firewall_policy_iam_binding#expression GoogleComputeFirewallPolicyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_compute_firewall_policy_iam_binding#title GoogleComputeFirewallPolicyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.43.0/docs/resources/google_compute_firewall_policy_iam_binding#description GoogleComputeFirewallPolicyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

