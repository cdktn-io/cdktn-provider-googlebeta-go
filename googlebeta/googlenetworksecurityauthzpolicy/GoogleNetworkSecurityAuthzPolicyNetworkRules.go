// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyNetworkRules struct {
	// from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_network_security_authz_policy#from GoogleNetworkSecurityAuthzPolicy#from}
	From *GoogleNetworkSecurityAuthzPolicyNetworkRulesFrom `field:"optional" json:"from" yaml:"from"`
	// to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.38.0/docs/resources/google_network_security_authz_policy#to GoogleNetworkSecurityAuthzPolicy#to}
	To *GoogleNetworkSecurityAuthzPolicyNetworkRulesTo `field:"optional" json:"to" yaml:"to"`
}

