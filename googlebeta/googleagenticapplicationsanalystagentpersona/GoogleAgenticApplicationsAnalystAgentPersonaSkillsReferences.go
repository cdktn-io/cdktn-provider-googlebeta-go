// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaSkillsReferences struct {
	// The content of the reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#content GoogleAgenticApplicationsAnalystAgentPersona#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The identifier of the reference within the skill. Use a descriptive string that reflects the reference's function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#reference_id GoogleAgenticApplicationsAnalystAgentPersona#reference_id}
	ReferenceId *string `field:"required" json:"referenceId" yaml:"referenceId"`
}

