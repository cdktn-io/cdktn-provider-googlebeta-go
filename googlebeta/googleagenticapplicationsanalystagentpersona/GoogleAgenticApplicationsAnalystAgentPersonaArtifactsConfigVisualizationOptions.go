// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptions struct {
	// visualization_examples block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_agentic_applications_analyst_agent_persona#visualization_examples GoogleAgenticApplicationsAnalystAgentPersona#visualization_examples}
	VisualizationExamples interface{} `field:"optional" json:"visualizationExamples" yaml:"visualizationExamples"`
	// Mode for generating visualizations. Possible values: VISUALIZATION_MODE_EXPLICIT_ONLY VISUALIZATION_MODE_WHEN_NECESSARY VISUALIZATION_MODE_WHEN_HELPFUL VISUALIZATION_MODE_ALWAYS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_agentic_applications_analyst_agent_persona#visualization_mode GoogleAgenticApplicationsAnalystAgentPersona#visualization_mode}
	VisualizationMode *string `field:"optional" json:"visualizationMode" yaml:"visualizationMode"`
}

