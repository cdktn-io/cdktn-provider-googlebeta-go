// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextDatasourceReferencesBqTableReferences struct {
	// Required. The dataset that the table belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#dataset_id GoogleGeminiDataAnalyticsDataAgent#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// Required. The project that the table belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#project_id GoogleGeminiDataAnalyticsDataAgent#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Required. The table ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#table_id GoogleGeminiDataAnalyticsDataAgent#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
}

