// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegkehubfleet


type GoogleGkeHubFleetDefaultClusterConfigCompliancePostureConfig struct {
	// compliance_standards block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_gke_hub_fleet#compliance_standards GoogleGkeHubFleet#compliance_standards}
	ComplianceStandards interface{} `field:"optional" json:"complianceStandards" yaml:"complianceStandards"`
	// Sets which mode to use for Compliance Posture features. Possible values: ["DISABLED", "ENABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_gke_hub_fleet#mode GoogleGkeHubFleet#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

