// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesPrincipalsPrincipal struct {
	// The input string must match exactly the string specified here. Examples: * abc only matches the value abc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_network_security_authz_policy#exact GoogleNetworkSecurityAuthzPolicy#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
}

