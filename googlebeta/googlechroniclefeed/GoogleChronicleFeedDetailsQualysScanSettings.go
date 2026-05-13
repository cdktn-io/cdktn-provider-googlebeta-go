// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlechroniclefeed


type GoogleChronicleFeedDetailsQualysScanSettings struct {
	// Supported Qualys Scan api type. Possible values: SCAN_SUMMARY_OUTPUT SCAN_COMPLIANCE_OUTPUT SCAN_COMPLIANCE_CONTROL_OUTPUT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_chronicle_feed#api_type GoogleChronicleFeed#api_type}
	ApiType *string `field:"optional" json:"apiType" yaml:"apiType"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsQualysScanSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_chronicle_feed#hostname GoogleChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
}

