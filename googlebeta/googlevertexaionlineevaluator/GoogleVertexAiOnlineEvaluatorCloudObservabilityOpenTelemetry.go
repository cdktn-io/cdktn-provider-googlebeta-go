// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaionlineevaluator


type GoogleVertexAiOnlineEvaluatorCloudObservabilityOpenTelemetry struct {
	// Defines which version of the OTel Semantic Convention the data follows. Can be "1.39.0" or newer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_vertex_ai_online_evaluator#semconv_version GoogleVertexAiOnlineEvaluator#semconv_version}
	SemconvVersion *string `field:"required" json:"semconvVersion" yaml:"semconvVersion"`
}

