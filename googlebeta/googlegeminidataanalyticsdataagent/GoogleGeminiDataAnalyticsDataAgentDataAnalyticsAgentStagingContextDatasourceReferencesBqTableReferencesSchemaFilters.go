// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferencesBqTableReferencesSchemaFilters struct {
	// Optional. The field to filter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#field GoogleGeminiDataAnalyticsDataAgent#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Optional. The type of filter present on a datasource, such as ALWAYS_FILTER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#type GoogleGeminiDataAnalyticsDataAgent#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Optional. The default value used for this filter if the filter is not overridden in a query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#value GoogleGeminiDataAnalyticsDataAgent#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

