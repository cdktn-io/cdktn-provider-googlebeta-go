// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlecesapp


type GoogleCesAppErrorHandlingSettings struct {
	// end_session_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_app#end_session_config GoogleCesApp#end_session_config}
	EndSessionConfig *GoogleCesAppErrorHandlingSettingsEndSessionConfig `field:"optional" json:"endSessionConfig" yaml:"endSessionConfig"`
	// The strategy to use for error handling. Possible values: NONE FALLBACK_RESPONSE END_SESSION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_app#error_handling_strategy GoogleCesApp#error_handling_strategy}
	ErrorHandlingStrategy *string `field:"optional" json:"errorHandlingStrategy" yaml:"errorHandlingStrategy"`
	// fallback_response_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_ces_app#fallback_response_config GoogleCesApp#fallback_response_config}
	FallbackResponseConfig *GoogleCesAppErrorHandlingSettingsFallbackResponseConfig `field:"optional" json:"fallbackResponseConfig" yaml:"fallbackResponseConfig"`
}

