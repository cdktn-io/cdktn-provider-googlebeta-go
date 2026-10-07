// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecuritysecurityprofile


type GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfile struct {
	// wildfire_inline_cloud_analysis_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_security_security_profile#wildfire_inline_cloud_analysis_rules GoogleNetworkSecuritySecurityProfile#wildfire_inline_cloud_analysis_rules}
	WildfireInlineCloudAnalysisRules interface{} `field:"optional" json:"wildfireInlineCloudAnalysisRules" yaml:"wildfireInlineCloudAnalysisRules"`
	// wildfire_inline_ml_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_security_security_profile#wildfire_inline_ml_overrides GoogleNetworkSecuritySecurityProfile#wildfire_inline_ml_overrides}
	WildfireInlineMlOverrides interface{} `field:"optional" json:"wildfireInlineMlOverrides" yaml:"wildfireInlineMlOverrides"`
	// wildfire_inline_ml_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_security_security_profile#wildfire_inline_ml_setting GoogleNetworkSecuritySecurityProfile#wildfire_inline_ml_setting}
	WildfireInlineMlSetting *GoogleNetworkSecuritySecurityProfileWildfireAnalysisProfileWildfireInlineMlSetting `field:"optional" json:"wildfireInlineMlSetting" yaml:"wildfireInlineMlSetting"`
	// wildfire_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_security_security_profile#wildfire_overrides GoogleNetworkSecuritySecurityProfile#wildfire_overrides}
	WildfireOverrides interface{} `field:"optional" json:"wildfireOverrides" yaml:"wildfireOverrides"`
	// Whether to hold the transfer of a file while the WildFire real-time signature cloud performs a signature lookup.
	//
	// Default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_security_security_profile#wildfire_realtime_lookup GoogleNetworkSecuritySecurityProfile#wildfire_realtime_lookup}
	WildfireRealtimeLookup interface{} `field:"optional" json:"wildfireRealtimeLookup" yaml:"wildfireRealtimeLookup"`
	// wildfire_submission_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_security_security_profile#wildfire_submission_rules GoogleNetworkSecuritySecurityProfile#wildfire_submission_rules}
	WildfireSubmissionRules interface{} `field:"optional" json:"wildfireSubmissionRules" yaml:"wildfireSubmissionRules"`
	// wildfire_threat_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_network_security_security_profile#wildfire_threat_overrides GoogleNetworkSecuritySecurityProfile#wildfire_threat_overrides}
	WildfireThreatOverrides interface{} `field:"optional" json:"wildfireThreatOverrides" yaml:"wildfireThreatOverrides"`
}

