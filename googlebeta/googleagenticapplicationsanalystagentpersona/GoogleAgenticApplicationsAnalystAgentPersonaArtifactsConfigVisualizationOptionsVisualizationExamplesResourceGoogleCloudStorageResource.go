// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResourceGoogleCloudStorageResource struct {
	// The Google Cloud Storage object or folder.
	//
	// Format: /
	// or: //
	//
	// Note that to refer to a folder, it _must_ end in a slash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona#google_cloud_storage_object GoogleAgenticApplicationsAnalystAgentPersona#google_cloud_storage_object}
	GoogleCloudStorageObject *string `field:"required" json:"googleCloudStorageObject" yaml:"googleCloudStorageObject"`
	// If non-empty, only files with these extensions are included when expanding the resource.  If empty, all files are included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona#file_extension_restrictions GoogleAgenticApplicationsAnalystAgentPersona#file_extension_restrictions}
	FileExtensionRestrictions *[]*string `field:"optional" json:"fileExtensionRestrictions" yaml:"fileExtensionRestrictions"`
}

