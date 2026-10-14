// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminidataanalyticsdataagent


type GoogleGeminiDataAnalyticsDataAgentDataAnalyticsAgentPublishedContextGlossaryTerms struct {
	// Required. The description or meaning of the term.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#description GoogleGeminiDataAnalyticsDataAgent#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Required. User friendly display name of the glossary term.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#display_name GoogleGeminiDataAnalyticsDataAgent#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Optional. A list of general purpose labels associated with this term.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_gemini_data_analytics_data_agent#labels GoogleGeminiDataAnalyticsDataAgent#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
}

