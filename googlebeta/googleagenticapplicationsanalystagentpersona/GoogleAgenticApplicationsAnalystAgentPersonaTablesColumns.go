// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaTablesColumns struct {
	// The data type of the column.
	//
	// This should be a GoogleSQL data type.
	// Parameterized types such as PROTO, ENUM, ARRAY, STRUCT<...>, and
	// RANGE are not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#data_type GoogleAgenticApplicationsAnalystAgentPersona#data_type}
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// The name of the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#name GoogleAgenticApplicationsAnalystAgentPersona#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#description GoogleAgenticApplicationsAnalystAgentPersona#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

