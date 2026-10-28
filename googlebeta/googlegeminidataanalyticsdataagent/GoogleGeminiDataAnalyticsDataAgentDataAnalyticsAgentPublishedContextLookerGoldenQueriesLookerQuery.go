// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextLookerGoldenQueriesLookerQuery struct {
	// Required. The LookML explore used to generate the query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#explore GoogleGeminiDataAnalyticsDataAgent#explore}
	Explore *string `field:"required" json:"explore" yaml:"explore"`
	// Required. The LookML model used to generate the query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#model GoogleGeminiDataAnalyticsDataAgent#model}
	Model *string `field:"required" json:"model" yaml:"model"`
	// Optional. The short alphanumeric identifier for the query, used for share links and Explore URLs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#client_id GoogleGeminiDataAnalyticsDataAgent#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Optional. The fields to retrieve from the explore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#fields GoogleGeminiDataAnalyticsDataAgent#fields}
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#filters GoogleGeminiDataAnalyticsDataAgent#filters}
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// Optional. Limit in the query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#limit GoogleGeminiDataAnalyticsDataAgent#limit}
	Limit *string `field:"optional" json:"limit" yaml:"limit"`
	// Optional. The primary identifier for the query resource in Looker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#query_id GoogleGeminiDataAnalyticsDataAgent#query_id}
	QueryId *string `field:"optional" json:"queryId" yaml:"queryId"`
	// Optional. The sorts to apply to the explore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#sorts GoogleGeminiDataAnalyticsDataAgent#sorts}
	Sorts *[]*string `field:"optional" json:"sorts" yaml:"sorts"`
}

