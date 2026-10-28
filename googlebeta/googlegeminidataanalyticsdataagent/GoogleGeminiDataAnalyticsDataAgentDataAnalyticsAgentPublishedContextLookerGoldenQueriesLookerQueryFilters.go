// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesLookerQueryFilters struct {
	// Required. The field to filter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#field GoogleGeminiDataAnalyticsDataAgent#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// Optional. The value for the field to filter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#value GoogleGeminiDataAnalyticsDataAgent#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

