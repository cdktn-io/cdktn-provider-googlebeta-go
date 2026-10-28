// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferencesBqPropertyGraphReferences struct {
	// Required. The dataset that the property graph belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#dataset_id GoogleGeminiDataAnalyticsDataAgent#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// Required. The project that the property graph belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#project_id GoogleGeminiDataAnalyticsDataAgent#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Required. The property graph ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#property_graph_id GoogleGeminiDataAnalyticsDataAgent#property_graph_id}
	PropertyGraphId *string `field:"required" json:"propertyGraphId" yaml:"propertyGraphId"`
}

