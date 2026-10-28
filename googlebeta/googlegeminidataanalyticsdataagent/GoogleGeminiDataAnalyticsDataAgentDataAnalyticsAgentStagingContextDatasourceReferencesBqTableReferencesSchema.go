// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferencesBqTableReferencesSchema struct {
	// Optional. A textual description of the table's content and purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#description GoogleGeminiDataAnalyticsDataAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional. Table display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#display_name GoogleGeminiDataAnalyticsDataAgent#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#fields GoogleGeminiDataAnalyticsDataAgent#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#filters GoogleGeminiDataAnalyticsDataAgent#filters}
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// Optional. A list of alternative names or synonyms that can be used to refer to the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#synonyms GoogleGeminiDataAnalyticsDataAgent#synonyms}
	Synonyms *[]*string `field:"optional" json:"synonyms" yaml:"synonyms"`
	// Optional. A list of tags or keywords associated with the table, used for categorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#tags GoogleGeminiDataAnalyticsDataAgent#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

