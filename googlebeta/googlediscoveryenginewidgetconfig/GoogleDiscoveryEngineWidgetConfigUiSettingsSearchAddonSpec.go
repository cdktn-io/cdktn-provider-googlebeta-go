// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlediscoveryenginewidgetconfig


type GoogleDiscoveryEngineWidgetConfigUiSettingsSearchAddonSpec struct {
	// If true, generative answer add-on is disabled. Generative answer add-on includes natural language to filters and simple answers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_discovery_engine_widget_config#generative_answer_add_on_disabled GoogleDiscoveryEngineWidgetConfig#generative_answer_add_on_disabled}
	GenerativeAnswerAddOnDisabled interface{} `field:"optional" json:"generativeAnswerAddOnDisabled" yaml:"generativeAnswerAddOnDisabled"`
	// If true, disables event re-ranking and personalization to optimize KPIs & personalize results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_discovery_engine_widget_config#kpi_personalization_add_on_disabled GoogleDiscoveryEngineWidgetConfig#kpi_personalization_add_on_disabled}
	KpiPersonalizationAddOnDisabled interface{} `field:"optional" json:"kpiPersonalizationAddOnDisabled" yaml:"kpiPersonalizationAddOnDisabled"`
	// If true, semantic add-on is disabled. Semantic add-on includes embeddings and jetstream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_discovery_engine_widget_config#semantic_add_on_disabled GoogleDiscoveryEngineWidgetConfig#semantic_add_on_disabled}
	SemanticAddOnDisabled interface{} `field:"optional" json:"semanticAddOnDisabled" yaml:"semanticAddOnDisabled"`
}

