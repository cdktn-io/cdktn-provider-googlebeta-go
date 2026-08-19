// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfigGenerationTriggerConfigGenerationRule struct {
	// Optional. Specifies to trigger generation when the event count reaches this limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#event_count GoogleVertexAiReasoningEngine#event_count}
	EventCount *float64 `field:"optional" json:"eventCount" yaml:"eventCount"`
	// Optional. Specifies to trigger generation at a fixed interval. The duration must have a minute-level granularity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#fixed_interval GoogleVertexAiReasoningEngine#fixed_interval}
	FixedInterval *string `field:"optional" json:"fixedInterval" yaml:"fixedInterval"`
	// Optional.
	//
	// Specifies to trigger generation if the stream is inactive for the
	// specified duration after the most recent event. The duration must have a
	// minute-level granularity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#idle_duration GoogleVertexAiReasoningEngine#idle_duration}
	IdleDuration *string `field:"optional" json:"idleDuration" yaml:"idleDuration"`
	// Optional. Re-include the last N already-processed events in the next window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.45.0/docs/resources/google_vertex_ai_reasoning_engine#overlap_event_count GoogleVertexAiReasoningEngine#overlap_event_count}
	OverlapEventCount *float64 `field:"optional" json:"overlapEventCount" yaml:"overlapEventCount"`
}

