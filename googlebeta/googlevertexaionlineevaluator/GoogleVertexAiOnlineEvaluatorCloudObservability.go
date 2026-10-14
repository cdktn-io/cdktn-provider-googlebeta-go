// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaionlineevaluator


type GoogleVertexAiOnlineEvaluatorCloudObservability struct {
	// open_telemetry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#open_telemetry GoogleVertexAiOnlineEvaluator#open_telemetry}
	OpenTelemetry *GoogleVertexAiOnlineEvaluatorCloudObservabilityOpenTelemetry `field:"required" json:"openTelemetry" yaml:"openTelemetry"`
	// Optional log view that will be used to query logs. If empty, the project's default view ('projects/{{project_id}}') will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#log_view GoogleVertexAiOnlineEvaluator#log_view}
	LogView *string `field:"optional" json:"logView" yaml:"logView"`
	// session_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#session_scope GoogleVertexAiOnlineEvaluator#session_scope}
	SessionScope *GoogleVertexAiOnlineEvaluatorCloudObservabilitySessionScope `field:"optional" json:"sessionScope" yaml:"sessionScope"`
	// trace_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#trace_scope GoogleVertexAiOnlineEvaluator#trace_scope}
	TraceScope *GoogleVertexAiOnlineEvaluatorCloudObservabilityTraceScope `field:"optional" json:"traceScope" yaml:"traceScope"`
	// Optional trace view that will be used to query traces.
	//
	// If empty, the
	// '_AllSpans' view from the '_Trace' US bucket will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.7.0/docs/resources/google_vertex_ai_online_evaluator#trace_view GoogleVertexAiOnlineEvaluator#trace_view}
	TraceView *string `field:"optional" json:"traceView" yaml:"traceView"`
}

