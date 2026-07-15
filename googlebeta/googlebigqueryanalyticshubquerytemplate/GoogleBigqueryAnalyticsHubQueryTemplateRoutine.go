// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlebigqueryanalyticshubquerytemplate


type GoogleBigqueryAnalyticsHubQueryTemplateRoutine struct {
	// SQL query logic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_bigquery_analytics_hub_query_template#definition_body GoogleBigqueryAnalyticsHubQueryTemplate#definition_body}
	DefinitionBody *string `field:"optional" json:"definitionBody" yaml:"definitionBody"`
	// Type of routine (e.g., TABLE_VALUED_FUNCTION). Possible values: ["TABLE_VALUED_FUNCTION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.40.0/docs/resources/google_bigquery_analytics_hub_query_template#routine_type GoogleBigqueryAnalyticsHubQueryTemplate#routine_type}
	RoutineType *string `field:"optional" json:"routineType" yaml:"routineType"`
}

