// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgent struct {
	// published_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#published_context GoogleGeminiDataAnalyticsDataAgent#published_context}
	PublishedContext *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContext `field:"optional" json:"publishedContext" yaml:"publishedContext"`
	// staging_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#staging_context GoogleGeminiDataAnalyticsDataAgent#staging_context}
	StagingContext *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContext `field:"optional" json:"stagingContext" yaml:"stagingContext"`
}

