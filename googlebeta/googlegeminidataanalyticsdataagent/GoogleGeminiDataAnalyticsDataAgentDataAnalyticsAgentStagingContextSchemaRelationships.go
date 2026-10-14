// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextSchemaRelationships struct {
	// left_schema_paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#left_schema_paths GoogleGeminiDataAnalyticsDataAgent#left_schema_paths}
	LeftSchemaPaths *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextSchemaRelationshipsLeftSchemaPaths `field:"required" json:"leftSchemaPaths" yaml:"leftSchemaPaths"`
	// right_schema_paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#right_schema_paths GoogleGeminiDataAnalyticsDataAgent#right_schema_paths}
	RightSchemaPaths *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextSchemaRelationshipsRightSchemaPaths `field:"required" json:"rightSchemaPaths" yaml:"rightSchemaPaths"`
	// Optional. A confidence score for the suggested relationship.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#confidence_score GoogleGeminiDataAnalyticsDataAgent#confidence_score}
	ConfidenceScore *float64 `field:"optional" json:"confidenceScore" yaml:"confidenceScore"`
	// Optional. Sources which generated the schema relation edge. Possible values: BIGQUERY_JOB_HISTORY, LLM_SUGGESTED, BIGQUERY_TABLE_CONSTRAINTS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#sources GoogleGeminiDataAnalyticsDataAgent#sources}
	Sources *[]*string `field:"optional" json:"sources" yaml:"sources"`
}

