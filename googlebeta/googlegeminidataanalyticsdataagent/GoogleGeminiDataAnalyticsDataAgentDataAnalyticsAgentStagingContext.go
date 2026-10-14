// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContext struct {
	// datasource_references block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#datasource_references GoogleGeminiDataAnalyticsDataAgent#datasource_references}
	DatasourceReferences *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferences `field:"required" json:"datasourceReferences" yaml:"datasourceReferences"`
	// example_queries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#example_queries GoogleGeminiDataAnalyticsDataAgent#example_queries}
	ExampleQueries interface{} `field:"optional" json:"exampleQueries" yaml:"exampleQueries"`
	// glossary_terms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#glossary_terms GoogleGeminiDataAnalyticsDataAgent#glossary_terms}
	GlossaryTerms interface{} `field:"optional" json:"glossaryTerms" yaml:"glossaryTerms"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#options GoogleGeminiDataAnalyticsDataAgent#options}
	Options *GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptions `field:"optional" json:"options" yaml:"options"`
	// schema_relationships block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#schema_relationships GoogleGeminiDataAnalyticsDataAgent#schema_relationships}
	SchemaRelationships interface{} `field:"optional" json:"schemaRelationships" yaml:"schemaRelationships"`
	// Optional. The basic entry point for data owners creating domain knowledge for the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#system_instruction GoogleGeminiDataAnalyticsDataAgent#system_instruction}
	SystemInstruction *string `field:"optional" json:"systemInstruction" yaml:"systemInstruction"`
}

