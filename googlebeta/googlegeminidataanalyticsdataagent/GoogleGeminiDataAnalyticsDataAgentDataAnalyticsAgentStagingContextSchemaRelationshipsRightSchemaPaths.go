// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextSchemaRelationshipsRightSchemaPaths struct {
	// Required. The ordered list of paths within the table schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#paths GoogleGeminiDataAnalyticsDataAgent#paths}
	Paths *[]*string `field:"required" json:"paths" yaml:"paths"`
	// Required. The service-qualified full resource name of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#table_fqn GoogleGeminiDataAnalyticsDataAgent#table_fqn}
	TableFqn *string `field:"required" json:"tableFqn" yaml:"tableFqn"`
}

