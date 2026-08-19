// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamples struct {
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#resource GoogleAgenticApplicationsAnalystAgentPersona#resource}
	Resource *GoogleAgenticApplicationsAnalystAgentPersonaArtifactsConfigVisualizationOptionsVisualizationExamplesResource `field:"required" json:"resource" yaml:"resource"`
	// The type of the visualization (e.g. "Bar Chart", "Line Chart").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#visualization_type GoogleAgenticApplicationsAnalystAgentPersona#visualization_type}
	VisualizationType *string `field:"required" json:"visualizationType" yaml:"visualizationType"`
}

