// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegkehubfleet


type GoogleGkeHubFleetDefaultClusterConfig struct {
	// binary_authorization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_gke_hub_fleet#binary_authorization_config GoogleGkeHubFleet#binary_authorization_config}
	BinaryAuthorizationConfig *GoogleGkeHubFleetDefaultClusterConfigBinaryAuthorizationConfig `field:"optional" json:"binaryAuthorizationConfig" yaml:"binaryAuthorizationConfig"`
	// compliance_posture_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_gke_hub_fleet#compliance_posture_config GoogleGkeHubFleet#compliance_posture_config}
	CompliancePostureConfig *GoogleGkeHubFleetDefaultClusterConfigCompliancePostureConfig `field:"optional" json:"compliancePostureConfig" yaml:"compliancePostureConfig"`
	// security_posture_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_gke_hub_fleet#security_posture_config GoogleGkeHubFleet#security_posture_config}
	SecurityPostureConfig *GoogleGkeHubFleetDefaultClusterConfigSecurityPostureConfig `field:"optional" json:"securityPostureConfig" yaml:"securityPostureConfig"`
}

