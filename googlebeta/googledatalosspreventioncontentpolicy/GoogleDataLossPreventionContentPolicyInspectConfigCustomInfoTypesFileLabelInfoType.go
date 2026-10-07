// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy


type GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoType struct {
	// google_drive_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#google_drive_label GoogleDataLossPreventionContentPolicy#google_drive_label}
	GoogleDriveLabel *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabel `field:"optional" json:"googleDriveLabel" yaml:"googleDriveLabel"`
	// sensitivity_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#sensitivity_label GoogleDataLossPreventionContentPolicy#sensitivity_label}
	SensitivityLabel *GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeSensitivityLabel `field:"optional" json:"sensitivityLabel" yaml:"sensitivityLabel"`
}

