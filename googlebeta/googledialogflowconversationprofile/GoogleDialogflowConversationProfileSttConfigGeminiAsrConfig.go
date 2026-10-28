// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileSttConfigGeminiAsrConfig struct {
	// End of speech sensitivity. Possible values: ["END_SENSITIVITY_HIGH", "END_SENSITIVITY_LOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_dialogflow_conversation_profile#end_of_speech_sensitivity GoogleDialogflowConversationProfile#end_of_speech_sensitivity}
	EndOfSpeechSensitivity *string `field:"optional" json:"endOfSpeechSensitivity" yaml:"endOfSpeechSensitivity"`
	// The Gemini ASR model ID used for transcription. This value overrides the default model ID configured on the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_dialogflow_conversation_profile#model_id GoogleDialogflowConversationProfile#model_id}
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
	// The required duration of detected speech before start-of-speech is committed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_dialogflow_conversation_profile#prefix_padding_ms GoogleDialogflowConversationProfile#prefix_padding_ms}
	PrefixPaddingMs *float64 `field:"optional" json:"prefixPaddingMs" yaml:"prefixPaddingMs"`
	// The required duration of detected silence (or non-speech) before end-of-speech is committed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_dialogflow_conversation_profile#silence_duration_ms GoogleDialogflowConversationProfile#silence_duration_ms}
	SilenceDurationMs *float64 `field:"optional" json:"silenceDurationMs" yaml:"silenceDurationMs"`
	// Start of speech sensitivity. Possible values: ["START_SENSITIVITY_HIGH", "START_SENSITIVITY_LOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.9.0/docs/resources/google_dialogflow_conversation_profile#start_of_speech_sensitivity GoogleDialogflowConversationProfile#start_of_speech_sensitivity}
	StartOfSpeechSensitivity *string `field:"optional" json:"startOfSpeechSensitivity" yaml:"startOfSpeechSensitivity"`
}

