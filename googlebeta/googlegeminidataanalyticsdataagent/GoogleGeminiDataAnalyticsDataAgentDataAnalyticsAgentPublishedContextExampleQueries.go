// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextExampleQueries struct {
	// Optional. A natural language question that a user might ask.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#natural_language_question GoogleGeminiDataAnalyticsDataAgent#natural_language_question}
	NaturalLanguageQuestion *string `field:"optional" json:"naturalLanguageQuestion" yaml:"naturalLanguageQuestion"`
	// Optional. The SQL query that should be generated to answer the natural language question.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#sql_query GoogleGeminiDataAnalyticsDataAgent#sql_query}
	SqlQuery *string `field:"optional" json:"sqlQuery" yaml:"sqlQuery"`
}

