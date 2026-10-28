// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptions struct {
	// analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#analysis GoogleGeminiDataAnalyticsDataAgent#analysis}
	Analysis *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptionsAnalysis `field:"optional" json:"analysis" yaml:"analysis"`
	// datasource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#datasource GoogleGeminiDataAnalyticsDataAgent#datasource}
	Datasource *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptionsDatasource `field:"optional" json:"datasource" yaml:"datasource"`
	// Optional. The model to use for the agent loop.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#model GoogleGeminiDataAnalyticsDataAgent#model}
	Model *string `field:"optional" json:"model" yaml:"model"`
}

