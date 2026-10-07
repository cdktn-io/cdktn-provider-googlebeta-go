// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy


type GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesFileLabelInfoTypeGoogleDriveLabel struct {
	// The label ID of the Google Drive label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#label_id GoogleDataLossPreventionContentPolicy#label_id}
	LabelId *string `field:"required" json:"labelId" yaml:"labelId"`
	// label_fields_to_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#label_fields_to_match GoogleDataLossPreventionContentPolicy#label_fields_to_match}
	LabelFieldsToMatch interface{} `field:"optional" json:"labelFieldsToMatch" yaml:"labelFieldsToMatch"`
}

