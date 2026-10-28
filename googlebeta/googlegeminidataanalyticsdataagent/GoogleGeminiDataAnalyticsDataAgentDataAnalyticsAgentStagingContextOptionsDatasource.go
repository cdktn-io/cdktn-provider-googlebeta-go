// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentStagingContextOptionsDatasource struct {
	// Optional. Limits the bytes billed for each BigQuery query job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#big_query_max_billed_bytes GoogleGeminiDataAnalyticsDataAgent#big_query_max_billed_bytes}
	BigQueryMaxBilledBytes *float64 `field:"optional" json:"bigQueryMaxBilledBytes" yaml:"bigQueryMaxBilledBytes"`
}

