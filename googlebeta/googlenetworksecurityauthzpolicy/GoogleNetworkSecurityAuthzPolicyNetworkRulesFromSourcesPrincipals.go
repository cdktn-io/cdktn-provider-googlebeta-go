// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesPrincipals struct {
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_network_security_authz_policy#principal GoogleNetworkSecurityAuthzPolicy#principal}
	Principal *GoogleNetworkSecurityAuthzPolicyNetworkRulesFromSourcesPrincipalsPrincipal `field:"optional" json:"principal" yaml:"principal"`
	// An enum to decide what principal value the principal rule will match against.
	//
	// If not specified, the PrincipalSelector is CLIENT_CERT_URI_SAN. Default value: "CLIENT_CERT_URI_SAN" Possible values: ["PRINCIPAL_SELECTOR_UNSPECIFIED", "CLIENT_CERT_URI_SAN", "CLIENT_CERT_DNS_NAME_SAN", "CLIENT_CERT_COMMON_NAME"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_network_security_authz_policy#principal_selector GoogleNetworkSecurityAuthzPolicy#principal_selector}
	PrincipalSelector *string `field:"optional" json:"principalSelector" yaml:"principalSelector"`
}

