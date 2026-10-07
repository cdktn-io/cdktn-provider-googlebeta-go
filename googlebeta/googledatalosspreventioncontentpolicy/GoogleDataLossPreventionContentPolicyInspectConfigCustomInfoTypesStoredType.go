// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventioncontentpolicy


type GoogleDataLossPreventionContentPolicyInspectConfigCustomInfoTypesStoredType struct {
	// Resource name of the requested StoredInfoType, for example 'organizations/433245324/storedInfoTypes/432452342' or 'projects/project-id/storedInfoTypes/432452342'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_data_loss_prevention_content_policy#name GoogleDataLossPreventionContentPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

