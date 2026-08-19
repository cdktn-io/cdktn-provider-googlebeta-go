// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfig struct {
	// document_generation_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#document_generation_options GoogleAgenticApplicationsAnalystAgentPersona#document_generation_options}
	DocumentGenerationOptions *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigDocumentGenerationOptions `field:"optional" json:"documentGenerationOptions" yaml:"documentGenerationOptions"`
	// slide_generation_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#slide_generation_options GoogleAgenticApplicationsAnalystAgentPersona#slide_generation_options}
	SlideGenerationOptions *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigSlideGenerationOptions `field:"optional" json:"slideGenerationOptions" yaml:"slideGenerationOptions"`
	// visualization_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#visualization_options GoogleAgenticApplicationsAnalystAgentPersona#visualization_options}
	VisualizationOptions *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptions `field:"optional" json:"visualizationOptions" yaml:"visualizationOptions"`
}

