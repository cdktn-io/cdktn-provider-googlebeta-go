// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextDatasourceReferencesLookerExploreReferencesPrivateLookerInstanceInfo struct {
	// Optional. The Looker instance ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#looker_instance_id GoogleGeminiDataAnalyticsDataAgent#looker_instance_id}
	LookerInstanceId *string `field:"optional" json:"lookerInstanceId" yaml:"lookerInstanceId"`
	// Optional. The service directory name of the Looker instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_gemini_data_analytics_data_agent#service_directory_name GoogleGeminiDataAnalyticsDataAgent#service_directory_name}
	ServiceDirectoryName *string `field:"optional" json:"serviceDirectoryName" yaml:"serviceDirectoryName"`
}

