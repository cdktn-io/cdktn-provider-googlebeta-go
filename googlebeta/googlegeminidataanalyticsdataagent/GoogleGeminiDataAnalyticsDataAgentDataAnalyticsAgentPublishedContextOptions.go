// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptions struct {
	// analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#analysis GoogleGeminiDataAnalyticsDataAgent#analysis}
	Analysis *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsAnalysis `field:"optional" json:"analysis" yaml:"analysis"`
	// datasource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#datasource GoogleGeminiDataAnalyticsDataAgent#datasource}
	Datasource *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextOptionsDatasource `field:"optional" json:"datasource" yaml:"datasource"`
}

