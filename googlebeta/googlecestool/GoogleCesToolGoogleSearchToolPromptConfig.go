// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecestool


type GoogleCesToolGoogleSearchToolPromptConfig struct {
	// Optional.
	//
	// Defines the prompt used for the system instructions when interacting with the
	// agent in chat conversations. If not set, default prompt will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_ces_tool#text_prompt GoogleCesTool#text_prompt}
	TextPrompt *string `field:"optional" json:"textPrompt" yaml:"textPrompt"`
	// Optional.
	//
	// Defines the prompt used for the system instructions when interacting with the
	// agent in voice conversations. If not set, default prompt will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.39.0/docs/resources/google_ces_tool#voice_prompt GoogleCesTool#voice_prompt}
	VoicePrompt *string `field:"optional" json:"voicePrompt" yaml:"voicePrompt"`
}

