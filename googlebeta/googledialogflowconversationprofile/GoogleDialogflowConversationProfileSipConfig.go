// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileSipConfig struct {
	// Allows interactions with a Dialogflow virtual agent even if the call is connected for SIPREC purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_dialogflow_conversation_profile#allow_virtual_agent_interaction GoogleDialogflowConversationProfile#allow_virtual_agent_interaction}
	AllowVirtualAgentInteraction interface{} `field:"optional" json:"allowVirtualAgentInteraction" yaml:"allowVirtualAgentInteraction"`
	// List of inbound call leg headers to be copied to outbound call legs created later.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_dialogflow_conversation_profile#copy_inbound_call_leg_headers GoogleDialogflowConversationProfile#copy_inbound_call_leg_headers}
	CopyInboundCallLegHeaders *[]*string `field:"optional" json:"copyInboundCallLegHeaders" yaml:"copyInboundCallLegHeaders"`
	// Asks Dialogflow Telephony to create the conversation provided in the SIP header on the fly when the call comes in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_dialogflow_conversation_profile#create_conversation_on_the_fly GoogleDialogflowConversationProfile#create_conversation_on_the_fly}
	CreateConversationOnTheFly interface{} `field:"optional" json:"createConversationOnTheFly" yaml:"createConversationOnTheFly"`
	// Ignores any media direction in the reINVITE SDP offer. Reuse the previous media direction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_dialogflow_conversation_profile#ignore_reinvite_media_direction GoogleDialogflowConversationProfile#ignore_reinvite_media_direction}
	IgnoreReinviteMediaDirection interface{} `field:"optional" json:"ignoreReinviteMediaDirection" yaml:"ignoreReinviteMediaDirection"`
	// Starts the conversation with inactive SDP directives.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_dialogflow_conversation_profile#inactive_start GoogleDialogflowConversationProfile#inactive_start}
	InactiveStart interface{} `field:"optional" json:"inactiveStart" yaml:"inactiveStart"`
	// Keeps the conversation running even if the call is disconnected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_dialogflow_conversation_profile#keep_conversation_running GoogleDialogflowConversationProfile#keep_conversation_running}
	KeepConversationRunning interface{} `field:"optional" json:"keepConversationRunning" yaml:"keepConversationRunning"`
	// Max duration for audio recording. Overrides the default value of 15 min. Max value is 8 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.46.1/docs/resources/google_dialogflow_conversation_profile#max_audio_recording_duration GoogleDialogflowConversationProfile#max_audio_recording_duration}
	MaxAudioRecordingDuration *string `field:"optional" json:"maxAudioRecordingDuration" yaml:"maxAudioRecordingDuration"`
}

