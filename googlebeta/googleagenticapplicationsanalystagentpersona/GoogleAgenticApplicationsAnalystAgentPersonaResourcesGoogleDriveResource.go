// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaResourcesGoogleDriveResource struct {
	// If non-empty, only files with these extensions are included when expanding the resource.  If empty, all files are included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#file_extension_restrictions GoogleAgenticApplicationsAnalystAgentPersona#file_extension_restrictions}
	FileExtensionRestrictions *[]*string `field:"optional" json:"fileExtensionRestrictions" yaml:"fileExtensionRestrictions"`
	// Points to a drive file to use.
	//
	// May refer to workspace files or folders
	// as well.  If folder is specifically, all files in the folder
	// (recursively) are used.
	//
	// Expected Format:
	// files/{file_id}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#file_reference GoogleAgenticApplicationsAnalystAgentPersona#file_reference}
	FileReference *string `field:"optional" json:"fileReference" yaml:"fileReference"`
}

