// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecuritysecurityprofile


type GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSetting struct {
	// file_exceptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_security_security_profile#file_exceptions GoogleNetworkSecuritySecurityProfile#file_exceptions}
	FileExceptions interface{} `field:"optional" json:"fileExceptions" yaml:"fileExceptions"`
	// inline_ml_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_security_security_profile#inline_ml_configs GoogleNetworkSecuritySecurityProfile#inline_ml_configs}
	InlineMlConfigs interface{} `field:"optional" json:"inlineMlConfigs" yaml:"inlineMlConfigs"`
}

