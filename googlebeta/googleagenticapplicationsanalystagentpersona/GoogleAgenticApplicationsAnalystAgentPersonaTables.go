// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaTables struct {
	// The name of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona#name GoogleAgenticApplicationsAnalystAgentPersona#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona#columns GoogleAgenticApplicationsAnalystAgentPersona#columns}
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// The description of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_agentic_applications_analyst_agent_persona#description GoogleAgenticApplicationsAnalystAgentPersona#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

