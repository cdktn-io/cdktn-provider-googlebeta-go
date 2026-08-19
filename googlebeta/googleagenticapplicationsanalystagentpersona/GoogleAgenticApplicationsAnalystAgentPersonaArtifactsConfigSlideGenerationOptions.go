// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptions struct {
	// Format for slide export. Possible values: PDF PNG PPTX GOOGLE_SLIDES.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#export_format GoogleAgenticApplicationsAnalystAgentPersona#export_format}
	ExportFormat *string `field:"optional" json:"exportFormat" yaml:"exportFormat"`
	// slide_examples block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#slide_examples GoogleAgenticApplicationsAnalystAgentPersona#slide_examples}
	SlideExamples interface{} `field:"optional" json:"slideExamples" yaml:"slideExamples"`
}

