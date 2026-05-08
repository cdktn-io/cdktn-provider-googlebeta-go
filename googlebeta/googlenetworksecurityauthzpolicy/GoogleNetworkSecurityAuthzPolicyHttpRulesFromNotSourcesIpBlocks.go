// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyHttpRulesFromNotSourcesIpBlocks struct {
	// The length of the address range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_authz_policy#length GoogleNetworkSecurityAuthzPolicy#length}
	Length *float64 `field:"required" json:"length" yaml:"length"`
	// The address prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_authz_policy#prefix GoogleNetworkSecurityAuthzPolicy#prefix}
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
}

