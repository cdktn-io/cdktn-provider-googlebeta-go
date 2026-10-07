// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsAudioTranscription struct {
	// The transcription text of this audio segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_reasoning_engine#text GoogleVertexAiReasoningEngine#text}
	Text *string `field:"required" json:"text" yaml:"text"`
	// A label identifying the speaker of this audio segment (e.g. spk_1, spk_2). Present when diarization is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_reasoning_engine#speaker_label GoogleVertexAiReasoningEngine#speaker_label}
	SpeakerLabel *string `field:"optional" json:"speakerLabel" yaml:"speakerLabel"`
	// words block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_reasoning_engine#words GoogleVertexAiReasoningEngine#words}
	Words interface{} `field:"optional" json:"words" yaml:"words"`
}

