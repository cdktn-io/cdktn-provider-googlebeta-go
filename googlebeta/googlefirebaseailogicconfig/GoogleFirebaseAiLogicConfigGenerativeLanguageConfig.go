// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package googlefirebaseailogicconfig


type GoogleFirebaseAiLogicConfigGenerativeLanguageConfig struct {
	// The value of the API key.
	//
	// Firebase AI Logic no longer requires API key
	// material; values provided here may be silently ignored on input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_firebase_ai_logic_config#api_key GoogleFirebaseAiLogicConfig#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// The value of the API key.
	//
	// Firebase AI Logic no longer requires API key
	// material; values provided here may be silently ignored on input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_firebase_ai_logic_config#api_key_wo GoogleFirebaseAiLogicConfig#api_key_wo}
	ApiKeyWo *string `field:"optional" json:"apiKeyWo" yaml:"apiKeyWo"`
	// Triggers update of 'api_key_wo' write-only.
	//
	// Increment this value when an update to 'api_key_wo' is needed. For more info see [updating write-only arguments](/docs/providers/google/guides/using_write_only_arguments.html#updating-write-only-arguments)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/8.6.0/docs/resources/google_firebase_ai_logic_config#api_key_wo_version GoogleFirebaseAiLogicConfig#api_key_wo_version}
	ApiKeyWoVersion *string `field:"optional" json:"apiKeyWoVersion" yaml:"apiKeyWoVersion"`
}

