// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecurityfirewallendpoint


type GoogleNetworkSecurityFirewallEndpointWildfireSettingsWildfireInlineCloudAnalysisSettings struct {
	// The maximum duration for which the firewall endpoint holds a file while waiting for an inline cloud analysis verdict before taking the configured timeout action (1ms to 240000ms, default 30000ms).
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'.
	// Example: "30s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_security_firewall_endpoint#max_analysis_duration GoogleNetworkSecurityFirewallEndpoint#max_analysis_duration}
	MaxAnalysisDuration *string `field:"optional" json:"maxAnalysisDuration" yaml:"maxAnalysisDuration"`
	// Indicates whether logging is disabled when a WildFire inline cloud analysis submission times out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_security_firewall_endpoint#submission_timeout_logging_disabled GoogleNetworkSecurityFirewallEndpoint#submission_timeout_logging_disabled}
	SubmissionTimeoutLoggingDisabled interface{} `field:"optional" json:"submissionTimeoutLoggingDisabled" yaml:"submissionTimeoutLoggingDisabled"`
	// The action to take when WildFire inline cloud analysis times out. Possible values: ["ALLOW", "DENY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_network_security_firewall_endpoint#timeout_action GoogleNetworkSecurityFirewallEndpoint#timeout_action}
	TimeoutAction *string `field:"optional" json:"timeoutAction" yaml:"timeoutAction"`
}

