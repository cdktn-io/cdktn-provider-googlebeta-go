// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueries struct {
	// looker_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#looker_query GoogleGeminiDataAnalyticsDataAgent#looker_query}
	LookerQuery *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesLookerQuery `field:"optional" json:"lookerQuery" yaml:"lookerQuery"`
	// Optional. Natural language questions that a user might ask.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#natural_language_questions GoogleGeminiDataAnalyticsDataAgent#natural_language_questions}
	NaturalLanguageQuestions *[]*string `field:"optional" json:"naturalLanguageQuestions" yaml:"naturalLanguageQuestions"`
}

