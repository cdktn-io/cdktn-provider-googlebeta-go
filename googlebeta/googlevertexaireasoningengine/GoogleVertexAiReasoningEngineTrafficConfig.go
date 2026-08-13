// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineTrafficConfig struct {
	// traffic_split_always_latest block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_reasoning_engine#traffic_split_always_latest GoogleVertexAiReasoningEngine#traffic_split_always_latest}
	TrafficSplitAlwaysLatest *GoogleVertexAiReasoningEngineTrafficConfigTrafficSplitAlwaysLatest `field:"optional" json:"trafficSplitAlwaysLatest" yaml:"trafficSplitAlwaysLatest"`
	// traffic_split_manual block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.44.0/docs/resources/google_vertex_ai_reasoning_engine#traffic_split_manual GoogleVertexAiReasoningEngine#traffic_split_manual}
	TrafficSplitManual *GoogleVertexAiReasoningEngineTrafficConfigTrafficSplitManual `field:"optional" json:"trafficSplitManual" yaml:"trafficSplitManual"`
}

