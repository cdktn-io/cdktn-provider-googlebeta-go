// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesapp


type GoogleCesAppLoggingSettingsMetricAnalysisSettings struct {
	// Whether to collect conversation data for llm analysis metrics.
	//
	// If true,
	// conversation data will not be collected for llm analysis metrics;
	// otherwise, conversation data will be collected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_app#llm_metrics_opted_out GoogleCesApp#llm_metrics_opted_out}
	LlmMetricsOptedOut interface{} `field:"optional" json:"llmMetricsOptedOut" yaml:"llmMetricsOptedOut"`
}

