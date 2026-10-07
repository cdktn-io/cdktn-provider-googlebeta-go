// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlegeminigibqobservabilitysetting


type GoogleGeminiGibqObservabilitySettingConversationalAnalyticsSetting struct {
	// Whether to enable feedback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_gemini_gibq_observability_setting#feedback_enabled GoogleGeminiGibqObservabilitySetting#feedback_enabled}
	FeedbackEnabled interface{} `field:"optional" json:"feedbackEnabled" yaml:"feedbackEnabled"`
	// Whether to enable logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_gemini_gibq_observability_setting#logging_enabled GoogleGeminiGibqObservabilitySetting#logging_enabled}
	LoggingEnabled interface{} `field:"optional" json:"loggingEnabled" yaml:"loggingEnabled"`
	// Whether to enable metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_gemini_gibq_observability_setting#metrics_enabled GoogleGeminiGibqObservabilitySetting#metrics_enabled}
	MetricsEnabled interface{} `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// Whether to enable traces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_gemini_gibq_observability_setting#traces_enabled GoogleGeminiGibqObservabilitySetting#traces_enabled}
	TracesEnabled interface{} `field:"optional" json:"tracesEnabled" yaml:"tracesEnabled"`
}

