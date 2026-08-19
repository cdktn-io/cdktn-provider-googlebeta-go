// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googleagenticapplicationsanalystagentpersona


type GoogleAgenticApplicationsAnalystAgentPersonaResourcesBigqueryResource struct {
	// Points to a bigquery dataset to use.
	//
	// Expected Format:
	// projects/{project_id_or_number}/datasets/{dataset_id}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#bigquery_dataset GoogleAgenticApplicationsAnalystAgentPersona#bigquery_dataset}
	BigqueryDataset *string `field:"optional" json:"bigqueryDataset" yaml:"bigqueryDataset"`
	// Points to a bigquery table to use.
	//
	// Expected Format:
	// projects/{project_id_or_number}/datasets/{dataset_id}/tables/{table_id}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#bigquery_table GoogleAgenticApplicationsAnalystAgentPersona#bigquery_table}
	BigqueryTable *string `field:"optional" json:"bigqueryTable" yaml:"bigqueryTable"`
	// A map of column names to column descriptions for the bigquery_table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_agentic_applications_analyst_agent_persona#column_descriptions GoogleAgenticApplicationsAnalystAgentPersona#column_descriptions}
	ColumnDescriptions *map[string]*string `field:"optional" json:"columnDescriptions" yaml:"columnDescriptions"`
}

