// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigCustomizationConfigsGenerateMemoriesExamplesConversationSourceEventsContentPartsAudioTranscriptionWords struct {
	// Transcript of the word.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_reasoning_engine#word GoogleVertexAiReasoningEngine#word}
	Word *string `field:"required" json:"word" yaml:"word"`
	// End offset in time of the word relative to the start of the audio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_reasoning_engine#end_offset GoogleVertexAiReasoningEngine#end_offset}
	EndOffset *string `field:"optional" json:"endOffset" yaml:"endOffset"`
	// Start offset in time of the word relative to the start of the audio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_vertex_ai_reasoning_engine#start_offset GoogleVertexAiReasoningEngine#start_offset}
	StartOffset *string `field:"optional" json:"startOffset" yaml:"startOffset"`
}

