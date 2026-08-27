// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaSkills struct {
	// The markdown text content of the skill.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#content GoogleAgenticApplicationsAnalystAgentPersona#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The identifier of the skill. Use a descriptive string that reflects the skill's function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#skill_id GoogleAgenticApplicationsAnalystAgentPersona#skill_id}
	SkillId *string `field:"required" json:"skillId" yaml:"skillId"`
	// The description of the skill.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#description GoogleAgenticApplicationsAnalystAgentPersona#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// references block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.0/docs/resources/google_agentic_applications_analyst_agent_persona#references GoogleAgenticApplicationsAnalystAgentPersona#references}
	References interface{} `field:"optional" json:"references" yaml:"references"`
}

