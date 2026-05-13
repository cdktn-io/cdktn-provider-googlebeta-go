// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesapp


type GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholds struct {
	// expectation_level_metrics_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_app#expectation_level_metrics_thresholds GoogleCesApp#expectation_level_metrics_thresholds}
	ExpectationLevelMetricsThresholds *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsExpectationLevelMetricsThresholds `field:"optional" json:"expectationLevelMetricsThresholds" yaml:"expectationLevelMetricsThresholds"`
	// turn_level_metrics_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_app#turn_level_metrics_thresholds GoogleCesApp#turn_level_metrics_thresholds}
	TurnLevelMetricsThresholds *GoogleCesAppEvaluationMetricsThresholdsGoldenEvaluationMetricsThresholdsTurnLevelMetricsThresholds `field:"optional" json:"turnLevelMetricsThresholds" yaml:"turnLevelMetricsThresholds"`
}

